package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 初始化一个包含两个文件路径的切片
	// 注意，home.page.tmpl 文件必须是切片中的 *第一个* 文件。
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// 使用 template.ParseFiles() 函数将模板文件读取到模板集中
	// 如果出错，记录详细错误信息并使用 http.Error() 函数向用户发送
	// 500 Internal Server Error 响应
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// 然后使用模板集的 Execute() 方法将模板内容写入响应体
	// Execute() 的最后一个参数用于传入动态数据，目前暂设为 nil
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
