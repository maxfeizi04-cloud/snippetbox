package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/maxfeizi04-cloude/snippetbox/pkg/models"
)

// 定义 templateData 类型，作为向 HTML 模板传递动态数据的载体结构体
// 目前它只包含一个字段，但随着构建的推进，我们会继续添加更多字段
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// 创建一个 humanDate 函数，返回 time.Time 对象的格式化字符串表示
func humanDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 初始化一个 template.FuncMap 对象并存储在全局变量中
// 这本质上是一个字符串键的 map，用于自定义模板函数名称与实际函数之间的查找
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// 初始化一个新的 map 作为缓存
	cache := map[string]*template.Template{}

	// 使用 filepath.Glob 函数获取所有扩展名为 '.page.tmpl' 的文件路径切片
	// 这本质上给了我们一个应用中所有 '页面' 模板的切片
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// 逐个遍历所有页面
	for _, page := range pages {
		// 从完整文件路径中提取文件名（如 'home.page.tmpl'）
		// 并将其赋值给 name 变量
		name := filepath.Base(page)

		// template.FuncMap 必须在调用 ParseFiles() 方法之前注册到模板集中
		// 这意味着我们必须使用 template.New() 创建一个空的模板集，
		// 使用 Funcs() 方法注册 template.FuncMap，
		// 然后正常解析文件
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// 使用 ParseGlob 方法将所有 '布局' 模板添加到模板集合中
		// (在我们的例子中，目前只有一个 'base' 布局)
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// 使用 ParseGlob 方法将所有 '局部' 模板添加到模板集合中
		// (在我们的例子中，目前只有一个 'footer' 局部模板)
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		// 将模板集合添加到缓存中，使用页面名称(如 'home.page.tmpl')作为键
		cache[name] = ts
	}

	return cache, nil
}
