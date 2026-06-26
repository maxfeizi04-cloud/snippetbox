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



#### 2.5. 自定义 HTTP 请求头

---

现在让我们更新我们的应用程序，使 /snippet/create路由仅响应使用 POST方法的 HTTP 请求，如下所示：

| Method | Pattern         | Handler       | Action                     |
| ------ | --------------- | ------------- | -------------------------- |
| ANY    | /               | home          | Display the home page      |
| ANY    | /snippet        | showSnippet   | Display a specific snippet |
| POST   | /snippet/create | createSnippet | Create a new snippet       |

​	进行这种更改非常重要，因为在应用程序构建的后期，对/snippet/create路由的请求将导致在数据库中创建一个新的片段。在数据库中创建新片段是一项非幂等性操作，它会改变我们服务器的状态，因此我们应该遵循HTTP的良好实践，并将此路由限制为仅以POST请求执行。

​	不过，我现在要专门讲这个，主要是因为它是引出 HTTP 响应头话题的好契机，顺便也能说明如何定制这些响应头。

##### HTTP 状态代码

​	首先更新 `snippetCreate` 处理函数，使其在请求方法非 POST 时返回 **405（Method Not Allowed）** 状态码。为此需使用 `w.WriteHeader()` 方法，示例如下：

```
File: main.go
```

```go
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// 使用 r.Method 检查请求是否使用 POST 方法。注意，
	// http.MethodPost 是一个常量，其值为字符串 "POST"。
	if r.Method != http.MethodPost {
		// 若不是 POST，则调用 w.WriteHeader() 发送 405 状态码，
		// 并调用 w.Write() 写入 "Method Not Allowed" 作为响应体。
		// 随后从函数返回，使后续代码不会被执行。
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
```

​	尽管此更改看起来很简单，但我应该解释一些细微差别：

​	每次响应只能调用 w.WriteHeader() 一次，状态码写入后便无法更改。如果您尝试第二次调用 w.WriteHeader()，Go 将记录一条警告消息。

​	如果您没有显式调用 w.WriteHeader()，那么第一次调用 w.Write()会自动向用户发送 200 OK状态码。因此，如果您想发送非 200 状态代码，则必须在调用 w.Write() 之前调用 w.WriteHeader()。

让我们看一下实际效果。

​	重新启动服务器，然后打开第二个终端窗口并使用 curl 向 http://localhost:4000/snippet/create发出 POST请求。您应该会收到一个带有 200 OK状态代码的 HTTP 响应，类似于：

```bash
$ curl -i -X POST http://localhost:4000/snippet/create
HTTP/1.1 200 OK
Date: Fri, 26 Jun 2026 10:02:38 GMT
Content-Length: 23
Content-Type: text/plain; charset=utf-8

Create a new snippet...
```

​	但是，如果您使用不同的请求方法——例如 GET、PUT或 DELETE——您现在应该会收到带有 405 Method Not Allowed状态代码的响应。例如：

```bash
$ curl -i -X PUT http://localhost:4000/snippet/create
HTTP/1.1 405 Method Not Allowed
Date: Fri, 26 Jun 2026 10:04:32 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

Method not allowed
```

##### 自定义标题

​	我们可以做的另一项改进是在 405 Method Not Allowed 响应中包含一个 Allow 标头，让用户知道该特定 URL 支持哪些请求方法。
​	我们可以通过使用 w.Header().Set() 方法将新标头添加到响应头映射来完成此操作，如下所示：

```bash
File: mian.go
```

```go
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// 使用 r.Method 检查请求是否使用 POST 方法。注意，
	// http.MethodPost 是一个常量，其值为字符串 "POST"。
	if r.Method != http.MethodPost {
		// 若不是 POST，则调用 w.WriteHeader() 发送 405 状态码，
		// 并调用 w.Write() 写入 "Method Not Allowed" 作为响应体。
		// 随后从函数返回，使后续代码不会被执行。
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
```

> [!IMPORTANT]
>
> 对响应标头映射的更改，若发生在 `w.WriteHeader()` 或 `w.Write()` 调用之后，则对用户实际接收到的标头无效。因此，必须在调用这些方法之前，确保标头映射已包含全部所需标头。

再次向 `/snippet/create` 发送非 POST 请求以观察实际效果，示例如下：

```bash
$ curl -i -X PUT http://localhost:4000/snippet/create
HTTP/1.1 405 Method Not Allowed
Allow: POST
Date: Fri, 26 Jun 2026 10:11:51 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

Method not allowed
```

注意响应现在包含了新的 `Allow: POST` 标头

##### http.Error 快捷方式

​	如果你想发送一个非200状态码和纯文本响应体（就像我们在上面的代码中所做的那样），那么这是使用http.Error() 快捷方式的好机会。这是一个轻量级的辅助函数，它接受给定的消息和状态码，然后在幕后为我们调用 w.WriteHeader() 和 w.Write() 方法。

让我们更新代码以使用它。

```
File: main.go
```

```go
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// 使用 r.Method 检查请求是否使用 POST 方法。注意，
	// http.MethodPost 是一个常量，其值为字符串 "POST"。
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// 使用 http.Error() 函数发送 405 状态码
		// 并以 "Method Not Allowed" 字符串作为响应体
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
```

​	在功能方面，这几乎是完全相同的。最大的区别是我们现在将我们的http.ResponseWriter 传递给另一个函数，它为我们向用户发送响应。

​	将 http.ResponseWriter 传递给其他函数的模式在 Go 语言中非常常见，而且在本书中我们将经常使用它。在实践中，直接使用 w.Write() 和 w.WriteHeader() 方法像我们迄今所做的那样是相当罕见的。但我想提前介绍它们，因为它们是发送响应的更高级和有趣,方法的基础。

##### net/http 常量

​	我们可以做的最后一项调整是使用 net/http 包中的常量作为 HTTP 方法和状态代码，而不是我们自己编写字符串和整数。

​	具体来说，我们可以使用常量 http.MethodPost代替字符串 "POST"，使用常量 http.StatusMethodNotAllowed 代替整数 405。像这样：

```
File: main.go
```

```go
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// 使用 r.Method 检查请求是否使用 POST 方法。注意，
	// http.MethodPost 是一个常量，其值为字符串 "POST"。
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// 使用 http.Error() 函数发送 405 状态码
		// 并以 "Method Not Allowed" 字符串作为响应体
        // http.StatusMethodNotAllowed = 405
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) 
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
```

​	使用这些常量是一种很好的做法，因为它有助于防止因拼写错误而导致的运行时错误，并且在 HTTP 状态代码常量的情况下，它还可以帮助使您的代码更清晰和自我记录——尤其是在处理不常用的状态时代码。

##### 附加信息

**系统生成的标头和内容嗅探**

​	发送响应时，Go 会自动为您设置三个系统生成的标头：Date 和 Content-Length 以及 Content-Type。

​	`Content-Type` 标头尤其值得关注。Go 会尝试通过 `http.DetectContentType()` 函数对响应体进行内容嗅探，自动设置正确的值。若该函数无法推断出内容类型，Go 会将 `Content-Type` 标头回退设置为 `application/octet-stream`。

​	`http.DetectContentType()` 函数通常运行良好，但对于刚接触 Go 的 Web 开发者来说，一个常见的陷阱是它无法区分 JSON 与纯文本。因此，JSON 响应默认会带有 `Content-Type: text/plain; charset=utf-8` 标头发送。你可以通过手动设置正确的标头来避免此问题，如下所示：

```go
w.Heander().Set("Content-Type", "application/json")
w.Write([]byte(`{"name": "Alex"}`))
```

##### 操作标头映射

​	上述代码中使用 `w.Header().Set()` 向响应标头映射添加了新标头。此外，还可以使用 `Add()`、`Del()`、`Get()` 和 `Values()` 方法来读取和操作标头映射。

```go
// 设置一个新的 Cache-Control 标头。如果已存在 "Cache-Control" 标头，它将被覆盖。
w.Header().Set("Cache-Control", "public, max-age=31536000")

// 相比之下，Add() 方法会追加一个新的 "Cache-Control" 标头，并且可以被多次调用。
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=31536000")

// 删除 "Cache-Control" 标头的所有值。
w.Header().Del("Cache-Control")

// 获取 "Cache-Control" 标头的第一个值。
w.Header().Get("Cache-Control")
```

**标头规范化**

​	对标头映射使用 `Set()`、`Add()`、`Del()`、`Get()` 和 `Values()` 方法时，标头名称始终会通过 `textproto.CanonicalMIMEHeaderKey()` 函数进行规范化。该函数将首字母及连字符后的字母转为大写，其余字母转为小写。实际影响是，调用这些方法时标头名称不区分大小写。

​	如需避免此规范化行为，可直接编辑底层的标头映射（其类型为 `map[string][]string`）。例如：

```go
w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
```

> [!note]
>
> 若使用 HTTP/2 连接，Go 会始终根据 HTTP/2 规范自动将标头名称和值转换为小写。

**抑制系统生成的标头**

​	`Del()` 方法不会移除系统生成的标头。若要移除这些标头，需直接访问底层标头映射并将其值设为 `nil`。例如，若要移除 `Date` 标头，需写入：

```go
w.Header()["Date"] = nil
```



#### 2.6. 网址查询字符串

---

##### 网址查询字符串

​	趁现在讲到路由，我们来更新 `snippetView` 处理函数，使其接受用户传入的 `id` 查询字符串参数，如下所示：

| Method | Pattern         | Handler       | Action                     |
| ------ | --------------- | ------------- | -------------------------- |
| ANY    | /               | jhome         | Display the home page      |
| ANY    | /snippet?id=1   | showSnippet   | Dispaly a specific snippet |
| POST   | /snippet/create | createSnippet | Create a new snippet       |

​	稍后我们会用这个 `id` 参数从数据库中查询对应的 snippet 并展示给用户。但目前，我们只读取 `id` 参数的值，并将其拼接到占位响应中。

为实现该功能，需要更新 `showSnippet` 处理函数，完成两件事：

1. 从 URL 查询字符串中获取 `id` 参数的值，可使用 `r.URL.Query().Get()` 方法。该方法始终返回参数的字符串值，若不存在匹配参数则返回空字符串 `""`。
2. 由于 `id` 参数来自不可信的用户输入，需对其进行验证，确保合理有效。就本 Snippetbox 应用而言，需检查其是否为正整数值。可通过 `strconv.Atoi()` 函数将字符串值转换为整数，然后检查该值是否大于零。

就是这样：

```
File: main.go
```

```go
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// 从查询字符串中提取 id 参数的值，并尝试使用 strconv.Atoi() 函数将其转换为整数。
	// 如果无法转换为整数，或值小于 1，则返回 404 页面未找到响应。
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// 使用 fmt.Fprintf() 函数将 id 值拼接到响应中，并写入 http.ResponseWriter。
	fmt.Fprintf(w, "Display a specific with ID %d...", id)
}
```

我们来试一下。

重启应用，然后访问类似 `http://localhost:4000/snippet?id=123` 的 URL。你应该会看到如下响应：

![image-20260626184934338](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260626184934338.png)

您可能还想尝试访问一些包含 id参数无效值或根本没有参数的 URL。例如：

- http://localhost:4000/snippet
- http://localhost:4000/snippet?id=-1
- http://localhost:4000/snippet?id=foo

对于所有这些请求，您应该得到 404 page not found响应。

##### io.writer 接口

​	上述代码在底层还引入了一个新知识点。查看 `fmt.Fprintf()` 函数的文档会发现，它的第一个参数是 `io.Writer` 类型……

```go
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
```

​	但我们传入的是 `http.ResponseWriter` 对象，并且运行正常。

​	之所以如此，是因为 `io.Writer` 是一个接口类型，而 `http.ResponseWriter` 对象实现了该接口（因为它有 `Write()` 方法）。如果你是 Go 新手，接口的概念可能会有些令人困惑，现在我不想深入展开。你只需知道，在实践中，凡是看到 `io.Writer` 类型参数的地方，都可以传入 `http.ResponseWriter` 对象。写入的内容随后会作为 HTTP 响应体发送出去。



#### 2.7. 项目结构和组织

---

​	在继续向 `main.go` 添加更多代码之前，现在是时候考虑如何组织和构建这个项目了。

​	需要事先说明的是，Go 中并没有唯一正确——甚至推荐——的 Web 应用结构方式。这既是好事也是坏事。一方面，你可以自由灵活地组织代码；另一方面，在决定最佳结构时又很容易陷入不确定的困境。

​	随着 Go 使用经验的积累，你会逐渐形成一套在不同场景下行之有效的模式。但作为起点，我能给的最好建议就是：不要过度复杂化。尽量只在确实需要时才引入结构和复杂性。

​	对于本项目，我们将采用一种流行且久经考验的结构方案。这是一个坚实的起点，其通用结构可以复用到多种项目中。

请确保当前位于项目仓库根目录下，然后执行以下命令：

```bash
// windows
$ cd /code/snippetbox
$ rm main.go
$ mkdir -p cmd/web,pkg,ui/html,ui/static
$ type nul > cmd/web/main.go
$ type nul > cmd/web/handlers.go

// Linux/Unix
$ cd $HOME/code/snippetbox
$ rm main.go
$ mkdir -p cmd/web pkg ui/html ui/static
$ touch cmd/web/main.go
$ touch cmd/web/handlers.go
```

您的项目存储库的结构现在应该如下所示：

![image-20260626190656393](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260626190656393.png)



/docs 目录是我自己的文档目录当**不存在**。

我们花点时间来讨论一下每个目录的用途。

​	`cmd` 目录将存放项目中可执行程序的应用特定代码。目前我们只有一个可执行程序——Web 应用——它将位于 `cmd/web` 目录下。

​	`pkg` 目录将存放项目中非应用特定的辅助代码。我们将用它来存放可能可重用的代码，例如验证辅助函数和项目的 SQL 数据库模型。

​	`ui` 目录将存放 Web 应用所使用的用户界面资源。具体来说，`ui/html` 目录存放 HTML 模板，`ui/static` 目录存放静态文件（如 CSS 和图片）。

那么为什么要采用这种结构呢？主要有两大好处：

1. 它清晰地分离了 Go 代码和非 Go 资源。我们编写的所有 Go 代码将专属于 `cmd` 和 `pkg` 目录之下，项目根目录则留作存放非 Go 资源，如 UI 文件、Makefile 和模块定义（包括 `go.mod` 文件）。这在将来构建和部署应用时有助于简化管理。
2. 如果将来想在项目中添加另一个可执行程序，这种结构扩展性很好。例如，将来你可能想添加一个 CLI（命令行界面）来自动化一些管理任务。在这种结构下，你可以将 CLI 应用放在 `cmd/cli` 下，并且它能够导入和复用你在 `pkg` 目录下编写的所有代码。

##### 重构现有代码

我们快速将已编写的代码迁移到新结构中。

```go
File: main.go
```

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
		
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
```

```
File: cmd/web/handlers.go
```

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
```

现在我们的 Web 应用由 `cmd/web` 目录下的多个 Go 源文件组成。要运行它们，可以使用 `go run` 命令，如下所示：

```bash
$ go run ./cmd/web        
2026/06/26 19:20:46 Starting server on :4000
```

##### 附加信息

**内部目录**

​	需要指出的是，目录名称 internal 在 Go 中具有特殊的意义和行为：存放在这个目录下的任何包只能被父级 internal 目录的内部代码导入。在我们的情况下，这意味着任何存放在 internal 中的包都只能被我们的 snippetbox 项目目录中的代码导入。

​	或者换句话说，这意味着任何 internal 目录下的包都不能被我们项目之外的代码导入。

​	这很有用，因为它可以防止其他代码库导入和依赖于我们的internal目录中的（可能是未经版本化和不受支持的）包 - 即使项目代码在GitHub等公共位置上可用。



#### 2.8. HTML 模板和继承

---

​	让我们为项目注入一点活力，并为我们的 Snippetbox Web 应用程序开发一个合适的主页。在接下来的几章中，我们将努力创建一个如下所示的页面：

![image-20260626192401294](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260626192401294.png)

​	让我们首先在 ui/html/pages/home.tmpl创建一个模板文件，其中包含我们主页的 HTML 内容。像这样：

```bash
$ cd ui\html\
$ type nul > home.page.tmpl
```

```
File: ui/html/home.page.tmpl
```

```html
<!doctype html>
<html lang='en'>
    <head>
        <meta charset='utf-9'>
        <title>Home - Snippetbox</title>
    </head>
    <body>
        <header>
            <h1><a href='/'>Snippetbox</a> </h1>
        </header>
        <nav>
            <a href="/">Home</a>
        </nav>
        <main>
            <h2>Latest Snippets</h2>
            <p>There's nothing to see here yet!</p>
        </main>
    </body>
</html>
```

> [!NOTE]
>
> 本书中，我们将使用 `<name>.<role>.tmpl` 的命名约定来命名模板文件，其中 `<role>` 为 `page`、`partial` 或 `layout`。通过文件名即可确定模板的角色，这有助于在本书后续内容中创建模板缓存。
>
> .tmpl 扩展在这里没有传达任何特殊含义或行为。我之所以选择这个扩展名，是因为它是一种很好的方式，可以在您浏览文件列表时清楚地表明该文件包含一个 Go 模板。但是，如果你愿意，你可以使用扩展名 .html来代替（这可能会让你的文本编辑器将文件识别为 HTML，以便语法高亮或自动完成）——或者你甚至可以使用“双重扩展名”，比如 .tmpl.html。选择权在你，但我们将在整本书中坚持使用 .tmpl 作为我们的模板。

​	现在已经创建了包含首页 HTML 标记的模板文件，下一个问题是如何让 `home` 处理函数渲染它。

​	为此，我们需要导入 Go 的 `html/template` 包，该包提供了一系列用于安全解析和渲染 HTML 模板的函数。我们可以使用该包中的函数来解析模板文件，然后执行模板。

我来演示一下。打开 `cmd/web/handlers.go` 文件，添加以下代码：

```
File: cmd/web/handlers.go
```

```go
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 使用 template.ParseFiles() 函数将模板文件读取到模板集中
	// 如果出错，记录详细错误信息并使用 http.Error() 函数向用户发送
	// 500 Internal Server Error 响应
	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 然后使用模板集的 Execute() 方法将模板内容写入响应体
	// Execute() 的最后一个参数用于传入动态数据，目前暂设为 nil
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
```

​	需要指出的是，传给 `template.ParseFiles()` 函数的文件路径必须是相对于当前工作目录的路径，或绝对路径。上述代码中，我使用的是相对于项目根目录的路径。

因此，请确保当前位于项目根目录下，然后重启应用：

```bash
$ go run ./cmd/web
2026/06/26 19:53:37 Starting server on :4000
```

然后在浏览器中打开 `http://localhost:4000`，你应该会看到首页的 HTML 页面已经基本成形。

![image-20260626195602954](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260626195602954.png)

##### 模板组成

​	随着我们为 Web 应用添加更多页面，会有一些共享的样板 HTML 标记需要包含在每个页面中——例如 `<head>` HTML 元素中的页头、导航和元数据。

​	为了避免重复输入，最好创建一个包含这些共享内容的布局（或主）模板，然后将其与各个页面的专属标记组合使用。

​	现在创建 `ui/html/base.layout.tmpl` 文件：

```bash
$ type nul > ui/html/base.layout.tmpl
```

并添加以下标记（我们希望出现在每个页面上）：

```
File: ui/html/base.layout.tmpl
```

```html
{{define "base"}}
<!doctype html>
<html lang='en'>
    <head>
        <meta charset='utf-8'>
        <title>{{template "title" .}} - Snippetbox</title>
    </head>
    <body>
        <header>
            <h1><a href='/'>Snippetbox</a> </h1>
        </header>
        <nav>
            <a href='/'>Home</a>
        </nav>
        <main>
            {{template "main" .}}
        </main>
    </body>
</html>
{{end}}
```

​	如果你之前用过其他语言的模板，应该会觉得这很熟悉。本质上它只是常规的 HTML，加上一些用双花括号包裹的额外动作。

​	这里我们使用 `{{define "base"}}...{{end}}` 动作定义了一个名为 `base` 的命名模板，其中包含我们希望在每个页面中显示的内容。

​	在其中，我们使用 `{{template "title" .}}` 和 `{{template "main" .}}` 动作来表示我们希望在 HTML 的特定位置调用其他命名模板（分别名为 `title` 和 `main`）。

> [!NOTE]
>
> 如果你有疑问，`{{template "title" .}}` 动作末尾的点表示要传递给被调用模板的任何动态数据。我们稍后会讨论。

现在回到 `ui/html/home.page.tmpl`，将其更新为定义 `title` 和 `main` 命名模板，包含首页的特定内容：

```
File: ui/html/home.page.tmpl
```

```html
{{template "base" .}}

{{define "title"}}Home{{end}}

{{define "main"}}
<h2>Latest Snippet</h2>
<p>There's nothing to see here yet!</p>
{{end}}
```

​	该文件顶部可以说是最重要的部分——`{{template "base" .}}` 动作。它告诉 Go，在执行 `home.page.tmpl` 文件时，我们要调用名为 `base` 的模板。

​	而 `base` 模板又包含调用 `title` 和 `main` 命名模板的指令。我知道一开始这可能感觉有点循环，但请耐心——在实践中，这种模式效果非常好。

​	完成后，下一步是更新 `home` 处理函数中的代码，使其同时解析两个模板文件，如下所示：

```go
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 初始化一个包含两个文件路径的切片
	// 注意，home.page.tmpl 文件必须是切片中的 *第一个* 文件。
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	
	// 使用 template.ParseFiles() 函数将模板文件读取到模板集中
	// 如果出错，记录详细错误信息并使用 http.Error() 函数向用户发送
	// 500 Internal Server Error 响应
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 然后使用模板集的 Execute() 方法将模板内容写入响应体
	// Execute() 的最后一个参数用于传入动态数据，目前暂设为 nil
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
```

​	现在，模板集不再直接包含 HTML，而是包含三个命名模板（`base`、`title` 和 `main`），以及一个调用 `base` 模板的指令（`base` 模板又会调用 `title` 和 `main` 模板）。

​	请随意重启服务器并尝试一下。你应该会看到输出与之前相同（只是 HTML 源码中动作所在位置会多出一些空白）。

​	使用这种模式组合模板的一大好处是，你可以将页面专属内容清晰地在磁盘上按单独文件定义，并在这些文件中控制页面使用哪种布局模板。这对于大型应用尤其有用，因为应用中的不同页面可能需要使用不同的布局。

##### 嵌入部份

​	对于某些应用，你可能希望将部分 HTML 片段提取为局部模板，以便在不同的页面或布局中复用。为了演示，我们为 Web 应用创建一个包含页脚内容的局部模板。

​	新建 `ui/html/footer.partial.tmpl` 文件，并添加一个名为 `footer` 的命名模板，如下所示：

```bash
$ type nul > ui/html/footer.partial.tmpl
```

```
File: ui/html/footer.partial.tmpl
```

```html
{{define "footer"}}
<footer>Powered by <a href="https://baidu.com/">Go</a> </footer>
{{end}}
```

​	然后更新 `base` 模板，使用 `{{template "footer" .}}` 动作调用 `footer`：

​	最后，需要更新 `home` 处理函数，在解析模板文件时加入新的 `ui/html/footer.partial.tmpl` 文件：

```
File: cmd/web/handlers.go
```

```go
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 初始化一个包含两个文件路径的切片
	// 注意，home.page.tmpl 文件必须是切片中的 *第一个* 文件。
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.page.tmpl",
	}

	// 使用 template.ParseFiles() 函数将模板文件读取到模板集中
	// 如果出错，记录详细错误信息并使用 http.Error() 函数向用户发送
	// 500 Internal Server Error 响应
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 然后使用模板集的 Execute() 方法将模板内容写入响应体
	// Execute() 的最后一个参数用于传入动态数据，目前暂设为 nil
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
```

​	重启服务器后，`base` 模板现在会调用 `footer` 模板，首页应该如下所示：

![image-20260626203458260](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260626203458260.png)

##### 附加信息

**块动作**

​	上面的代码中，我们使用 `{{template}}` 动作从一个模板调用另一个模板。但 Go 还提供了 `{{block}}...{{end}}` 动作，同样可以用于此目的。

​	它的行为类似于 `{{template}}` 动作，区别在于，如果被调用的模板在当前模板集中不存在，它会允许你指定一些默认内容。

​	在 Web 应用的上下文中，当你希望提供一些默认内容（如侧边栏），而各个页面在需要时可以在个案基础上覆盖它们时，这会很有用。

语法上使用方式如下：

```bash
{{define "base"}}
	<h1>An example template</h1>
	{{block "sidebar" .}}
		<p>My default sidebar content</p>
	{{end}}
{{end}}
```

​	但是——如果你愿意的话——你不需要在 {{block}} 和 {{end}} 操作之间包含任何默认内容。在这种情况下，调用的模板就像它是“可选的”一样。如果该模板存在于模板集中，那么它将被渲染。但如果没有，则不会显示任何内容。



#### 2.9. 提供静态文件

---

​	现在我们来改善首页的外观，为项目添加一些静态 CSS 和图片文件，再加上一小段 JavaScript 来高亮当前导航项。

​	如果你在跟着操作，可以用以下命令获取所需文件并解压到之前创建的 `ui/static` 文件夹中：

```bash
$ cd $HOME/code/snippetbox
$ curl https://www.alexedwards.net/static/sb.v130.tar.gz | tar -xvz -C ./ui/static/
```

​	也可以直接输入 https://www.alexedwards.net/static/sb.v130.tar.gz 进行下载在解压,再移动到对应文件夹下.

​	ui/static目录的内容现在应该如下所示：

![image-20260626213935660](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260626213935660.png)

##### http.Fileserver 处理程序











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



