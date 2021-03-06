# golang中的协程与通道（一）

## 写在前面

在上一篇文章[《golang中的错误处理》](https://bingfenglai.github.io/2021/12/11/golang/14_golang%E4%B8%AD%E7%9A%84%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86/)当中，我们简单介绍了Golang当中的错误处理部分的内容。接下来，我们将学习Golang当中的**协程**（goroutine）和**通道**（channel）两部分的内容。

## 概述

作为一门 21 世纪的语言，Go 原生支持应用之间的通信和程序的并发。程序可以在不同的处理器和计算机上同时执行不同的代码段。Go 语言为构建并发程序的基本代码块是 **协程** (go-routine) 与**通道** (channel)。

## 并发、并行和协程

### 协程的定义

在讲协程的定义之前，我们首先需要来了解什么是并发和并行，以及他们之间的区别。

一般来说，一个应用程序运行在机器上时，它对应着一个进程。一个进程由一个或多个操作系统线程组成，这些线程共享进程的内存地址空间。那么我们说，**一个并发程序可以在一个处理器或者内核上使用多个线程来执行任务，但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行。**

公认的，使用多线程的应用难以做到准确，最主要的问题是内存中的数据共享，它们会被多线程以无法预知的方式进行操作，导致一些无法重现或者随机的结果（称作 `竞态`）。那么解决这个问题的方法便是对数据加锁。这样同时就只有一个线程可以变更数据。在 Go 的标准库 `sync` 中有一些工具用来在低级别的代码中实现加锁。

为了方便大家理解加锁这个问题，下面给大家举一个形象的例子：

例如，在现实生活当中，妹子可以类比做数据。为了数据访问合法性，假设有一线程男要对某一数据女进行操作，那么在这之前，他需要带着妹子先去领证。那么领完证之后，就类比于对数据加了锁（排他锁、独占锁）。那么线程对数据才能进行操作，其他线程想要操作这个数据，那么就得等到原先的线程释放锁。而释放锁的过程正好与加锁的过程相反，就可以类比于离婚了。

上面的例子当中只是为了方便部分读者去理解锁的相关概念，例子当中涉及到的只有独占锁一类锁。实际编程当中还涉及到了读写锁、更新锁、乐观锁、悲观锁等等概念，关于这部分，大家可以去Google，这并不是今天的重点。

在过去的软件开发经验告诉我们这会带来更高的复杂度，更容易使代码出错以及更低的性能，所以这个经典的方法明显不再适合现代多核/多处理器编程。

在Golang当中，应用程序并发处理这部分被称作**go-routines**（协程或者go协程）（需要注意的是：这里的协程与Python当中的协程概念并不相同，这是两个完全不同的东西。），并鼓励开发者使用channels进行协程同步（后面会详细讲），而不是通过`sync`包当中的锁来实现。

### 协程的特点

1. 协程与操作系统线程之间并没有一一对应的关系：协程是根据一个或多个线程的可用性，映射（多路复用）执行于他们之上的。
2. 当系统调用（比如等待 I/O）阻塞协程时，其他协程会继续在其他线程上工作。协程的设计隐藏了许多线程创建和管理方面的复杂工作。
3. 协程是轻量的，它比线程更轻。使用4K的栈内存就可以在堆当中创建它们。栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。
4. 协程可以运行在多个操作系统线程之间，也可以运行在线程之内，让你可以很小的内存占用就可以处理大量的任务。即使有协程阻塞，该线程的其他协程也可以被`runtime`调度，转移到其他可运行的线程上。并且这个细节对于开发者来说是透明的，降低了编程的难度。
5. 线程是运行协程的实体，协程调度器将可运行的协程分配到工作线程上。

### 协程的调用

协程是通过使用关键字 `go` 调用（执行）一个函数或者方法来实现的（也可以是匿名或者 lambda 函数）。这样会在当前的计算过程中开始一个同时进行的函数，在相同的地址空间中并且分配了独立的栈，比如：`go sum(bigArray)`，在后台计算总和。

```go
go 函数名（实参列表）
```

### GOMAXPROCS参数的使用

在上文中提到：线程是运行协程的实体，协程调度器将可运行的协程分配到工作线程上。那么，如何设置多少系统线程用于执行协程呢？这时就需要用到**GOMAXPROCS**参数。GOMAXPROCS参数默认值为1。这时，程序的所有协程都由1个线程执行，也就是N-1模式（关于其他几种模式，后面有时间再详细讲解）。

在N-1模式下，协程在用户态线程即完成切换，不会陷入到内核态，这种切换非常的轻量快速。但是缺点也很明显：无法使用多核加速能力，一旦某协程阻塞，就会造成线程阻塞。也因此，我们需要通过设置该参数，来充分利用多核CPU。

假设 n 是机器上处理器或者核心的数量。如果你设置环境变量 GOMAXPROCS>=n，或者执行 `runtime.GOMAXPROCS(n)`，接下来协程会被分割（分散）到 n 个处理器上。更多的处理器并不意味着性能的线性提升。有这样一个经验法则，对于 n 个核心的情况设置 GOMAXPROCS 为 n-1 以获得最佳性能，也同样需要遵守这条规则：协程的数量 > 1 + GOMAXPROCS > 1。

一句话概括：GOMAXPROCS参数值等同于（并发的）线程数量，在一台核心数多于1个的机器上，会尽可能有等同于核心数的线程在并行运行。

### 协程的简单应用demo

下面，我们将通过一个简单的的demo来实际应用协程。

```go
package main

import (
	"runtime"
	"time"
)

func main() {
	i := make([]int,0)
	i = append(i, 2,4,6)
	runtime.GOMAXPROCS(2)
	go sum(i)
	go multiply(i)
	// 为了保证协程逻辑执行完
	time.Sleep(1 * 1e9)
}

// 求和函数
func sum(num []int) {
	var total int
	for i, _ := range num {

		total = num[i]+total
	}

	println("求和结果为：",total)

}

// 乘积
func multiply(num []int) {
	var total int
	total = 1
	for i, _ := range num {

		total = num[i]*total
	}
	println("乘积结果为：",total)
}



```

程序输出：

```
乘积结果为： 48
求和结果为： 12
```

在这个例子当中，我们需要对输入的一组数据分别进行求和运算与乘积运算。这两个运算可以同时进行。我们通过创建2个协程对其分别进行乘积与求和运算并打印最终的结果。

等我们学习完通道的相关知识点后，我们将对其进行一个综合的应用。

## 写在最后

在这篇文章当中，我们初步认识了Go语言当中的协程，并通过一个简单的demo跟大家分享协程的使用。在下一篇文章当中，我们将介绍通道的相关知识点。

本文当中涉及到的例子可以[点击此处下载](https://github.com/bingfenglai/golang)。如果我的学习笔记能够给你带来帮助，还请多多点赞鼓励。文章如有错漏之处还请各位小伙伴帮忙斧正。





