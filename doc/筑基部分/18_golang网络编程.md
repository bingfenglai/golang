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
		fmt.Printf("服务器启动完成 " + time.Now().Format("2006-01-02 15:04:05"))
	}

	for {

		accept, err := listen.Accept()
		if err != nil {
			println(err)
			return

		}

		go doServerStuff(accept)

	}

}

func doServerStuff(conn net.Conn) {

	for {
		buf := make([]byte, 1024)

		read, err := conn.Read(buf)

		if err != nil {
			println(err)
			return

		}

		println("\n 接收到数据：\n", string(buf[:read]))

		conn.Write([]byte(" hello client " + conn.RemoteAddr().String()))
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

	if err != nil {
		println("建立连接失败", err)
		return

	} else {
		fmt.Println("建立连接成功")
	}
	conn.Write([]byte("hello server " + time.Now().Format("2006-01-02 15:04:05")))


	buf := make([]byte, 1024)
	read, err := conn.Read(buf)

	println("收到服务器发送来的消息")

	if err != nil {
		println("读取数据失败 ", err)
		return
	}

	fmt.Println(string(buf[:read]))

}
```

客户端启动后，与服务器建立连接，并向服务器发送数据。

服务器程序输出：

```go
服务器启动完成 2022-01-16 16:06:56
 接收到数据：
 hello server 2022-01-16 17:52:41
```



客户端程序输出：

```go
建立连接成功
收到服务器发送来的消息
 hello client 127.0.0.1:2722

进程完成，并显示退出代码 0
```









