package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// 更新 routes() 方法的签名，使其返回一个 http.Handler
// 而不是 *http.ServeMux。
func (app *application) routes() http.Handler {
	// 创建一个包含"标准"中间件的中间件链，该链将用于应用接收到的每个请求
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// 返回"标准"中间件链后跟 servemux
	return standardMiddleware.Then(mux)
}
