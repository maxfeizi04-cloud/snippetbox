package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
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
