package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/maxfeizi04-cloude/snippetbox/pkg/models"
	"github.com/maxfeizi04-cloude/snippetbox/pkg/models/mysql"
)

// 在 application 结构体中添加 snippets 字段
// 这将使 SnippetModel 对象在 handlers 中可用
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// 因为 Pat 会精确匹配 "/" 路径，所以我们现在可以移除该 handler 中
	// 对 r.URL.Path != "/" 的手动检查

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Pat 不会自动去掉命名捕获参数中的冒号，因此我们需要从查询字符串中获取
	// ":id" 的值，而不是 "id"
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})

}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// 首先我们调用 r.ParseForm()，它会将 POST 请求体中的数据添加到
	// r.PostForm 映射中。对于 PUT 和 PATCH 请求，它的工作方式也是相同的
	// 如果出现任何错误，我们使用 app.clientError 辅助函数向用户发送
	// 400 Bad Request 响应
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// 使用 r.PostForm.Get() 方法从 r.PostForm 映射中提取相关的数据字段
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// 将数据传给 SnippetModel.Insert() 方法，并接收新记录的 ID
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// 将客户端重定向到新创建的 snippet 详情页面
	// 使用 http.StatusSeeOther (303) 状态码,这是 POST 请求后重定向的标准做法
	// 可以防止用户刷新页面时重复提交表单
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
