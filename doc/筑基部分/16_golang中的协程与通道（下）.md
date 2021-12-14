# golang中的协程与通道（二）

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

	<-ch1


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

## 通道阻塞

在上面的例子我们可以发现：在默认情况下，使用通道通信是同步且无缓冲的,必须要一个接收者准备好接收通道的数据后发送者才可以将数据发送给接收者，在这之前，通道是阻塞的。

在默认情况下：

1. 对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。

2. 对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。

## 使用带缓冲的通道

下面，我们将通过一个例子来学习带缓冲通道的使用。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
    runtime.GOMAXPROCS(2)
	// 缓冲元素个数
	buf := 3
	ch := make(chan string,buf)

	ch1 := make(chan string)
	go sendData(ch)
	go receivingData(ch,ch1)

	<-ch1

}

func sendData(ch chan string) {

	data := [...]string{
		"韩立",
		"厉飞羽",
		"张铁",
		"墨大夫",
		"南宫婉",
		"六道传人",
		"董萱儿",
		"EOF",
	}
	fmt.Println("开始发送数据",time.Now().Format("2006-01-02 15:04:05"))
	for i, _ := range data {


		ch <- data[i]
		fmt.Println("发送数据：", data[i])
		if data[i] == "EOF" {
			fmt.Println("数据发送完毕",time.Now().Format("2006-01-02 15:04:05"))
			break

		}

	}

}

func receivingData(ch, ch1 chan string) {
	// 为了演示缓冲效果，先让接收者函数休眠3秒
	time.Sleep(3*1e9)
	fmt.Println("开始接收数据",time.Now().Format("2006-01-02 15:04:05"))
	for {
		input := <-ch
		if input != "EOF" {
			fmt.Println("接收到数据：", input)
			fmt.Println("数据处理中...")
			// 模拟数据处理耗时
			time.Sleep(1*1e9)
		} else {
			fmt.Println("数据接收完毕",time.Now().Format("2006-01-02 15:04:05"))
			break
		}

	}

	ch1 <- "EOF"
}
```

程序输出：

```go
开始发送数据 2021-12-14 22:25:24
发送数据： 韩立
发送数据： 厉飞羽
发送数据： 张铁
开始接收数据 2021-12-14 22:25:27
接收到数据： 韩立
数据处理中...
发送数据： 墨大夫
接收到数据： 厉飞羽
发送数据： 南宫婉
数据处理中...
接收到数据： 张铁
数据处理中...
发送数据： 六道传人
接收到数据： 墨大夫
数据处理中...
发送数据： 董萱儿
接收到数据： 南宫婉
数据处理中...
发送数据： EOF
数据发送完毕 2021-12-14 22:25:31
接收到数据： 六道传人
数据处理中...
接收到数据： 董萱儿
数据处理中...
数据接收完毕 2021-12-14 22:25:34


```

在这个例子当中，`ch`是一个缓冲大小为3的通道。这意味着数据发送方可以在接收方未准备好的情况下先往通道里面塞3个数据，等接收方拿第一个数据后发送方就可以继续往里面塞数据。

总结如下：

1. 以上demo中buf 是通道可以同时容纳的元素（这里是 string）个数

2. 在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。

3. 缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的 `cap` 函数可以返回缓冲区的容量。

4. 如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是0或者未设置，通信仅在收发双方准备好的情况下才可以成功。

## 信号量模式

在上面的例子当中，为了告诉数据接收方数据已经发送完了，双方约定好：当接收到的数据等于"EOF"符号时，表示数据已发送完毕。数据接收方处理完数据后通过通道`ch1`发送"EOF"告诉主程序数据处理完毕，使得主程序退出。这里的"EOF"就是一个信号.

除此之外，信号量还经常用以实现互斥锁，限制对资源的并发访问。

请看下面的例子：

```go
package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	ch :=make(chan interface{},3)

	for i :=0;i<10;i++ {
		go buyGoods(ch)
	}

	time.Sleep(3* 1e9)

}

func buyGoods(ch chan interface{})  {
	println("准备下单")
	ch <- 1
	println("购买成功！库存剩余",cap(ch)-len(ch))
}
```

程序输出：

```go
准备下单
购买成功！库存剩余 2
准备下单
购买成功！库存剩余 1
准备下单
购买成功！库存剩余 0
准备下单
准备下单
准备下单
准备下单
准备下单
准备下单
准备下单
```

在这个例子当中，我们使用通道缓存元素个数来表示商品的库存，使用协程模拟并发访问，来模拟下单场景。可以看到使用通道缓存来实现信号量机制，可以保护我们的共享资源的并发访问。

除此之外，大家可以考虑一种计算场景：输入是一个很长的数据序列，我们要对这个序列求和，也就是1+2+...+n-1+n。这时，我们可以通过协程加通道的形式，使用并行计算的方式，计算好每个分段的和，然后将其发送到通道当中，主程序收到后进行最后的运算，这样可以极大的提高资源利用率，加快运算速度。关于这一块，这里只是顺嘴提一句，大家可以尝试着实现一下。

## 通道的方向与习惯用法

### 通道的方向

通道类型可以用注解来表示它只发送或者只接收：

```go
// 只发送
var send_only chan<- int 
// 只接收
var recv_only <-chan int       
```

只接收的通道（<-chan T）无法关闭，因为关闭通道是发送者用来表示不再给通道发送值了，所以对只接收通道是没有意义的。

通道在创建时都是双向的，但是我们可以分配有方向的通道变量。

请看下面的例子：

```go
package main

func main() {
	ch :=make(chan string)
	go sendData(ch)
	go receivingData(ch)
	
}


func sendData(ch chan<- string) {
	
}

func receivingData(ch <-chan string) {
	
}

```

### 习惯用法1：通道迭代器

```go
package main

type container struct {
	items []string
}

func (c *container) Len()int {
	return len(c.items)
}




func (c *container) Iter () <- chan string {
	ch := make(chan string)

	go func () {
		for i:= 0; i < c.Len(); i++{
			ch <- c.items[i]
		}
		close(ch)
	} ()
	return ch
}

func main() {
	c := container{items: []string{"韩立","南宫婉"}}

	for s := range c.Iter() {
		println(s)
	}

}

```

程序输出：

```go
韩立
南宫婉
```

以上就是给通道使用for循环实现的迭代器,其中`container`为存放资源的容器。使用for循环遍历通道，意味着它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。**写入完成后必须要关闭通道** 。因为Iter函数返回的是一个只读通道，它是没法关闭的。

### 习惯用法2 ：生产者消费者模式

假设存在生产者函数Produce()不断产生消费者consume()所需要的值，它们都可以运行在独立的协程中,那么你可以使用一下的写法:

```go
package main

func main() {
	
	Consume(Produce(10))
}

func Produce(size int) <-chan int{
	ch := make(chan int)
	go func() {
		for i:=0;i<size;i++ {
			ch<-i
		}
		close(ch)
	}()

	return ch
}

func Consume(ch <-chan int) {
	for i := range ch {
		println("收到数据",i)
	}
}
```

程序输出：

```go
收到数据 0
收到数据 1
收到数据 2
收到数据 3
收到数据 4
收到数据 5
收到数据 6
收到数据 7
收到数据 8
收到数据 9
```

## 写在最后

关于















