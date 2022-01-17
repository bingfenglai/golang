# Golang网络编程

## 写在前面

在前面的文章当中，我们学习了go当中的通道（channel）与协程（go-routine）及其应用。接下来，我们将学习go关于网络编程方面的内容。

## Golang网络编程初识

在正式的跟大家讲解网络编程之前，需要提前跟大家说明一下：网络编程是一个很庞大的领域，本文当中只是通过一些简单的demo带大家初步了解golang在网络编程方面的应用。

为了大家能够更好地理解本文，大家需要先了解互联网协议的相关知识点。这包括OSI七层模型，了解各层的作用、协议等。如果大家学习过其他的高级语言，例如Java、Python等，了解socket编程的相关知识点，那这再好不过。

## Golang socket编程

### socket是什么？

socket是一种操作系统提供的进程间通信机制，中文翻译为“套接字”。它用于描述IP地址和端口，是一个通信链的句柄。



> socket最初被翻译为把socket译为“介质(字)”。不久，ARPANET的socket就被翻译为“套接字”，其理由是：
>
> 由于每个主机系统都有各自命名进程的方法，而且常常是不兼容的，因此，要在全网范围内硬把进程名字统一起来是不现实的。所以，每个计算机网络中都要引入一种起介质作用的、全网一致的标准名字空间。这种标准名字，在ARPA网中称作套接字，而在很多其他计算机网中称作信口。更确切地说，进程之间的连接是通过套接字或信口构成的。

在操作系统当中，通常会为应用程序提供一组API，称之为套接字接口。应用程序可以通过套接字接口，向网络发出请求或者应答网络请求，进行数据交换。

一个套接字对（socket pairs）由本地的IP地址和端口、远程的IP地址和端口以及建立连接所使用的协议（protocol）组成。

### 对于socket的理解

socket位于应用层与TCP/IP协议族通信的中间，相当于设计模式当中的门面。对于程序员来说，不需要关系通信协议的具体实现，只需要调用socket提供的相关API，就可以进行网络数据交换。

## Golang的TCP client-server Demo

服务端程序：

```go
package main

import (
	"fmt"
	"go_code/web/tcp/constants"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen(constants.Protocol, constants.Addr)

	if err != nil {
		println("监听端口出错", err)
		return

	} else {
		fmt.Println("服务器启动完成 " + time.Now().Format("2006-01-02 15:04:05"))
	}

	for {

		accept, err := listen.Accept()
		if err != nil {
			println("出现错误:\n", err.Error())
			return

		}

		go doServerStuff(accept)

	}

}

func doServerStuff(conn net.Conn) {

	defer func() {
		println("关闭连接")
		conn.Close()
	}()

	for {
		buf := make([]byte, 256)

		read, err := conn.Read(buf)

		if err != nil {
			if "EOF" == err.Error() {
				println("数据读取结束")
				return
			}
			println("出现错误:", err.Error())
			return

		}

		println("接收到数据：")
		println(string(buf[:read]))
		_, err = conn.Write([]byte(" hello client " + conn.RemoteAddr().String()))

		if err != nil {
			println("出现错误\n", err.Error())
		}

	}

}
```

在服务端程序当中，服务启动后，使用for无限循环接收客户端的请求并启动协程处理来自客户端的请求。



客户端程序：

```go
package main

import (
	"fmt"
	"go_code/web/tcp/constants"
	"net"
	"time"
)

func main() {

	// 建立连接

	conn, err := net.Dial(constants.Protocol, constants.Addr)

	defer conn.Close()

	if err != nil {
		println("建立连接失败", err.Error())
		return

	} else {
		fmt.Println("建立连接成功")
	}
	conn.Write([]byte("hello server " + time.Now().Format("2006-01-02 15:04:05")))

	buf := make([]byte, 1024)
	read, err := conn.Read(buf)

	println("收到服务器发送来的消息")

	if err != nil {
		println("读取数据失败 ", err.Error())
		return
	}

	fmt.Println(string(buf[:read]))

}
```

客户端启动后，与服务器建立连接，并向服务器发送数据。

服务器程序输出：

```go
服务器启动完成 2022-01-16 18:37:53
接收到数据：
hello server 2022-01-16 18:38:01
数据读取结束
关闭连接
```



客户端程序输出：

```go
建立连接成功
收到服务器发送来的消息
 hello client 127.0.0.1:3471

进程完成，并显示退出代码 0
```

以上就是Go 关于tcp编程的一个简单的demo，其中主要应用到了net包和协程的相关知识点。

## http 服务器demo

接下来，我们通过一个小demo来学习Go当中简单的http服务编程的相关知识点。

`constants.go`

```go
package constants

const (
	ADDR         = "localhost:9527"
	CONTEXT_PATH = "/"
)
```

`http_server.go`

```go
package main

import (
	"go_code/web/http/constants"
	"net/http"
	"strings"
)

var handlerMap = map[string]http.HandlerFunc{}

func main() {

	handlerMap["/sayHello"] = SayHelloHandler
	handlerMap["/goodbye"] = SayGoodbyeHandler

	http.HandleFunc(constants.CONTEXT_PATH, BaseHandler)

	err := http.ListenAndServe(constants.ADDR, nil)

	if err != nil {
		println("启动http服务器错误\n", err.Error())
	}

}

func BaseHandler(w http.ResponseWriter, r *http.Request) {

	uri := strings.Split(r.RequestURI, "?")[0]

	println("请求URI ", uri)
	//println(i)

	handlerFunc, ok := handlerMap[uri]
	if ok {
		handlerFunc(w, r)
	} else {

		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("请求资源不存在"))

	}
}

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {

	// 获取get请求参数
	name := r.URL.Query().Get("name")

	if name != "" {
		_, _ = w.Write([]byte("hello ! " + name))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("缺少参数name"))
	}
}

func SayGoodbyeHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name != "" {
		_, _ = w.Write([]byte("goodbye ! " + name))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("缺少参数name"))
	}
}
```

在这个例子当中，`http.ListenAndServe`函数主要用来监听并启动http服务，如果启动出现异常，则会返回相应的error。`http.HandleFunc`函数用于注册对应请求路径的处理函数，在这个例子当中，我们对`/`路径注册了对应的处理函数`BaseHandler`,它实现于http包中的`Handler`接口。在BaseHandler接口当中我们再对请求进一步的转发到具体的处理器函数。

而具体的处理函数也很简单，分别是一个sayHello和sayGoodbye的模拟业务处理的函数。当程序启动后，使用`curl`进行测试

程序输出：

```shell
curl http://localhost:9527/sayHello?name=韩立
hello ! 韩立
curl http://localhost:9527/sayHello
缺少参数name
curl http://localhost:9527/sayHell
请求资源不存在
curl http://localhost:9527/goodbye?name=陈师姐
goodbye ! 陈师姐
```

## 写在最后

在这篇文章当中，我们主要通过两个例子初步地认识了go语言网络编程的相关知识点。本文当中涉及到的例子可以[点击此处下载](https://github.com/bingfenglai/golang)。如果我的学习笔记能够给你带来帮助，还请多多点赞鼓励。文章如有错漏之处还请各位小伙伴帮忙斧正。







