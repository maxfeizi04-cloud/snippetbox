![image-20260625162532281](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260625162532281.png)

​	Let's Go 逐步教您如何使用出色的编程语言 Go 创建快速、安全和可维护的 Web 应用程序。

​	本书背后的理念是帮助您边做边学。我们将一起完成 Web 应用程序从头到尾的构建过程——从构建您的工作区到会话管理、验证用户、保护您的服务器和测试您的应用程序。

本书背后的理念是帮助您边做边学。我们将一起完成 Web 应用程序从头到尾的构建过程——从构建您的工作区到会话管理、验证用户、保护您的服务器和测试您的应用程序。

​	以这种方式构建完整的 Web 应用程序有几个好处。它有助于将你正在学习的东西置于上下文中，它展示了你的代码库的不同部分如何链接在一起，并迫使我们解决在现实生活中编写软件时出现的边缘情况和困难。从本质上讲，你会比仅仅阅读 Go 的（伟大的）文档或独立的博客文章学到更多。

​	到本书结束时，您将了解并有信心使用 Go 构建您自己的生产就绪 Web 应用程序。

​	虽然您可以从头到尾阅读本书，但它是专门设计的，因此您可以自己跟着项目构建。

​	打破你的文本编辑器，快乐编码！

— Alex



# Contents 内容

### 1. 简介

#### 介绍 

---

​	在本书中，我们将构建一个名为 Snippetbox 的 Web 应用程序，它可以让人们粘贴和共享文本片段——有点像 Pastebin 或 GitHub 的 Gists。在构建结束时，它看起来有点像这样：

![image-20260625165707082](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260625165707082.png)

​	我们的应用程序一开始非常简单，只有一个网页。然后在每一章中，我们将逐步构建它，直到用户能够通过该应用程序保存和查看片段。这将带我们了解如何构建项目、路由请求、使用数据库、处理表单和安全地显示动态数据等主题。

​	然后在本书的后面，我们将添加用户帐户，并限制应用程序，以便只有注册用户才能创建片段。这将带我们了解更高级的主题，例如配置 HTTPS 服务器、会话管理、用户身份验证和中间件。



#### 惯例

---

​	在本书中，代码块显示为银色背景，如下所示。如果代码特别长，不相关的部分可以用省略号代替。为了便于跟进，大多数代码块的顶部还有一个标题栏，指示我们正在处理的文件的名称。

```
file: hello.go
```

```go
package main

func sayHello() {
    fmt.Println("Hello world!")
}
```

> [!TIP]
>
> 提示：如果您跟随应用程序构建，我建议使用本书的HTML版本，而不是PDF或EPUB版本。HTML版本适用于所有浏览器，如果您想直接从书中复制和粘贴代码，则保留代码块的适当格式。

​	终端（命令行）指令以银色背景显示并以美元符号开头。这些命令应该适用于任何基于 Unix 的操作系统，包括 Mac OSX 和 Linux。示例输出以银色显示在命令下方，如下所示：

```bash
$ echo "Hello world!"
Hello world!
```

​	如果您使用的是 Windows，则应将命令替换为 DOS 等效命令或通过普通 Windows GUI 执行操作。

​	本书的一些章节以附加信息部分结束。这些部分包含与我们的应用构建无关但仍然重要（有时只是有趣）的信息。如果您对 Go 还不太熟悉，您可能需要跳过这些部分并稍后再回头看。



#### 关于作者

---

​	嘿，我是 Alex Edwards，全栈网络开发人员和作家。我住在奥地利因斯布鲁克附近。

​	我使用 Go 已经 8 年多了，为我自己和商业客户构建生产应用程序，并帮助世界各地的人们提高他们的 Go 技能。

​	你可以在我的博客上看到更多我的文章（我在那里发布详细的教程），在 GitHub 上看到我的一些开源作品，你也可以在 Instagram 和 Twitter 上关注我。



#### 版权和免责声明

---

​	Let's Go：学习使用 Go 构建专业的 Web 应用程序。版权所有 © 2022 亚历克斯·爱德华兹。

​	最后更新时间为 2022-08-04 11:31:25 UTC。版本 2.19.0。

​	Go gopher 由 Renee French 设计，并在 Creative Commons 3.0 Attributions 许可下使用。 Cover gopher 改编自 Egon Elbre 的矢量。

​	本书中提供的信息仅供一般参考。虽然作者和出版商已尽一切努力确保本书中信息的准确性在出版时是正确的，但对于完整性、准确性、可靠性、适用性或可用性，我们不作任何明示或暗示的陈述或保证出于任何目的而使用本书中包含的信息、产品、服务或相关图形。使用此信息的风险由您自行承担。



#### 先决条件

---

#### 背景知识

​	这本书是为刚接触 Go 的人设计的，但如果你先对 Go 的语法有一个大致的了解，你可能会发现它更有趣。如果你发现自己在语法上苦苦挣扎，Karl Seguin 的 Little Book of Go 是一个很棒的教程，或者如果你想要更具交互性的东西，我建议你完成 Go 之旅。

​	我还假设您对 HTML/CSS 和 SQL 有（非常）基本的了解，并且对使用终端（或 Windows 用户的命令行）有一定的了解。如果您以前用任何其他语言构建过 Web 应用程序——无论是 Ruby、Python、PHP 还是 C#——那么这本书应该非常适合您。



#### 其他软件

---

​	如果您想完全按照说明进行操作，则应确保您的计算机上还有其他一些软件可用。他们是：

​	用于处理来自终端的 HTTP 请求和响应的 curl 工具。在 MacOS 和 Linux 机器上，它应该预先安装或在您的软件存储库中可用。否则，您可以从此处下载最新版本。

​	具有良好开发工具的 Web 浏览器。我将在本书中使用 Firefox，但 Chromium、Chrome 或 Microsoft Edge 也可以。

### 2. 基础

​	好了，让我们开始吧！在本书的第一部分中，我们将为我们的项目奠定基础，并解释您在构建应用程序的其余部分时需要了解的主要原则。

您将学习如何：

1. 设置一个遵循 Go 约定的项目目录。
2. 启动 Web 服务器并侦听传入的 HTTP 请求。
3. 根据请求路径将请求路由到不同的处理程序。
4. 向用户发送不同的 HTTP 响应、标头和状态代码。
5. 从 URL 查询字符串参数中获取并验证不受信任的用户输入。
6. 以合理且可扩展的方式构建您的项目。
7. 呈现 HTML 页面并使用模板继承使您的标记没有重复的样板代码。
8. 从您的应用程序提供静态文件，如图像、CSS 和 JavaScript。

#### 2.1 安装 Go

---



​	本书中的信息适用于 Go 的最新主要版本（1.19 版），如果您想在构建应用程序的同时进行编码，则应该安装它。

​	如果你已经安装了 Go，你可以使用 go version命令从你的终端检查版本号。输出应与此类似：

```bash
$ go version
go version go1.19 linux/amd64
```

​	如果您需要升级您的 Go 版本——或者从头开始安装 Go——那么请现在就开始吧。可以在此处找到针对不同操作系统的详细说明：

- 删除旧版本的 Go
- 在 Mac OS X 上安装 Go
- 在 Windows 上安装 Go
- 在 Linux 上安装 Go

#### 2.2. 项目设置和创建模块 

​	在我们编写任何代码之前，您需要在您的计算机上创建一个 snippetbox 目录作为该项目的顶级“家”。我们在整本书中编写的所有 Go 代码以及其他项目特定的资产（如 HTML 模板和 CSS 文件）都将保存在这里。

​	因此，如果您正在跟随，请打开您的终端并在您计算机上的任意位置创建一个名为 snippetbox的新项目目录。我将在 $HOME/code下找到我的项目目录，但您可以根据需要选择其他位置。

```bash
$ mkdir -p $HOME/code/snippetbox
```

##### 创建模块

​	接下来您需要做的是为您的项目选择一个模块路径。

​	如果您还不熟悉 Go 模块，您可以将模块路径视为基本上是项目的规范名称或标识符。

​	您几乎可以选择任何字符串作为您的模块路径，但重要的是要专注于唯一性。为了避免未来与其他人的项目或标准库可能发生的导入冲突，您需要选择一个全局唯一且不太可能被其他东西使用的模块路径。在 Go 社区中，常见的约定是将您的模块路径基于您拥有的 URL。

​	在我的例子中，这个项目的一个清晰、简洁且不太可能被其他任何东西使用的模块路径是 snippetbox.alexedwards.net，我将在本书的其余部分使用它。如果可能的话，您应该将其换成您独有的东西。

​	现在我们已经确定了一个唯一的模块路径，接下来我们需要做的就是将我们的项目目录变成一个模块。

​	确保您位于目录的根目录中，然后运行 go mod init命令——将您的模块路径作为参数传递，如下所示：

```bash
$ cd $HOME/code/snippetbox
$ go mod edit -module github.com/<you-github-name>/snippetbox    
```

此时您的项目目录应该看起来有点像下面的屏幕截图。注意到已经创建的 go.mod文件了吗？

![image-20260625173521167](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260625173521167.png)

​	目前这个文件中没有太多内容，如果你在你的文本编辑器中打开它，它应该看起来像这样（但最好使用你自己独特的模块路径）：

```
File: go.mod
```

```go
module github.com/<you-github-name>/snippetbox  

go 1.26.1
```

​	我们将在本书后面更详细地讨论模块，但现在只要知道当项目目录的根目录中有一个有效的 go.mod文件时，您的项目就是一个模块就足够了。将您的项目设置为模块有很多优势——包括更容易管理第三方依赖项、避免供应链攻击，以及确保您的应用程序在未来的可重现构建。

##### Hello world!

​	在我们继续之前，让我们快速检查一下一切设置是否正确。继续在您的项目目录中创建一个新的 main.go，其中包含以下代码：

```bash
$ touch main.go
```

```
File: main.go
```

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

​	保存此文件，然后在终端中使用 go run .命令编译并执行当前目录中的代码。一切顺利，您将看到以下输出：

```bash
go run .
Hello world!
```

##### 附加信息

​	如果你正在创建一个可以被其他人和程序下载和使用的项目，那么你的模块路径最好等于代码可以从中下载的位置。

​	例如，如果你的包托管在 https://github.com/foo/bar 上，则项目的模块路径应该是 github.com/foo/bar。



#### 2.3. 网络应用基础

---

​	现在一切都已正确设置，让我们对我们的 Web 应用程序进行第一次迭代。我们将从三个绝对要素开始：

​	我们首先需要的是处理程序。如果您来自 MVC 背景，您可以将处理程序视为有点像控制器。他们负责执行您的应用程序逻辑并编写 HTTP 响应标头和正文。

​	第二个组件是路由器（或 Go 术语中的 servemux）。这会存储应用程序的 URL 模式与相应处理程序之间的映射。通常，您的应用程序有一个包含所有路由的 servemux。

​	我们最不需要的是网络服务器。 Go 的一大优点是您可以建立一个 Web 服务器并监听传入的请求，作为应用程序本身的一部分。您不需要像 Nginx 或 Apache 这样的外部第三方服务器。

​	让我们将这些组件放在 main.go文件中以制作一个可用的应用程序。

```go
package main

import (
	"log"
	"net/http"
)

// 定义一个 home 处理函数，它写入一个字节片
// 其中包含来自 Snippetbox 的 Hello 作为响应体
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}
func main() {
	// 使用 http.NewServeMux() 函数初始化一个新的 servemux
	// 然后将 home 函数注册为 URL 模式的处理程序
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	
	// 使用 http.ListenAndServe() 函数启动一个新的 web 服务器
	// 我们传入两个参数:要侦听的 TCP 网络地址(在本例中为 4000)和我们刚刚创建的 servemux
	// 如果 http.ListenAndServe() 返回一个错误,我们使用 log.Fatal() 函数记录错误消息并退出
	// 注意, ttp.ListenAndServe() 返回的任何错误总是非nil
	log.Println("Starting server :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
```

> [!NOTe]
>
> ​	home 处理函数只是一个带有两个参数的常规 Go 函数。http.ResponseWriter 参数提供了组装 HTTP 响应并将其发送给用户的方法，而 *http.Request 参数是指向结构体的指针，该结构体保存有关当前请求的信息（例如 HTTP 方法和正在请求的 URL）。我们将在本书中逐渐讲解这些参数，并演示如何使用它们。

​	当你运行这段代码时，它应该会启动一个 Web 服务器，监听你本地机器的 4000 端口。每次服务器收到一个新的 HTTP 请求时，它会将请求传递给 servemux，然后 servemux 将检查 URL 路径并将请求分派给匹配的处理程序。

​	让我们试一试。保存您的 main.go文件，然后尝试使用 go run命令从您的终端运行它。

```bash
PS D:\code\snippetbox> go run .
2026/06/25 18:02:10 Starting server :4000
```

​	在服务器运行时，打开网络浏览器并尝试访问 http://localhost:4000。如果一切都按计划进行，您应该会看到一个如下所示的页面：

![image-20260625180324048](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260625180324048.png)

> [!CAUTION]
>
> ​	在我们继续之前，我应该解释一下 Go 的 servemux 将 URL 模式 "/" 视为万能匹配。因此，目前我们服务器上的所有 HTTP 请求都将由 home 函数处理，而不管它们的 URL 路径如何。例如，你可以访问不同的URL路径，如 http://localhost:4000/foo ，但你将收到完全相同的响应。

​	如果您返回终端窗口，可以通过按键盘上的 Ctrl+c来停止服务器。

##### 网络地址

​	您传递给 http.ListenAndServe() 的 TCP 网络地址应采用 "host:port" 格式。如果您省略主机（就像我们对 ":4000" 所做的那样），那么服务器将监听您计算机的所有可用网络接口。通常，如果您的计算机有多个网络接口并且您只想监听其中一个，则只需在地址中指定一个主机即可。

​	在其他 Go 项目或文档中，您有时可能会看到使用 ":http" 或 ":http-alt" 之类的命名端口而不是数字编写的网络地址。如果您使用命名端口，那么 Go 将在启动服务器时尝试从您的 /etc/services文件中查找相关端口号，或者如果找不到匹配项将返回错误。

##### 使用 Go 运行

​	在开发过程中，go run 命令是测试代码的便捷方式。它本质上是一种编译代码的快捷方式，在 /tmp目录中创建可执行二进制文件，然后一步运行该二进制文件。

​	它接受一个以空格分隔的 .go 文件列表、特定包的路径（其中.符号表示当前目录）或完整的模块路径。对于我们的应用程序，目前以下三个命令是等效的：

```bash
$ go run main.go
$ go run .
$ go run github.com/<you-github-name>/snippetbox
```



#### 2.4. 路由请求

---

​	只有一个路由的 Web 应用程序不是很令人兴奋……或有用！让我们添加更多的路由，以便应用程序开始像这样成形：

| URL Pattern     | Handler       | Action                                   |
| --------------- | ------------- | ---------------------------------------- |
| /               | home          | Display the home page(显示主页)          |
| /snippet        | showSnippet   | Display a specific snippet(显示特定片段) |
| /snippet/create | createSnippet | Create a new snippet(创建一个新的片段)   |

重新打开 main.go 文件并更新如下：

```
File: main.go
```

```go
package main

import (
	"log"
	"net/http"
)

// 定义一个 home 处理函数，它写入一个字节片
// 其中包含来自 Snippetbox 的 Hello 作为响应体
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// 添加一个 showSnippet 处理函数
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// 添加 createSnippet 处理函数
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// 向 servemux 注册两个新的处理程序函数和相应的URL模式
	// 与我们之前所做的完全相同
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
```

确保保存这些更改，然后重新启动 Web 应用程序：

```bash
PS D:\code\snippetbox> go run main.go      
2026/06/25 18:29:03 Starting server on :4000
```

如果您在 Web 浏览器中访问以下链接，您现在应该会为每条路线获得适当的响应：

http://localhost:4000/snippet

![image-20260625183102945](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260625183102945.png)

http://localhost:4000/snippet/create

![image-20260625183240691](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260625183240691.png)

##### 固定路径和子树模式

现在两条新路线已经启动并开始运行，让我们谈谈一些理论。

Go 的 servemux 支持两种不同类型的 URL 模式：固定路径和子树路径。固定路径不以尾部斜杠结尾，而子树路径确实以尾部斜杠结尾。

​	我们的两个新模式——"/snippet" 和 "/snippet/create" —— 都是固定路径的例子。在 Go 的 servemux 中，只有当请求 URL 路径与固定路径完全匹配时，才会匹配（并调用相应的处理程序）这样的固定路径模式。

​	相反，我们的模式 "/" 是子树路径的示例（因为它以斜杠结尾）。另一个例子是 "/static/"。只要请求 URL 路径的开头与子树路径匹配，就会匹配子树路径模式（并调用相应的处理程序）。如果它有助于您的理解，您可以认为子树路径的行为有点像它们末尾有一个通配符，例如 **"/****" 或 "/static/**"。

这有助于解释为什么 "/"模式表现得像一个万能的。该模式本质上意味着匹配一个斜线，然后是任何东西（或什么都没有）。

##### 限制根 url 模式

那么，如果您不希望 "/"模式像包罗万象呢？

例如，在我们正在构建的应用程序中，当且仅当请求 URL 路径与 "/" 完全匹配时，我们希望显示主页。否则，我们希望用户收到 404 page not found 响应。

不可能更改 Go 的 servemux 的行为来执行此操作，但您可以在 home 处理程序中包含一个简单的检查，最终具有相同的效果：
```File: main.go
File: main.go
```

```go
package main

import (
	"log"
	"net/http"
)

// 定义一个 home 处理函数，它写入一个字节片
// 其中包含来自 Snippetbox 的 Hello 作为响应体
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

...
```

继续进行更改，然后重新启动服务器并请求未注册的 URL 路径，例如 http://localhost:4000/missing。你应该得到一个 404的响应，看起来有点像这样：

![image-20260625184504655](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260625184504655.png)

如果您已经使用Go一段时间，可能已经遇到过 http.Handle() 和 http.HandleFunc() 函数。这些函数允许您在不声明 servemux 的情况下注册路由，如下所示：

```go
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet", showSnippet)
	http.HandleFunc("/snippet/create", createSnippet)
    
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
```

​	在幕后，这些函数将它们的路由与一个叫做 DefaultServeMux 的东西注册关联起来。这没什么特别之处——它只是我们已经使用过的常规 servemux，但其默认初始化并存储在 net/http 全局变量中。下面是Go源代码中相关的一行：

```go
var DefaultServeMux = NewServeMux()
```

尽管这种方法可以使您的代码稍微短一些，但我不建议将其用于生产应用程序。

​	因为 DefaultServeMux 是一个全局变量，任何包都可以访问它并注册路由——包括您的应用程序导入的任何第三方包。如果其中一个第三方包被攻击，它们可以使用 DefaultServeMux 将恶意处理程序公开到网络中。

​	因此，出于安全考虑，通常最好避免使用 DefaultServeMux 和相应的辅助函数。改为使用自己的本地范围servemux，就像我们迄今在这个项目中一直在做的那样。

##### 附加信息

**Servemux 特性和特点**

​	在 Go 的 servemux 中，较长的 URL 模式总是优先于较短的。因此，如果 servemux 包含多个匹配请求的模式，它总是会将请求分派给最长模式对应的处理程序。这有一个很好的副作用，您可以按任何顺序注册模式，并且不会改变 servemux 的行为方式。

​	请求 URL 路径会自动清理。如果请求路径包含任何或元素或重复的斜杠，用户将自动重定向到等效的干净 URL。例如，如果用户向 /foo/bar/..//baz 发出请求，他们将自动向 /foo/baz 发送 301 Permanent Redirect。

​	如果一个子树路径已经被注册并且收到了一个没有尾部斜杠的子树路径请求，那么用户将自动发送一个 301 Permanent Redirect 到带有斜杠的子树路径。例如，如果您注册了子树路径 /foo/，那么任何对 /foo的请求都将被重定向到 /foo/。

**主机名匹配**

​	可以在您的 URL 模式中包含主机名。当您希望将所有 HTTP 请求重定向到规范 URL，或者如果您的应用程序充当多个站点或服务的后端时，这会很有用。例如：

```go
mux := http.NewServeMux()
mux.HandleFunc("foo.example.org/", fooHandler)
mux.HandleFunc("bar.example.org/", barHandler)
mux.HandleFunc("/baz", bazHandler)
```

​	当谈到模式匹配时，将首先检查任何特定于主机的模式，如果匹配，请求将被分派到相应的处理程序。只有当没有找到特定于主机的匹配项时，才会检查非特定于主机的模式。

**RESTful 路由呢？**

​	必须承认，Go 的 servemux 提供的路由功能相当轻量级。它不支持基于请求方法的路由，不支持带有变量的干净URL，并且不支持基于正则表达式的模式。如果您曾经使用过Rails、Django 或 Laravel 等框架，可能会觉得有些受限制了...甚至感到惊讶！

​	但是不要让那些让你失望。事实上，Go 的 servemux 仍然可以让你走得更远，并且对于许多应用程序来说已经足够了。在您需要更多的时候，可以使用大量的第三方路由器来代替 Go 的 servemux。我们将在本书后面介绍一些流行的选项。

#### 2.5. 自定义 HTTP 标头

#### 2.6. 网址查询字符串

#### 2.7. 项目结构和组织

#### 2.8. HTML 模板和继承

#### 2.9. 提供静态文件

#### 2.10. http.Handler 接口

### 3. 配置和错误处理

#### 3.1. 管理配置设置

#### 3.2. 分级记录

#### 3.3. 依赖注入

#### 3.4. 集中错误处理

#### 3.5. 隔离应用程序路由

### 4. 数据库驱动的响应

#### 4.1. 设置 MySQL

#### 4.2. 安装数据库驱动程序

#### 4.3. 模块和可重现的构建

#### 4.4. 创建数据库连接池

#### 4.5. 设计数据库模型

#### 4.6. 执行 SQL 语句

#### 4.7. 单记录 SQL 查询

#### 4.8. 多记录 SQL 查询

#### 4.9. 交易和其他细节

### 5. 动态 HTML 模板

#### 5.1. 显示动态数据

#### 5.2. 模板操作和函数

#### 5.3. 缓存模板

#### 5.4. 捕获运行时错误

#### 5.5. 常用动态数据

#### 5.6. 创建数据库连接池

### 6. 中间件

#### 6.1.中间件如何工作

#### 6.2. 设置安全标头

#### 6.3. 请求记录

#### 6.4. 紧急恢复

#### 6.5. 可组合的中间件链

### 7. 进阶路由

#### 7.1. 选择路由器

#### 7.2. 干净的 URL 和基于方法的路由

### 8. 加工表格

#### 8.1. 设置 HTML 表单

#### 8.2.  解析表单数据

#### 8.3. 验证表单数据

#### 8.4. 显示错误并重新填充字段

#### 8.5. 创建验证助手

#### 8.6. 自动表单解析

### 9. 有状态的HTTP

#### 9.1. 选择会话管理器

#### 9.2. 设置会话管理器

#### 9.3. 使用会话数据

### 10. 安全改进

#### 10.1. 生成自签名 TLS 证书

#### 10.2. 运行 HTTPS 服务器

#### 10.3. 配置 HTTPS 设置

#### 10.4. 连接超时

### 11. 用户认证

#### 11.1. 路线设置

#### 11.2. 创建用户模型

#### 11.3. 用户注册和密码加密

#### 11.4. 用户登录

#### 11.5. 用户注销

#### 11.6. 用户授权

#### 11.7. CSRF保护

### 12. 使用请求上下文

#### 12.1. 请求上下文如何工作

#### 12.2. 身份验证/授权的请求上下文

### 13. 可选的 Go 功能

#### 13.1. 使用嵌入文件

#### 13.2. 使用泛型

### 14. 测试

#### 14.1. 单元测试和子测试

#### 14.2. 测试 HTTP 处理程序和中间件

#### 14.3. 端到端测试

#### 14.4. 自定义测试的运行方式

#### 14.5. 模拟依赖

#### 14.6. 测试 HTML 表单

#### 14.7. 集成测试

#### 14.8. 分析测试覆盖率

### 15. 结论

### 16. 进一步阅读和有用的链接

### 17. 指导练习

#### 17.1. 向应用程序添加“关于”页面

#### 17.2. 添加调试模式

#### 17.3. 测试 snippetCreate 处理程序

#### 17.4. 向应用程序添加“帐户”页面

#### 17.5. 登录后适当重定向用户

#### 17.6. 实施“更改密码”功能



