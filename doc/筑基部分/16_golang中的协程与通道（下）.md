# golang中的协程与通道（下）

## 写在前面

在上一篇文章[《golang中的协程与通道（上）》](https://bingfenglai.github.io/2021/12/12/golang/15_golang%E4%B8%AD%E7%9A%84%E5%8D%8F%E7%A8%8B%E4%B8%8E%E9%80%9A%E9%81%93/)当中，我们初步认识了golang当中的协程的相关知识点。接下来，我们将开始学习通道(channel) 的相关知识点。

## 通道的概念

在上一篇文章的demo当中，协程是独立执行的，他们之间没有进行通信。然而在实际情况下，协程之间必须要通信才会变得更加有用：协程之间通过发送和接收消息来协调或同步他们之间的工作。

一组协程组成一条流水线，他们通过皮带流水装配线来协同工作，以提升资源利用率和工作效率。

Golang中有一种特殊的类型，通道（channel），它是一个可以用于发送类型化数据的管道，由其负责协程之间的通信，通过这种通信方式从而保证了同步性。

## 通道的声明

通道的声明方式如下：

```go
var 通道标识符 chan datatype
```

未初始化的通道的值是`nil`.

从上面的声明语句当中我们知道，通道只能传输一种类型的数据，并且所有的数据类型都可用于通道。interface{}类型也是可以的。

实际上，通道是类型化消息的队列：它是先进先出（FIFO）的结构的，这保证了数据传输的顺序性。

通道是一种引用类型，因此我们可以使用make()函数来给它分配内存。

```go
var ch1 chan string
ch1 = make(chan string)
// 或者
ch1 := make(chan string)
```

通道是第一类对象：可以存储在变量中，作为函数的参数传递，从函数返回以及通过通道发送它们自身。另外它们是类型化的，允许类型检查，比如尝试使用整数通道发送一个指针。

## 通信操作符<-

通信操作符`<-`箭头的方向为数据的流向。

流向通道：

```go
ch :=make(chan int)
ch <- int1
```

从通道流出：

```go
int2 := <- ch
```

`<- ch` 可以单独调用获取通道的（下一个）值，当前值会被丢弃:

```go
if <- ch !=-1{
    do something
}
```

## 通道的特点

1. 为了可读性通道的命名通常以 `ch` 开头或者包含 `chan`。
2. 通道的发送和接收都是原子操作：它们总是互不干扰的完成的。

## 举个例子

在下面的这个demo当中，我们结合前面学到的协程实际运用一下通道。

```go
package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	ch := make(chan string)
	ch1 :=make(chan string)

	go sendData(ch)
	go sayHello(ch,ch1)

	for <-ch1!="EOF" {

	}


}

func sendData(ch chan string) {

	ch <- "韩立"
	ch <- "厉飞羽"
	ch <- "张铁"
	ch <- "墨大夫"
	ch <- "南宫婉"
	ch <- "六道传人"
	ch <- "董萱儿"
	ch <-"EOF"
}

func sayHello(ch,ch1 chan string){

	for {
		input := <-ch

		if input!="EOF" {
			println("hello !",input)
		}else {
			break
		}

	}

	ch1 <- "EOF"
}
```

程序输出：

```
hello ! 韩立
hello ! 厉飞羽
hello ! 张铁
hello ! 墨大夫
hello ! 南宫婉
hello ! 六道传人
hello ! 董萱儿
```

在这个例子当中，`sendData()`函数向通道`ch`发送数据,`sayHello()`函数接受并处理，处理完数据后向`ch1`通道发送结束符，主程序退出。在这个例子当中，很好地展示了通道以及协程的综合使用。



