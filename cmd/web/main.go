package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
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
	// 调用新 http.Server 结构体的 ListenAndServe() 方法
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
