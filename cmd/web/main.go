package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/maxfeizi04-cloude/snippetbox/pkg/models/mysql"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 为保持 main() 函数简洁，我将创建连接池的代码放到了下面的 openDB() 函数中
	// 我们将命令行标志中的 DSN 传给 openDB()
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// 我们还延迟调用 db.Close()
	// 以确保在 main() 函数退出前关闭连接池
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// 初始化一个新的 http.Server 结构体。设置 Addr 和 Handler 字段
	// 使服务器使用与之前相同的网络地址和路由，并设置 ErrorLog 字段
	// 以便服务器在出现任何问题时使用自定义的 errorLog logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	// 因为 err 变量在上面的代码中已经声明过了，所以这里需要使用赋值运算符 =，
	// 而不能使用 :=（声明并赋值）运算符
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB() 函数封装了 sql.Open()
// 返回给定 DSN 对应的 sql.DB 连接池
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
