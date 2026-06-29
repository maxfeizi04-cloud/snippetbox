package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/maxfeizi04-cloude/snippetbox/pkg/models"
	"github.com/maxfeizi04-cloude/snippetbox/pkg/models/mysql"
)

// 在 application 结构体中添加 snippets 字段
// 这将使 SnippetModel 对象在 handlers 中可用
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// 初始化一个包含两个文件路径的切片
	// 注意，home.page.tmpl 文件必须是切片中的 *第一个* 文件。
	//files := []string{
	//	"./ui/html/home.page.tmpl",
	//	"./ui/html/base.layout.tmpl",
	//	"./ui/html/footer.partial.tmpl",
	//}

	// 使用 template.ParseFiles() 函数将模板文件读取到模板集中
	// 如果出错，记录详细错误信息并使用 http.Error() 函数向用户发送
	//// 500 Internal Server Error 响应
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}

	// 然后使用模板集的 Execute() 方法将模板内容写入响应体
	// Execute() 的最后一个参数用于传入动态数据，目前暂设为 nil
	//err = ts.Execute(w, nil)
	//if err != nil {
	//	app.serverError(w, err)
	//}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// 使用 SnippetModel 对象的 Get 方法，根据 ID 检索特定记录的数据。
	// 如果未找到匹配记录，则返回 404 Not Found 响应
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// 将 snippet 数据作为纯文本 HTTP 响应体写入
	fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// 创建一些存有测试数据的变量。稍后构建时会移除这些数据
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n-Kobayashi Issa"
	expires := "7"

	// 将数据传给 SnippetModel.Insert() 方法，并接收新记录的 ID
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// 将用户重定向到该 snippet 对应的页面
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
