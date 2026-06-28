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

**背景知识**

​	这本书是为刚接触 Go 的人设计的，但如果你先对 Go 的语法有一个大致的了解，你可能会发现它更有趣。如果你发现自己在语法上苦苦挣扎，Karl Seguin 的 Little Book of Go 是一个很棒的教程，或者如果你想要更具交互性的东西，我建议你完成 Go 之旅。

​	我还假设您对 HTML/CSS 和 SQL 有（非常）基本的了解，并且对使用终端（或 Windows 用户的命令行）有一定的了解。如果您以前用任何其他语言构建过 Web 应用程序——无论是 Ruby、Python、PHP 还是 C#——那么这本书应该非常适合您。

**其他软件**

​	如果您想完全按照说明进行操作，则应确保您的计算机上还有其他一些软件可用。他们是：

​	用于处理来自终端的 HTTP 请求和响应的 curl 工具。在 MacOS 和 Linux 机器上，它应该预先安装或在您的软件存储库中可用。否则，您可以从此处下载最新版本。

​	具有良好开发工具的 Web 浏览器。我将在本书中使用 Firefox，但 Chromium、Chrome 或 Microsoft Edge 也可以。



### 2. 基础

---



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

---



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

​	Go 标准库中的 **`net/http` 包**内置了一个用于提供静态文件服务的处理器（Handler）——**`http.FileServer`**。

​	借助 `http.FileServer`，我们可以将某个指定目录中的文件通过 **HTTP 协议** 对外提供访问，而无需自己编写读取文件、设置响应头以及返回文件内容等逻辑。

​	下面，我们将在应用程序中新增一条路由，使**所有以 `"/static/"` 开头的请求**都交由 `http.FileServer` 进行处理，如下所示：

| Method | Pattern         | Handler         | Action                       |
| ------ | --------------- | --------------- | ---------------------------- |
| ANY    | /               | home            | Display the home page        |
| ANY    | /snippet?id=1   | showSnippet     | Display a specific snippet   |
| POST   | /snippet/create | createSnippet   | Create a new snippet         |
| ANY    | /static/        | http.FileServer | Serve a specific static file |

> [!NOTE]	
>
> 请注意，路由模式 **`"/static/"`** 属于一种**子树路径模式（Subtree Path Pattern）**。它的匹配规则类似于路径末尾隐式带有一个通配符，因此**所有以 `"/static/"` 开头的请求路径**都会匹配到这条路由。

​	要创建一个新的 **`http.FileServer`** 处理器（Handler），需要调用 **`http.FileServer()`** 函数，具体写法如下：

```go
fileServer := http:FileServer(http.Dir("./ui/static/"))
```

​	当 `http.FileServer` 处理器接收到请求时，它会先移除 **URL 路径开头的 `/`**，然后在 `./ui/static` 目录中查找与请求路径对应的文件，并将该文件返回给客户端。

​	因此，为了使其能够正确工作，我们必须在将请求交给 `http.FileServer` 之前，先从 URL 路径中移除前缀 `"/static"`。否则，`http.FileServer` 将会去查找一个并不存在的文件，最终向客户端返回 **404 Not Found（页面未找到）** 响应。

​	幸运的是，Go 标准库提供了专门用于完成这一任务的辅助函数 **`http.StripPrefix()`**。

​	请打开 `main.go` 文件，并添加如下代码，使文件最终内容如下所示：

```
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

	// 创建一个静态文件服务器，用于提供 "./ui/static" 目录中的文件
	// 注意，传递给 http.Dir() 的路径是相对于项目根目录（Project Root）的相对路径
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// 使用 mux.Handle() 注册路由，将 fileServer 设置为所有以 "/static/"
	// 开头的 URL 路径的处理器（Handler）。对于匹配到的请求，在交由
	// fileServer 处理之前，会先移除 URL 路径中的 "/static" 前缀
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
```

​	完成上述代码后，重新启动应用程序，并在浏览器中访问 `http://localhost:4000/static/`。

​	如果一切正常，你将看到 **`ui/static`** 目录的可浏览目录列表（Directory Listing），其效果如下图所示：

![image-20260627162236748](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260627162236748.png)

​	你可以自行浏览该目录列表，并查看其中的各个文件。

​	例如，访问 `http://localhost:4000/static/css/main.css`，浏览器将直接显示 `main.css` 文件的内容，如下图所示：

![image-20260627163215850](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260627163215850.png)

##### 使用静态文件

​	文件服务器正常工作后，我们现在可以更新 `ui/html/base.layout.tmpl` 文件以使用静态文件：

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
        <link rel='stylesheet' href='/static/css/main.css'>
        <link rel="shortcut icon" href='/static/img/favicon.ico' type="image/x-icon">
        <link rel="stylesheet" href='https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700'>
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
        <!-- Invoke the footer template -->
        {{template "footer" .}}
        <script src="/static/js/main.js" type="text/javascript"></script>
    </body>
</html>
{{end}}
```

​	保存上述修改后，访问 `http://localhost:4000`。

​	如果一切正常，首页现在应该如下图所示：

![image-20260627165026518](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260627165026518.png)

##### 附加信息

Go 的文件服务器（`http.FileServer`）提供了一些非常实用的特性，值得了解：

- 在查找文件之前，它会先通过 **`path.Clean()`** 函数对所有请求路径进行规范化处理，自动移除 URL 路径中的 `.` 和 `..` 等路径元素，从而有效防止**目录遍历攻击（Directory Traversal Attack）**。如果你使用的路由器（Router）不会自动对 URL 路径进行规范化处理，那么这一特性尤为重要。

- `http.FileServer` **完全支持 Range 请求（范围请求）**。这一特性对于提供大文件下载非常有用，因为它能够支持**断点续传**。例如，可以使用 `curl` 请求 `logo.png` 文件中 **100～199 字节**的数据，以验证这一功能，如下所示：

```bash
$ curl -i -H "Range: bytes=100-199" --output - http://localhost:4000/static/img/logo.png
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Length: 100
Content-Range: bytes 100-199/1075
Content-Type: image/png
Last-Modified: Thu, 04 May 2017 13:07:52 GMT
Date: Sat, 27 Jun 2026 08:52:57 GMT

[binary data]
```

- `http.FileServer` 会自动支持 **`Last-Modified`** 和 **`If-Modified-Since`** 请求头。如果文件自用户上次请求以来没有发生变化，`http.FileServer` 将返回 **`304 Not Modified`** 状态码，而不会再次发送文件内容。这有助于降低客户端和服务器之间的网络延迟，并减少双方的处理开销。

- `http.FileServer` 会根据文件扩展名，通过 **`mime.TypeByExtension()`** 函数自动设置 **`Content-Type`** 响应头。如有需要，还可以使用 **`mime.AddExtensionType()`** 函数为自定义文件扩展名添加对应的 MIME 类型。

**性能（Performance）**

​	在前面的代码中，我们将 `http.FileServer` 配置为从硬盘上的 **`./ui/static`** 目录提供静态文件服务。

​	不过，需要注意的是，当应用程序运行起来之后，`http.FileServer` 实际上通常并不会每次都从磁盘读取这些文件。无论是 Windows 还是类 Unix 操作系统，都会将最近访问过的文件缓存到内存（RAM）中。因此，对于经常访问的静态文件，`http.FileServer` 很可能直接从内存中读取，而无需每次都进行速度相对较慢的磁盘 I/O 操作。

 **提供单个文件（Serving Single Files）**

​	有时候，你可能只需要在某个处理器（Handler）中返回一个指定的文件。对于这种场景，可以使用 **`http.ServeFile()`** 函数，使用方式如下所示：

```go
func downloadHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./ui/static/file.zip")
}
```

> [!WARNING]
>
> **`http.ServeFile()` **不会自动对文件路径进行规范化处理。如果文件路径是根据**不可信的用户输入**构造的，为了防止**目录遍历攻击（Directory Traversal Attack）**，在调用 `http.ServeFile()` 之前，必须先使用 **`filepath.Clean()`** 对输入路径进行规范化处理。

**禁用目录列表**

​	如果你希望**禁用目录列表（Directory Listing）**，可以采用多种不同的方法。

​	其中，最简单的方法是在**需要禁用目录列表的目录**中添加一个空白的 **`index.html`** 文件。这样，当用户访问该目录时，服务器将返回这个 `index.html` 文件，而不是显示目录列表。由于该文件为空，因此客户端会收到一个 **`200 OK`** 响应，但响应体（Body）为空。

​	如果希望对 **`./ui/static`** 目录下的**所有子目录**都应用这一方式，可以执行下面的命令：

```bash
$ find ./ui/static -type d -exec touch {}/index.html \;
```

​	一种更复杂（但通常也更推荐）的解决方案是，自定义实现 **`http.FileSystem`** 接口，并在访问目录时返回 **`os.ErrNotExist`** 错误，使服务器将目录视为不存在，从而禁止目录列表的显示。

​	关于这种方案的完整原理和示例代码，可以参考作者提供的博客文章。



#### 2.10. http.Handler 接口

---

​	在继续后面的内容之前，我们需要先学习一些相关的理论知识。这部分内容稍微有些复杂，如果你觉得本章阅读起来比较吃力，也不用担心。可以先继续完成应用程序的开发，等对 Go 更加熟悉之后，再回过头来学习这一部分内容。

​	在前面的章节中，我们多次提到了 **Handler（处理器）** 这个术语，但一直没有详细说明它的真正含义。严格来说，**Handler** 是指**任何实现了 `http.Handler` 接口的对象**：

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

​	简单来说，一个对象要成为 **Handler（处理器）**，就必须实现一个**方法签名完全符合要求**的 **`ServeHTTP()`** 方法，其定义如下：

```go
ServeHTTP(http.ResponseWriter, *http.Request)
```

​	因此，一个最简单的 **Handler（处理器）** 可以写成下面这样：

```go
type home struct {}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is my home page"))
}
```

​	这里我们有一个对象（在这个例子中是一个 `home` 结构体，但它也可以是字符串、函数或其他任何类型），并且我们在该对象上实现了一个方法，其签名为 `ServeHTTP(http.ResponseWriter, *http.Request)`。这样就已经满足了 **Handler（处理器）** 的全部要求。

​	随后，你可以通过 **`servemux` 的 `Handle` 方法**将该 Handler 注册进去，如下所示：

```go
mux := http.NewServeMux()
mux.Handle("/", &home{})
```

​	当此 servemux 接收到 "/"的 HTTP 请求时，它将调用 home结构的 ServeHTTP()方法——这反过来写入 HTTP 响应。

##### Handler functions

​	不过，专门为了实现 `ServeHTTP()` 方法而去创建一个对象，这种方式显得比较冗长，也有些不够直观。因此在实际开发中，更常见的做法是像我们前面一直使用的那样，直接使用普通函数来编写 handler。例如：

```go
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
```

​	但是这个 `home` 函数只是一个普通函数，它并没有实现 `ServeHTTP()` 方法。因此从严格意义上来说，它并不是一个 **Handler（处理器）**。

​	因此，我们需要使用 **`http.HandlerFunc()` 适配器（adapter）**，将它转换为一个合法的 Handler，如下所示：

```go
mux := http.NewServeMux()
mux.Handle("/", http.HandlerFunc(home))
```

​	`http.HandlerFunc()` 适配器的工作原理是：它会自动为 `home` 函数“附加”一个 `ServeHTTP()` 方法。当该方法被调用时，它实际上只是转而执行原始 `home` 函数的内容。

​	从本质上讲，这是一种略显绕但非常方便的方式，用来将一个普通函数“转换”为满足 `http.Handler` 接口的对象。

​	在本项目的前面部分，我们一直使用 `HandleFunc()` 方法将 handler 函数注册到 `servemux` 中。实际上，这只是 Go 提供的一种语法糖（syntactic sugar），它在内部完成了“函数转 Handler 并注册”的两个步骤，而无需开发者手动操作。

​	上面的代码在功能上等价于：

```go
mux := http,NewServeMux()
mux.HandleFunc("/", home)
```

##### 链式组合

​	眼尖的读者可能已经在本项目一开始就注意到了一个有趣的细节：`http.ListenAndServe()` 函数的第二个参数，实际上要求传入一个 **`http.Handler` 对象**

```go
func ListenAndServe(addr string, handler Handler) error
```

​	但是我们一直传入的参数其实是一个 **servemux（请求多路复用器）**。

​	之所以可以这样做，是因为 `servemux` 同样实现了 `ServeHTTP()` 方法，因此它也满足 `http.Handler` 接口的要求。

​	从理解上来说，可以把 `servemux` 看作一种特殊的 **Handler（处理器）**：它本身并不直接返回响应，而是根据请求的 URL 路径，将请求转发给另一个具体的 handler 进行处理。

​	这种理解方式有助于简化整体模型，而且也并不抽象——在 Go 中，将多个 handler 进行**链式组合（chaining handlers）**是一种非常常见的编程模式，在后续项目中我们也会大量使用。

​	实际上，真实的执行过程如下：

​	当服务器接收到一个新的 HTTP 请求时，会首先调用 `servemux` 的 `ServeHTTP()` 方法。该方法会根据请求的 URL 路径查找对应的 handler，然后再调用该 handler 的 `ServeHTTP()` 方法进行处理。

​	可以把一个 Go Web 应用理解为：**一条由多个 `ServeHTTP()` 方法依次调用所构成的处理链条**。

##### 请求是并发处理的

​	还有一件非常重要的事情需要强调：所有进入的 HTTP 请求，都会在各自独立的 **goroutine** 中进行处理。对于高并发的服务器来说，这意味着你的 handler 中的代码（以及它所调用的代码）很可能是**并发执行的**。虽然这种机制让 Go 能够具备非常高的性能表现，但它的代价是：当多个 handler 同时访问共享资源时，你必须注意并发安全问题，并防止出现**竞态条件（race conditions）**。



### 3. 配置和错误处理

---

​	在本书的这一部分，我们将进行一些“整理工作（housekeeping）”。我们不会为应用程序增加太多新的功能，而是专注于改进其结构，使其在不断扩展时更易于维护和管理。

你将学习如何：

-  使用命令行参数（command-line flags），以一种简单且符合 Go 语言习惯的方式，在程序运行时配置应用的各项设置。
- 增强应用的日志（log）输出，使其包含更多信息，并能够根据日志类型（或级别）进行分类管理。
- 以一种可扩展、类型安全且不影响测试编写的方式，将依赖项传递给各个 handler。
-  集中化错误处理逻辑，从而避免在代码中重复编写相同的错误处理流程。



#### 3.1. 管理配置设置

---

​	目前，我们的 Web 应用程序在 `main.go` 文件中包含了几个**硬编码（hard-coded）**的配置项：

- 服务器监听的网络地址（当前为 `":4000"`）。
- 静态资源目录的路径（当前为 `"./ui/static"`）。

​	将这些配置直接写死在代码中并不是一种理想的做法。这样会导致**配置与业务代码耦合在一起**，并且无法在程序运行时动态修改配置。而在实际开发中，不同的运行环境（如开发环境、测试环境和生产环境）通常需要使用不同的配置，因此运行时可配置性非常重要。

​	在本章中，我们将首先对这一问题进行改进，使服务器的监听地址能够在**程序运行时进行配置**。

##### 命令行标志

​	在 Go 中，一种常见且符合 Go 语言编程习惯（idiomatic）的配置管理方式，是在**启动应用程序时使用命令行参数（Command-line Flags）**来指定配置项。例如：

```bash
$ go run ./cmd/web -addr=":80"
```

​	在 Go 中，接收并解析**命令行参数（Command-line Flag）**最简单的方式，就是使用下面这样的一行代码：

```go
addr := flag.String("addr", ":4000", "HTTP network address")
```

​	这行代码实际上定义了一个新的**命令行参数（Command-line Flag）**，其名称为 `addr`，默认值为 `":4000"`，并附带了一段简短的帮助信息，用于说明该参数的作用。

​	程序运行时，该命令行参数的值将保存到 `addr` 变量中。

​	接下来，我们将在应用程序中使用这个命令行参数，并将原来硬编码的服务器监听地址替换为可通过命令行指定的配置。

```
File: cmd/web/main.go	
```

```go
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// 定义一个名为 "addr" 的新命令行标志，默认值为 ":4000"
	// 并附带简短帮助信息说明该标志的用途
	// 该标志的值将在运行时存入 addr 变量
	addr := flag.String("addr", ":4000", "HTTP network address")

	// 关键：调用 flag.Parse() 解析命令行标志
	// 这会读取命令行标志的值并赋给 addr 变量
	// 必须在使用 addr 变量之前调用，否则它将始终包含默认值 ":4000"
	// 若解析过程中遇到任何错误，应用将被终止
	flag.Parse()
	
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// 创建一个静态文件服务器，用于提供 "./ui/static" 目录中的文件
	// 注意，传递给 http.Dir() 的路径是相对于项目根目录（Project Root）的相对路径
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// 使用 mux.Handle() 注册路由，将 fileServer 设置为所有以 "/static/"
	// 开头的 URL 路径的处理器（Handler）。对于匹配到的请求，在交由
	// fileServer 处理之前，会先移除 URL 路径中的 "/static" 前缀
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
```

​	保存上述修改后，在启动应用程序时尝试使用 **`-addr`** 命令行参数。

​	此时，你会发现服务器将监听你所指定的网络地址，例如：

```bash
$ go run ./cmd/web -addr=":9999"
2026/06/28 15:19:24 Starting server on :9999
```

> [!NOTE]
>
> **注意：**端口号 **`0~1023`** 属于**系统保留端口（Privileged Ports）**，通常只有具有 **Root（Linux/Unix）** 或 **Administrator（Windows）** 权限的进程才能使用。 如果应用程序尝试监听这些端口，而当前进程没有足够的权限，那么在启动时通常会出现 **`bind: permission denied`（绑定端口失败：权限不足）** 的错误提示。

##### 默认值

**命令行参数（Command-line Flags）是完全可选的。**

​	例如，如果在启动应用程序时没有指定 **`-addr`** 参数，服务器将自动使用 **`:4000`** 作为监听地址，也就是我们为该参数设置的**默认值（Default Value）**。

```bash
$ go run ./cmd/web
2026/06/28 15:22:30 Starting server on :4000
```

​	命令行参数的**默认值（Default Value）**并没有固定的规定，应根据实际需求进行设置。

​	作者通常会将默认值设置为适合**开发环境（Development Environment）**的配置，这样在开发应用程序时可以减少输入，提高开发效率。

​	不过，具体采用哪种方式因人而异（YMMV，Your Mileage May Vary）。你也可以选择一种更加稳妥的做法，将默认值设置为适用于**生产环境（Production Environment）**的配置。

##### 类型转换

​	在上面的代码中，我们使用了`flag.String()`函数来定义命令行标志。这样做的好处是将用户在运行时提供的任何值转换为字符串类型。如果该值无法转换为字符串，应用程序将记录错误并退出。

​	Go 标准库还提供了许多用于定义命令行参数的函数，例如 **`flag.Int()`**、**`flag.Bool()`** 和 **`flag.Float64()`** 等。

​	它们的使用方式与 **`flag.String()`** 完全相同，不同之处在于，这些函数会根据各自的类型，自动将命令行参数的值转换为对应的数据类型。

##### 自动帮助

​	另一个很棒的功能是您可以使用 `-help` 标志列出应用程序的所有可用命令行标志及其随附的帮助文本。试一试：

```bash
$ go run ./cmd/web -help
Usage of C:\...\Local\go-build\ea\eafe0774...\web.exe:
  -addr string
        HTTP network address (default ":4000")
```

​	总体来说，到目前为止，我们已经完成了一个不错的改进。

​	我们采用了一种**符合 Go 语言编程习惯（idiomatic）**的方式，在程序运行时管理应用程序的配置；同时，也为应用程序与其运行配置之间建立了一套**明确且具有文档说明的配置接口**。

##### 附加信息

**环境变量**

​	如果你之前开发和部署过 Web 应用程序，可能会想到：**环境变量（Environment Variables）**怎么办？将配置项存放在环境变量中，不也是一种良好的实践吗？

​	当然可以。你可以将应用程序的配置项保存在环境变量中，并通过 **`os.Getenv()`** 函数在程序中直接读取这些配置，例如：

```go
addr := os.Getenv("SNIPPETBOX_ADDR")
```

​	不过，与**命令行参数（Command-line Flags）**相比，使用环境变量也存在一些不足之处。

​	无法为配置项指定默认值。如果环境变量不存在，`os.Getenv()` 返回的将是一个空字符串（`""`）。

​	 无法像命令行参数一样，自动提供 `-help` 帮助信息。
​	`os.Getenv()` 的返回值始终是 **`string`** 类型，不会像 `flag.Int()`、`flag.Bool()` 等命令行参数函数那样，自动完成数据类型转换。

​	因此，一种更好的做法是**结合环境变量和命令行参数**：在启动应用程序时，将环境变量的值作为命令行参数传递给程序。这样既能够利用环境变量管理配置，又能够享受命令行参数带来的默认值、帮助信息以及自动类型转换等优势。例如：

```bash
# Windows cmd 
$ set SNIPPETBOX_ADDR=":9999"
# Windows PowerShell
$ $env:SNIPPETBOX_ADDR=":9999"
```

**Boolean Flags**

​	对于使用 **`flag.Bool()`** 定义的**布尔类型命令行参数**，如果省略参数值，则等同于将该参数设置为 **`true`**。

​	也就是说，下面两条命令的效果是完全相同的：

```bash
$ go run example.go -flag=true
$ go run example.go -flag
```

​	如果需要将**布尔类型命令行参数**设置为 **`false`**，则必须显式指定 **`-flag=false`**。

**预先存在的变量**

​	除了使用 `flag.String()`、`flag.Int()` 等函数外，Go 还提供了 **`flag.StringVar()`**、**`flag.IntVar()`**、**`flag.BoolVar()`** 等函数，用于将**命令行参数解析后的值直接存储到已有变量的内存地址中**。

​	当希望将所有配置项统一保存到一个**结构体（struct）**中时，这种方式会非常方便。

​	下面是一个简单的示例：

```go
type Config struct {
    Addr 	  String
    StaticDir String
}

...

cfg := new(Config)
flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
flag.StringVar(&cfg.StaticDir, ""static-dir", "./ui/static", "Path to static assets")
flag.Parse()
```



#### 3.2. 分级记录

---

​	目前，在我们的 `main.go` 文件中，日志信息是通过 **`log.Printf()`** 和 **`log.Fatal()`** 两个函数输出的。

​	这两个函数都会使用 Go 标准库提供的**默认日志记录器（Standard Logger）**输出日志。默认情况下，它会在每条日志前添加**本地日期和时间**，并将日志写入**标准错误输出（stderr）**，通常会显示在终端窗口中。

​	其中，`log.Fatal()` 在输出日志之后，还会调用 **`os.Exit(1)`**，使程序立即退出。

​	通常，我们可以将日志划分为两种不同的**日志级别（Log Level）**：

​	**信息日志（Informational Log）**：用于记录程序正常运行的信息，例如 `"Starting server on :4000"`。
​	**错误日志（Error Log）**：用于记录程序运行过程中发生的错误信息。

```go
log.Printf("Starting server on %s", *addr)	// Information message
err :=http.ListenAndServe(*addr, mux)
log.Fatal(err) // Error message
```

​	接下来，我们将为应用程序增加**分级日志（Leveled Logging）**功能，使信息日志和错误日志能够采用不同的管理方式。具体来说：

​	**信息日志（INFO）**：以 `"INFO"` 作为日志前缀，并输出到**标准输出（stdout）**。

​	**错误日志（ERROR）**：以 `"ERROR"` 作为日志前缀，并输出到**标准错误输出（stderr）**，同时记录调用日志的位置（包括文件名和行号），以便于程序调试。

​	实现这一功能有多种方式，其中一种简单且清晰的做法是使用 **`log.New()`** 函数创建两个自定义的日志记录器（Logger）。

​	打开 `main.go` 文件，并按照下面的内容进行修改：

```
File: cmd/web/main.go
```

```go
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

	// 使用 log.New() 创建一个用于记录信息日志的 logger
	// 它接受三个参数：日志写入目标 (os.Stdout)、消息前缀字符串 ("INFO" 后跟一个制表符)
	// 以及用于指定包含哪些附加信息的标志（本地日期和时间）
	// 注意标志通过按位或运算符 | 组合
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// 以同样方式创建用于记录错误日志的 logger，但使用 stderr 作为目标
	// 并使用 log.Lshortfile 标志包含相关文件名和行号
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// 创建一个静态文件服务器，用于提供 "./ui/static" 目录中的文件
	// 注意，传递给 http.Dir() 的路径是相对于项目根目录（Project Root）的相对路径
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 使用两个新的 logger 来写入消息，替代标准 logger。
	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
```

​	下面来实际体验一下这两个日志记录器的效果。

​	首先运行应用程序，然后再打开一个新的终端窗口，再次启动同一个应用程序。

​	由于服务器要监听的网络地址 **`:4000`** 已经被第一个应用程序占用，因此第二次启动时会产生一个错误。

​	此时，在第二个终端窗口中，你应该会看到类似下面的日志输出：

```bash
$ go run ./cmd/web
INFO    2026/06/28 15:49:46 Starting server on :4000
ERROR   2026/06/28 15:49:46 main.go:38: listen tcp :4000: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
exit status 1
```

​	请注意，这两条日志消息使用了**不同的前缀**，因此在终端中可以很容易地区分**信息日志（INFO）**和**错误日志（ERROR）**。

​	另外，错误日志还额外包含了**调用日志记录器的源文件名和代码行号**（例如 `main.go:38`），这对于定位问题和调试程序非常有帮助。

> [!TIP]
>
> **如果希望日志中显示**完整的文件路径，而不仅仅是文件名，可以在创建自定义 Logger 时使用 **`log.Llongfile`** 标志，而不是 **`log.Lshortfile`**。 此外，如果希望日志统一使用 **UTC 时间**（而不是本地时间），可以在 Logger 的配置中添加 **`log.LUTC`** 标志。

##### 解耦日志记录

​	将日志输出到标准输出流（**stdout**）和标准错误流（**stderr**）的一个重要优势是：可以实现应用程序与日志系统的**解耦（decoupling）**。

​	也就是说，应用程序本身不需要关心日志的路由方式或存储位置，从而可以根据不同环境更灵活地管理日志。

​	在开发环境中，由于标准输出会直接显示在终端，因此可以很方便地查看日志内容。

​	而在测试环境（staging）或生产环境（production）中，可以将标准输出和标准错误重定向到最终存储位置以便查看和归档。这些存储位置可以是本地磁盘文件，也可以是类似 **Splunk** 这样的日志服务平台。无论采用哪种方式，日志的最终去向都可以由运行环境独立控制，而不需要修改应用程序本身。

​	例如，我们可以在启动应用程序时，将 stdout 和 stderr 重定向到磁盘文件中，如下所示：

```bash
$ go run ./cmd/web 1>>tmp/info.log 2>>tmp/error.log
```

> [!NOTE]
>
> 使用双重重定向符号 **`>>`** 时，输出内容会被**追加到已有文件末尾**，而不是在启动应用程序时**覆盖（截断）原文件内容**。

##### http.Server 错误日志

​	我们还需要对应用程序做一个额外的修改。

​	默认情况下，如果 Go 的 HTTP 服务器在运行过程中遇到错误，它会使用标准日志记录器（standard logger）进行输出。

​	为了保持日志风格的一致性，更合理的做法是改用我们自定义的 **`errorLog` 日志记录器**来输出这些错误信息。

​	要实现这一点，我们需要显式创建一个 **`http.Server` 结构体**，用于集中管理服务器的配置参数，而不是继续使用 `http.ListenAndServe()` 这种快捷方式。

​	下面通过代码示例来说明这一做法：

````
File: cmd/web/main.go
```

```go
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 初始化一个新的 http.Server 结构体。设置 Addr 和 Handler 字段，
	// 使服务器使用与之前相同的网络地址和路由，并设置 ErrorLog 字段，
	// 以便服务器在出现任何问题时使用自定义的 errorLog logger。
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	
	infoLog.Printf("Starting server on %s", *addr)
	// 调用新 http.Server 结构体的 ListenAndServe() 方法。
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
```

##### 附加信息

**其他日志记录方法**

​	到目前为止，在本书中我们已经使用了 `Println()`、`Printf()` 和 `Fatal()` 等方法来输出日志信息，但 Go 还提供了许多其他日志方法，值得进一步熟悉和了解。

​	作为一个经验法则（rule of thumb），应尽量避免在 `main()` 函数之外使用 `Panic()` 和 `Fatal()` 这类方法。

​	更推荐的做法是：在其他函数中通过**返回错误（error）**的方式进行处理，而只有在 `main()` 函数中才根据情况直接触发 panic 或退出程序。

**并发日志记录**

​	通过 `log.New()` 创建的自定义日志记录器是**并发安全（concurrency-safe）**的。

​	你可以在多个 goroutine 之间共享同一个 logger，并在 handler 中同时使用，而无需担心发生竞态条件（race conditions）。

​	但需要注意的是，如果存在多个 logger 同时向同一个输出目标写入日志，那么必须确保该输出目标底层的 `Write()` 方法本身也是**并发安全的**。

**记录到文件**

​	如前所述，通常更推荐的做法是将日志输出到标准输出流（standard streams），并在运行时将其重定向到文件中。

​	但如果你不希望采用这种方式，也可以在 Go 中直接打开一个文件，并将其作为日志输出的目标。

​	下面是一个简单示例：

```go
f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
if err != nil {
    log.Fatal(err)
}
defer f.Close()

infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
```



#### 3.3. 依赖注入

---

​	我们的日志系统还存在一个需要修正的问题。

​	如果你打开 `handlers.go` 文件，会发现 `home` 这个 handler 函数仍然在使用 Go 的**标准日志记录器（standard logger）**来输出错误信息，而不是我们希望统一使用的 **`errorLog` 日志记录器**。

```go
func home(w http.ResponseWriter, r *http.Request) {
	...

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
```

​	这就引出了一个非常重要的问题：我们应该如何让在 `main()` 中创建的 **`errorLog` 日志记录器**，在 `home` handler 函数中也可以被访问到？

​	进一步来看，这个问题可以扩展为一个更通用的设计问题：在 Web 应用中，handler 往往需要依赖多种资源，例如数据库连接池、统一的错误处理器、模板缓存等。那么，本质上我们需要解决的是：**如何将这些依赖提供给各个 handler 使用？**

​	针对这个问题，有多种实现方式，其中最简单的一种是使用**全局变量（global variables）**。但通常来说，这并不是推荐做法。

​	更好的实践是采用**依赖注入（dependency injection）**的方式将依赖传递给 handler。这种方式可以让代码更加清晰、明确，同时减少出错概率，也更方便进行单元测试。

​	对于像我们这样所有 handler 都位于同一个 package 的项目来说，一个比较优雅的实现方式是：将所有依赖封装到一个自定义的 **application 结构体（struct）**中，然后将 handler 方法定义为该结构体的方法。

​	下面我们来演示具体实现。

​	打开 `main.go` 文件，并创建一个新的 `application` 结构体，如下所示：

```
File: cmd/web/handlers.go
```

```go
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
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 然后使用模板集的 Execute() 方法将模板内容写入响应体
	// Execute() 的最后一个参数用于传入动态数据，目前暂设为 nil
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
```

​	最后，我们将在 `main.go` 文件中把所有组件**串联（wire）起来**，完成整体整合：

`````
File: cmd/web/main.go
```

```go
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 初始化一个新的 http.Server 结构体。设置 Addr 和 Handler 字段
	// 使服务器使用与之前相同的网络地址和路由，并设置 ErrorLog 字段
	// 以便服务器在出现任何问题时使用自定义的 errorLog logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// 调用新 http.Server 结构体的 ListenAndServe() 方法
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
```

​	我理解这种实现方式一开始可能会让人觉得有些复杂甚至略显绕（convoluted），尤其是相比于直接将 `infoLog` 和 `errorLog` 定义为全局变量这种更简单的做法。

​	但请继续坚持这种方式。

​	随着应用程序规模逐渐增长，当我们的 handler 需要依赖的组件越来越多时，这种设计模式的优势将会逐渐体现出来。

##### 添加一个故意的错误

​	接下来，我们通过在应用中**故意制造一个错误（deliberate error）**来进行验证。

​	打开终端，将 `ui/html/home.page.tmpl` 文件重命名为 `ui/html/home.page.bak`。

​	当我们重新运行应用程序并请求首页时，由于 `ui/html/home.page.tmpl` 文件已经不存在，因此此时应该会触发一个错误。

​	现在可以执行上述修改操作：

```bash
$ ren ui/html/home.page.tmpl home.page.bak  
```

​	然后重新运行应用程序，并在浏览器中访问 ``。

​	此时，你应该会在浏览器中看到一个 **Internal Server Error（内部服务器错误）** 的 HTTP 响应，同时在终端中也会输出一条对应的错误日志，其格式类似如下：

```bash
$ go run ./cmd/web
INFO    2026/06/28 16:38:13 Starting server on :4000
ERROR   2026/06/28 16:38:48 handlers.go:35: open ./ui/html/home.page.tmpl: The system cannot find the file specified.
```

​	请注意，此时日志信息已经以 **`ERROR`** 作为前缀，并且错误来源显示为 `handlers.go` 文件的第 35 行。

​	这很好地说明了：我们自定义的 **`errorLog` 日志记录器**已经作为依赖成功传递到了 `home` handler 中，并且运行符合预期。

​	目前请先保留这个“人为制造的错误”，在下一章中我们还会用到它。

##### 附加信息

**依赖注入闭包**

​	如果你的 handler 分布在多个不同的 package 中，那么我们当前使用的这种**依赖注入模式**就不再适用。

​	在这种情况下，可以采用另一种替代方案：创建一个 `config` 包，并在其中导出一个 `Application` 结构体，然后让 handler 函数通过**闭包（closure）**的方式“捕获”该结构体，从而访问其中的依赖。

​	其大致实现方式如下：

```go
func mian() {
    app := &config.Application {
        ErrorLog: log.New((os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    }
    
	mux.Handle("/", handlers.Home(app)) 
}
                          
func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		...
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		...
	}
}
```

​	你可以在这个 Gist 中找到一个更完整、更具体的示例，展示如何使用**闭包（closure）模式**来实现依赖注入。



#### 3.4. 集中错误处理

---

​	接下来，我们将通过把一部分**错误处理代码移动到辅助方法（helper methods）中**，来整理（neaten up）当前的应用结构。

这样做可以更好地实现职责分离（separation of concerns），并避免在后续开发过程中重复编写相同的代码。

​	现在请在 `cmd/web` 目录下创建一个新的文件 `helpers.go`：

```bash
$ type nul > cmd/web/helpers.go 
```

​	然后在该文件中添加如下代码：

```
File: cmd/web/helpers.go 
```

```go
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
	app.errorLog.Println(trace)

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
```

​	这里新增的代码量不多，但引入了一些非常值得理解的新特性。

- 在 `serverError()` 辅助函数中，我们使用了 `debug.Stack()` 来获取当前 goroutine 的**堆栈信息（stack trace）**，并将其追加到日志中。通过堆栈信息，我们可以看到程序的执行路径，这在排查和调试错误时非常有帮助。

- 在 `clientError()` 辅助函数中，我们使用了 `http.StatusText()` 方法，用于自动生成 HTTP 状态码对应的**可读文本描述**。例如：`http.StatusText(400)` 会返回 `"Bad Request"`。

- 另外，我们开始使用 `net/http` 包中提供的**命名常量（named constants）**来表示 HTTP 状态码，而不是直接写数字。例如：在 `serverError()` 中使用 `http.StatusInternalServerError` 替代 `500`在 `notFound()` 中使用 `http.StatusNotFound` 替代 `404`。

  使用状态码常量是一种非常好的实践，它可以让代码更加清晰，并具备更强的自解释性，尤其是在处理一些不常见的状态码时更为明显。你可以在官方文档中查看完整的状态码常量列表。

​	完成以上修改后，回到 `handlers.go` 文件，并将代码更新为使用新的辅助函数：

```
File: cmd/web/handlers.go
```

```go
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
```

​	完成更新后，重新启动应用程序，并在浏览器中访问 `http://localhost:4000`。

​	此时，由于我们之前保留的“人为错误”仍然存在，程序会再次触发该错误。同时，你将在终端中看到对应的错误信息以及完整的堆栈追踪（stack trace）输出：

```bash
$ go run ./cmd/web
INFO    2026/06/28 17:04:15 Starting server on :4000
ERROR   2026/06/28 17:04:19 helpers.go:13: open ./ui/html/home.page.tmpl: The system cannot find the file specified.
goroutine 21 [running]:
runtime/debug.Stack()
        D:/Golang/GO/Gocode/go1.26.1/src/runtime/debug/stack.go:26 +0x5e
main.(*application).serverError(0x29d6fb96210, {0x7ff6b7029200, 0x29d6fcb8000}, {0x7ff6b7026680?, 0x29d6fc82120?})
        D:/code/snippetbox/cmd/web/helpers.go:12 +0x5d
main.(*application).home(0x29d6fb96210, {0x7ff6b7029200, 0x29d6fcb8000}, 0x29d6fc86020?)
        D:/code/snippetbox/cmd/web/handlers.go:35 +0x105
net/http.HandlerFunc.ServeHTTP(0x29d6fbac300?, {0x7ff6b7029200?, 0x29d6fcb8000?}, 0x7ff6b6eaab76?)
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:2286 +0x29
net/http.(*ServeMux).ServeHTTP(0x7ff6b6c00ab9?, {0x7ff6b7029200, 0x29d6fcb8000}, 0x29d6fc92000)
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:2828 +0x1c7
net/http.serverHandler.ServeHTTP({0x29d6fc88000?}, {0x7ff6b7029200?, 0x29d6fcb8000?}, 0x6?)
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:3311 +0x8e
net/http.(*conn).serve(0x29d6fc0c1b0, {0x7ff6b7029ba8, 0x29d6fb8d6b0})
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:2073 +0x650
created by net/http.(*Server).Serve in goroutine 1
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:3464 +0x485
```

​	仔细观察你会发现一个小问题：当前日志中显示的文件名和行号是 `helpers.go:13`，因为日志实际上是在这个位置被写入的。

但我们真正希望记录的是**调用发生位置（上一层调用栈）**的文件名和行号，这样才能更清楚地定位错误的真实来源。

为了解决这个问题，我们可以修改 `serverError()` 方法，改用 logger 的 `Output()` 方法，并将堆栈深度（frame depth）设置为 2。

重新打开 `helpers.go` 文件，并按如下方式进行修改：

```
File: cmd/web/helpers.go
```

```go
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
```

​	如果现在再尝试运行一次，你应该会发现日志中已经正确显示了**对应的文件名和行号（`handlers.go:35`）**，从而准确指向错误真正发生的位置：

```bash
$ go run ./cmd/web
INFO    2026/06/28 17:07:36 Starting server on :4000
ERROR   2026/06/28 17:07:40 handlers.go:35: open ./ui/html/home.page.tmpl: The system cannot find the file specified.
goroutine 8 [running]:
runtime/debug.Stack()
        D:/Golang/GO/Gocode/go1.26.1/src/runtime/debug/stack.go:26 +0x5e
main.(*application).serverError(0x236529f60250, {0x7ff7044b9200, 0x23652a138000}, {0x7ff7044b6680?, 0x23652a102120?})
        D:/code/snippetbox/cmd/web/helpers.go:12 +0x5a
main.(*application).home(0x236529f60250, {0x7ff7044b9200, 0x23652a138000}, 0x23652a106020?)
        D:/code/snippetbox/cmd/web/handlers.go:35 +0x105
net/http.HandlerFunc.ServeHTTP(0x236529fe0300?, {0x7ff7044b9200?, 0x23652a138000?}, 0x7ff70433ab76?)
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:2286 +0x29
net/http.(*ServeMux).ServeHTTP(0x7ff704090ab9?, {0x7ff7044b9200, 0x23652a138000}, 0x23652a112000)
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:2828 +0x1c7
net/http.serverHandler.ServeHTTP({0x23652a108000?}, {0x7ff7044b9200?, 0x23652a138000?}, 0x6?)
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:3311 +0x8e
net/http.(*conn).serve(0x23652a0a21b0, {0x7ff7044b9ba8, 0x236529f5d6e0})
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:2073 +0x650
created by net/http.(*Server).Serve in goroutine 1
        D:/Golang/GO/Gocode/go1.26.1/src/net/http/server.go:3464 +0x485
```

##### 恢复故意的错误

​	到目前为止，这个“人为制造的错误”已经不再需要了，因此我们可以将其修复，如下所示：

```bash
$ ren ui/html/home.page.bak home.page.tmpl
```



#### 3.5. 隔离应用程序路由

---

​	在进行代码重构的过程中，还有一个值得做的优化。

​	目前我们的 `main()` 函数已经开始变得有些臃肿（crowded），为了让结构更加清晰、职责更加单一，我们可以将应用程序的**路由定义（route declarations）**拆分出去，放到一个独立的 `routes.go` 文件中，例如：

```bash
$ type nul > cmd/web/routes.go 
```

```
File: cmd/web/routes.go
```

```go
package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	return mux
}
```

​	随后，我们可以更新 `main.go` 文件，使其改为使用这个新的实现：

```
File: cmd/web/main.go
```

```go
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
```

这样处理之后，代码结构明显更加清晰（neater）了。

当前应用的路由已经被独立封装在 `app.routes()` 方法中，实现了良好的隔离与封装。

同时，`main()` 函数的职责也被有效收敛，现在它只负责以下三件事情：

- 解析应用程序的运行时配置（runtime configuration）
- 初始化并构建 handler 所需的依赖（dependencies）
- 启动 HTTP 服务器（HTTP server）



### 4. 数据库驱动的响应

---

为了让我们的 Snippetbox Web 应用真正具备实用价值，我们需要一个用于**存储（persist）用户输入数据**的地方，并且能够在运行时对这些数据进行动态查询。

可用于本应用的数据存储方案有很多种，每种方案都有各自的优缺点，而我们将选择目前非常流行的**关系型数据库 MySQL**。

在这一部分中，你将学习以下内容：

- 如何从 Go Web 应用连接 MySQL（特别是如何建立一个**可复用的数据库连接池**）
- 如何创建一个独立的 `models` 包，使数据库相关逻辑可复用，并与 Web 应用解耦
- 如何使用 Go 的 `database/sql` 包中的不同函数执行各类 SQL 语句，以及如何避免常见错误（例如导致服务器资源耗尽的问题）
- 如何通过正确使用**占位符参数（placeholder parameters）**来防止 SQL 注入攻击
- 如何使用**事务（transactions）**，将多个 SQL 语句作为一个原子操作执行

#### 4.1. 设置 MySQL

如果你正在跟着本书进行实践，那么此时需要在你的计算机上安装 MySQL。

MySQL 官方文档提供了适用于各种操作系统的完整安装指南；如果你使用的是 Windows，可以通过以下方式进行安装：

Windows 下 MySQL **安装简易教程**

1. 下载 MySQL 安装包

访问 MySQL 官方下载页面：

https://dev.mysql.com/downloads/installer/

选择 **MySQL Installer for Windows**（推荐下载 `mysql-installer-web-community` 或完整版）。

---

2. 运行安装程序

双击下载的 `.msi` 文件启动安装向导。

常见选择：

- 选择 **Developer Default**（开发者默认配置，推荐）
  - 包含 MySQL Server + Workbench + 工具

或：

- 选择 **Custom**（自定义安装）

---

3. 安装 MySQL Server

在组件列表中确认：

- MySQL Server
- MySQL Workbench（可选但推荐）

点击 **Next → Execute** 开始安装。

---

4. 配置 MySQL Server

安装完成后进入配置界面：

4.1 类型选择

- Development Computer（开发环境，推荐）

4.2 端口

- 默认端口：`3306`（保持默认即可）

4.3 认证方式

- 推荐：Use Strong Password Encryption

4.4 设置 root 密码

- 设置并记住 root 密码（非常重要）

---

5. 创建 Windows 服务

- 勾选：**Run MySQL as a Windows Service**
- 服务名默认即可：MySQL80

---

6. 完成安装

点击 **Finish** 完成安装。

---

7. 验证安装是否成功

方法一：命令行

打开 CMD 或 PowerShell：

```bash
mysql -u root -p
```

输入密码后进入 MySQL 即表示成功。

---

方法二：MySQL Workbench

打开 Workbench：

- 新建连接
- Host：127.0.0.1
- Port：3306
- User：root

点击连接测试。

---

8. 常见问题

❗ mysql 不是内部命令

解决方法：  
把 MySQL 的 `bin` 目录加入系统环境变量 PATH，例如：

```
C:\Program Files\MySQL\MySQL Server 8.0\bin
```

---

##### 搭建数据库

当 MySQL 安装完成后，你应该可以在终端中以 **root 用户**身份连接到数据库。具体的连接命令会根据你安装的 MySQL 版本而有所不同。

例如，在 MySQL 5.7 版本中，你可以通过输入以下命令进行连接：

```bash
$ sudo mysql
mysql>
```

但如果上述命令无法正常工作，可以尝试下面的命令，并输入你在安装过程中设置的密码：

```bash
$ $ mysql -u root -p
Enter password:
mysql>
```

连接成功之后，我们首先需要在 MySQL 中创建一个数据库，用于存储本项目的所有数据。

请将以下命令复制并粘贴到 MySQL 命令行（mysql prompt）中，以 UTF8 编码创建一个新的 `snippetbox` 数据库。

```sql
-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Switch to using the `snippetbox` database.
USE snippetbox;
```

接下来，请复制并执行下面的 SQL 语句，在 `snippetbox` 数据库中创建一个 **`snippets`** 表，用于存储应用程序中的代码片段（text snippets）数据：

```sql
-- Create a `snippets` table.
CREATE TABLE snippets (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	created DATETIME NOT NULL,
	expires DATETIME NOT NULL
);
-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);
```

`snippets` 表中的每条记录都包含一个整型的 **`id`** 字段，作为该代码片段的**唯一标识符（Unique Identifier）**。

此外，每条记录还包含以下字段：

\- **`title`**：用于存储代码片段的标题
\- **`content`**：用于存储代码片段的具体内容
\- **`created`**：记录代码片段的创建时间
\- **`expires`**：记录代码片段的过期时间

接下来，我们再向 `snippets` 表中插入几条**示例数据（Placeholder Data）**，这些数据将在后续几章中使用。

书中使用的是几首简短的俳句（Haiku）作为代码片段的内容，不过具体内容并不重要，只要能够作为测试数据即可。

```sql
-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
	'An old silent pond',
	'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
	'Over the wintry forest',
	'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
	'First autumn morning',
	'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```

##### 创建新用户

从**安全性（Security）**的角度来看，在 Web 应用程序中直接使用 **`root`** 用户连接 MySQL 并不是一种好的做法。

更推荐的做法是创建一个**权限受限的数据库用户（Database User）**，只授予其应用程序实际需要的数据库权限。

因此，在保持 MySQL 命令行（MySQL Prompt）连接的情况下，请执行下面的 SQL 语句，创建一个新的 Web 用户，并仅授予其在当前数据库上的 **`SELECT`** 和 **`INSERT`** 权限。

```sql
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```

完成上述操作后，输入 `exit` 命令退出 MySQL 命令行（MySQL Prompt）。

##### 测试新用户

您现在应该能够使用以下命令以 web用户身份连接到 snippetbox数据库。出现提示时，输入您刚刚设置的密码。

```sql
mysql -D snippetbox -u web -p
Enter password: 

mysql>
```

如果权限配置正确，你应该能够正常执行 **`SELECT`** 和 **`INSERT`** 操作；而对于 **`DROP TABLE`**、**`GRANT`** 等超出授权范围的命令，则会执行失败。

```sql
mysql> SELECT id, title,expires FROM snippets;
+----+------------------------+---------------------+
| id | title                  | expires             |
+----+------------------------+---------------------+
|  1 | An old silent pond     | 2027-06-28 10:44:44 |
|  2 | Over the wintry forest | 2027-06-28 10:48:34 |
|  3 | First autumn morning   | 2026-07-05 10:50:55 |
+----+------------------------+---------------------+
3 rows in set (0.00 sec)

mysql> DROP TABLE snippets;
ERROR 1142 (42000): DROP command denied to user 'web'@'localhost' for table 'snippets'
```



#### 4.2. 安装数据库驱动程序

---

为了在 Go Web 应用程序中使用 MySQL，我们需要安装一个**数据库驱动（Database Driver）**。

数据库驱动充当 Go 程序与 MySQL 数据库之间的**桥梁**，负责将 Go 发出的数据库操作转换为 MySQL 能够识别和执行的命令，同时将数据库返回的结果转换为 Go 程序可以处理的数据。

Go 官方 Wiki 提供了多种数据库驱动可供选择，而本项目将使用目前应用最广泛的 **`go-sql-driver/mysql`** 驱动。

下载该驱动，请进入项目根目录，并执行下面的 `go get` 命令：

```bash
$ go get github.com/go-sql-driver/mysql@v1
go: downloading github.com/go-sql-driver/mysql v1.6.0
```

请注意，这里我们在包路径后面添加了 **`@v1`** 后缀，表示希望下载该依赖**主版本号（Major Version）为 1 的最新版本**。

本书编写时使用的版本是 **`v1.6.0`**，但你实际下载的版本可能是 **`v1.6.1`**、**`v1.7.0`** 或其他 `v1.x.x` 版本，这都是正常的。

这是因为 `go-sql-driver/mysql` 遵循**语义化版本（Semantic Versioning，SemVer）**规范进行版本管理，因此所有 **`v1.x.x`** 版本都能够与本书中的代码保持兼容。

另外，如果你希望**始终下载最新版本**，而不限制主版本号，只需省略 `@version` 后缀即可，例如：

```bash
$ go get github.com/go-sql-driver/mysql
go: added filippo.io/edwards25519 v1.2.0
go: added github.com/go-sql-driver/mysql v1.10.0
```

如果希望下载某个**指定版本**的依赖包，可以在包路径后面指定完整的版本号。例如：

```bash
$ go get github.com/go-sql-driver/mysql@v1.0.3
```

安装完成数据库驱动后，打开项目中的 **`go.mod`** 文件（这是我们在本书开始时创建的 Go 模块文件）。

你应该会看到新增了一条 **`require`** 声明，其中包含：

\- 下载的依赖包路径（Package Path）
\- 实际下载的精确版本号（Exact Version）

例如，它的内容大致如下：

```
File: go.mod
```

```go
module github.com/<your-github-name>/snippetbox

go 1.26.1

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/go-sql-driver/mysql v1.10.0 // indirect
)
```

此外，你还会发现，在项目根目录下新增了一个名为 **`go.sum`** 的文件。

`go.sum` 文件由 Go 工具自动生成，用于记录项目依赖包及其版本对应的**校验值（Checksum）**。

当项目下载或更新依赖时，Go 会根据 `go.sum` 中记录的校验值验证依赖包的完整性，确保下载的内容没有被篡改，从而提高依赖管理的安全性和可靠性。

简单来说：

\- **`go.mod`**：记录项目依赖了哪些模块，以及所需的版本。
\- **`go.sum`**：记录这些模块对应的校验值，用于保证依赖包的完整性和一致性。

![image-20260628203902347](C:\Users\Yang\AppData\Roaming\Typora\typora-user-images\image-20260628203902347.png)

`go.sum` 文件中保存的是项目依赖包的**加密校验和（Cryptographic Checksums）**。

这些校验和用于唯一标识每个依赖包的内容。当 Go 下载依赖时，会根据这些校验值验证文件是否完整、是否被篡改，从而保证项目依赖的安全性和一致性。

打开 `go.sum` 文件后，你应该会看到类似下面这样的内容：

```
File: go.sum
```

```go
github.com/go-sql-driver/mysql v1.10.0 h1:Q+1LV8DkHJvSYAdR83XzuhDaTykuDx0l6fkXxoWCWfw=
github.com/go-sql-driver/mysql v1.10.0/go.mod h1:M+cqaI7+xxXGG9swrdeUIoPG3Y3KCkF0pZej+SK+nWk=
```

与 `go.mod` 文件不同，`go.sum` 并不是设计给开发者手动编辑的，一般情况下你也不需要打开它。

不过，它具有两个非常重要的作用：

- 如果你在终端中运行 `go mod verify` 命令，它会验证你本地已下载依赖包的校验和（checksum）是否与 `go.sum` 文件中的记录一致，从而确保这些依赖包没有被篡改。
- 如果其他开发者需要下载项目的所有依赖（可以通过运行 `go mod download` 完成），那么当下载到的依赖与 `go.sum` 文件中记录的校验和不一致时，Go 会报错提示。

##### 附加信息

**更新包**

当一个依赖包被下载并添加到 `go.mod` 文件后，该依赖包及其版本就会被**固定（fixed）**下来。

不过，在今后的开发过程中，你可能会因为各种原因希望升级到该依赖包的更新版本。

如果要将某个依赖升级到**最新的次版本（Minor）**或**补丁版本（Patch）**，只需在执行 `go get` 命令时添加 `-u` 参数即可，例如：

```bash
$ go get -u github.com/foo/bar
```

或者，如果你希望将依赖升级到**指定版本**，可以使用相同的命令，并在包路径后添加对应的 `@version` 后缀。例如：

```bash
$ go get -u github.com/foo/bar@v2.0.0
```

**移除未使用的包**

有时候，你可能使用 `go get` 下载了某个依赖包，但在后续开发过程中发现实际上已经不再需要它。

遇到这种情况时，你有两种处理方式。

第一种方式是执行 `go get` 命令，并在包路径后添加 `@none` 后缀，以移除该依赖。例如：

```bash
$ go get github.com/foo/bar@none
```

另一种方式是，如果你已经从代码中移除了对该依赖包的所有引用，那么可以执行 `go mod tidy` 命令。

该命令会自动清理项目依赖，将 `go.mod` 和 `go.sum` 文件中所有不再使用的依赖包移除。

```bash
$ go mod tidy -v
```

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



