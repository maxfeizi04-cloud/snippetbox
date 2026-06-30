package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

// serverError 辅助函数将错误信息和堆栈跟踪写入 errorLog
// 然后向用户发送通用的 500 Internal Server Error 响应
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError 辅助函数向用户发送指定的状态码及对应的描述信息
// 本书后面将用它来发送诸如 400 "Bad Request" 之类的响应
// 用于处理用户请求中的问题
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// 为保持一致性，我们还实现一个 notFound 辅助函数
// 它只是对 clientError 的便捷封装，用于向用户发送 404 Not Found 响应
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// 创建一个 addDefaultData 辅助方法。该方法接收一个指向 templateData 结构体的指针
// 将当前年份添加到 CurrentYear 字段中，然后返回该指针
// 目前我们还没有使用 *http.Request 参数，但在本书后面会用到
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// 根据页面名称（如 'home.page.tmpl'）从缓存中检索对应的模板集合
	// 如果在缓存中找不到指定名称的条目，就调用之前创建的 serverError 辅助方法
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	// 初始化一个新的缓冲区
	buf := new(bytes.Buffer)

	// 将模板写入缓冲区，而不是直接写入 http.ResponseWriter
	// 如果出现错误，调用 serverError 辅助方法然后返回
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	// 将缓冲区的内容写入 http.ResponseWriter
	// 同样,这是我们将 http.ResponseWriter 传递给接受 io.Writer 的函数的另一次实践
	buf.WriteTo(w)
}
