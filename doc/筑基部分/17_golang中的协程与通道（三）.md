# golang中的协程与通道（三）

## 写在前面

在前一篇文章《[golang中的协程与通道（二） ](https://bingfenglai.github.io/2021/12/15/golang/16_golang中的协程与通道（二）/)》当中，我们初步学习了通道的相关内容，并结合协程进行了简单的应用，接下来，我们将进一步学习该知识点。



## 使用 select 切换协程

从不同的并发执行的协程中获取值可以通过关键字`select`来完成，它和`switch`控制语句非常相似，也被称作通信开关。

```go
package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case i := <-ch1:
		fmt.Println(i)
	case j := <-ch2:
		fmt.Println(j)
	default:
		fmt.Println(" default")

	}
}
```

关于`select` 使用的说明：

1. `default` 语句是可选的；

2. 在任何一个 case 中执行 `break` 或者 `return`，select 就结束了；

3. `select` 做的就是：选择处理列出的多个通信情况中的一个：

   3.1 如果通道都属于阻塞状态，则会等待直到其中一个可以处理（没有default的情况下）

   3.2 如果多个可以处理，则随机选择其中一个

   3.3 如果没有通道操作可以处理了并且写了`default` 语句，则会执行`default`语句：`default` 永远是可运行的。

## 惰性生成器的实现

生成器是指被调用时返回下一个序列中下一个值的函数。一般用来生成订单号等。因为生成器每次返回的是序列中下一个值而非整个序列，这种特性也被称之为**惰性求值**。

请看下面的例子：

```go
package interval

import (
	"errors"
	"sync"
)

type Interval struct {
	maxValue    int
	currentVale int
	step        int
	lock        sync.Mutex
}

func NewInterval(maxValue, step int) *Interval {

	return &Interval{
		maxValue:    maxValue,
		currentVale: 1,
		step:        step,
	}

}

// 获取序列的下一个区间
func (receiver *Interval) GetNextInterval() (min, max int,err error) {

	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	if receiver.currentVale >= receiver.maxValue {
		return 0, 0, errors.New("序列号已耗尽")
	}

	min = receiver.currentVale
	max = receiver.currentVale + receiver.step
	if max > receiver.maxValue {
		max = receiver.maxValue
	}
	receiver.currentVale = max + 1
	return min, max,nil
}

```

```go
package main

import (
	"fmt"
	"go_code/channel_demo/gen/interval"
	"runtime"
	"sync"
	"time"
)



// 声明一个chan int 类型的通道
var resume chan int

// 区间实例
var myInterval = interval.NewInterval(99999, 1000)

type count struct {
	i int
	mu sync.Mutex
}

func main() {
	runtime.GOMAXPROCS(10)

	// 初始化通道
	resume = intergers()

	count := count{
		i:  0,
	}

	// 不断地获取下一个序列值并格式化打印
	// 为了验证所有的协程获取的序列号不重复，每获取一个记录数count+1
	for j :=10;j>0;j-- {

		go func() {

			for {

				fmt.Println(fmt.Sprintf("%05d", generateInteger()))
				count.mu.Lock()
				count.i = count.i+1
				count.mu.Unlock()
			}
		}()

	}



	time.Sleep(1e9*3)
	println("总共获取：",count.i)



}

// 从chan int 通道当作获取下一个值
func generateInteger() int {
	return <-resume
}


// 获取一个chan int 类型的通道
func intergers() chan int {
	yield := make(chan int)
	min, max, _ := myInterval.GetNextInterval()
	var err error
	// 开启一个协程，不断地向通道当中写入下一个序列号（阻塞的）

	go func() {
		for {

			if min <= max {
				yield <- min
				min++
			} else {
				min, max, err = myInterval.GetNextInterval()
				if err!=nil {
					return
				}
			}

		}

	}()

	return yield
}
```

在这个例子当中，Interval存储的是整个序列的信息，包括序列的最大值，步长，当前开始值。函数`GetNextInterval()` 获取下一个区间值。函数`generateInteger()`获取序列当中的下一个值。我们在main函数当中，使用协程和无限循环不断地获取序列值并打印，每获取一个`count`+1，最终打印获取到的总数。

程序输出：

```
总共获取： 99999
```

## 实现 Futures 模式

所谓的**Futures**指的是：在某些场景当中，我们使用某一个值之前需要先对其进行计算。在这种情况下，我们可以在另一个处理器上进行该值的计算，到使用时，该值就已经计算完毕。

Futures类似于生成器，不同的地方在于Futures需要返回一个值。

请看下面的例子：

在某一场景当中，我们的程序需要不断地接收图片并对其进行特征提取。那么，我们应当如何实现呢？

实现方式一：

```go
package main

func main() {

	file := receiverImagesFile()
	resolve(file)

}

func receiverImagesFile()  string{
	return "1.png"
}

func resolve(s string) {
	println("对图片 "+s+" 进行特征提取")
}
```

在这个方式当中，接收文件跟对文件进行处理采用同步的方式进行，分为接收函数`receiverImagesFile()`和处理函数`resolve（）` 。在这个简单的模式当中，因为是同步的，当接收函数正在接收文件时（IO相对而言是缓慢的），处理函数可能会存在空闲状态，这样会导致资源的浪费。

那么，Futures模式是如何处理的呢？请看下面的例子：

```go
package main

import (
	"time"
)

func main() {
	resolve()

}

func resolve() {

	s := <-receiverImageFile()
	time.Sleep(1e9 * 0.3)
	println("文件" + s + "处理完毕")

}

func receiverImageFile() chan string {
	ch := make(chan string)

	go func() {

		// 模拟接收文件过程
		println("接收文件中...")
		time.Sleep(1e9 * 1)
		ch <- time.Now().Format("20060102150405") + ".png"

	}()

	return ch
}
```

在Futures当中，接收函数跟处理函数通过通道进行解耦。接收函数可以不断地接收新的图片文件，而处理函数不断地对图片进行运算。再给通道设置缓存后，两者之间可以（理想情况下）做到互不影响，使得资源利用率得到提升。并且，我们可以配置多数接收函数对应少数的处理函数，进而减少计算资源的等待时间。对于密集计算型任务，Futures模式可以使得API以异步的形式暴露出来，使得API调用方可以在等待结果的时间处理其他任务。

## 协程与通道在CS模式中的应用

客户端（client）可以是运行在任意设备上的任意程序，它会按需发送请求（rrequest）到服务器。服务器（server）接收到请求后开始对应的工作，并将结果（也称之为响应，response）返回给客户端。一般情况下，多个客户端对应少数的服务器。日常我们使用的浏览器，移动App等都属于客户端。

在Golang当中，请求的执行一般在协程当中进行。请求的响应将通过请求当中包含的通道将结果进行返回。而server程序将不断地从通道当中接收请求，并开启一个协程对其进行处理。

请看下面的demo:

```go
package main

import (
	"runtime"
	"strconv"
	"time"
)

// 请求
type request struct {
	args   string
	replyc chan string
}

type sayHello func(name string) string

// 模拟调度具体的业务方法，并通过通道返回结果
func run(do sayHello, req *request) {
	req.replyc <- do(req.args)
}

// 模拟服务器应用
func server(do sayHello, service chan *request, quit chan int) {
	defer func() {
		close(quit)
		for  {
			if cap(service) == 0 {
				close(service)
				println("程序退出")
				return
			}

		}
	}()
	for {
		select {
		case req := <-service:
			println("收到请求", req.args)
			go run(do, req)

		case <-quit:
			println("收到退出指令")
			return
		}

	}
}

// 启动服务器方法
// resp: 请求通道、退出信号通道
func startServer(do sayHello) (seivice chan *request, quit chan int) {
	seivice = make(chan *request)
	quit = make(chan int)
	go server(do, seivice, quit)
	return seivice, quit

}

func main() {
	runtime.GOMAXPROCS(4)
	service, quit := startServer(func(name string) string {
		return "hello! " + name
	})

	const N = 100

	var reqs [N]request

	// 初始化N个请求实例，并对服务器发器请求
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.args = strconv.Itoa(i) + " name"
		req.replyc = make(chan string)

		// 模拟接收响应结果
		go func() {
			s := <-req.replyc
			println("接收响应：",s)

		}()
		service <- req
		if i == N-1 {
			quit<-1
		}

	}

	time.Sleep(2 * 1e9)

}

```

在这个例子当中，我们封装了请求`request` ，里面包含了请求参数和结果返回的通道信息。在服务器当中，每收到一个请求就开启一个协程对该请求进行处理（go协程是非常轻量的），并将结果发送到请求的响应通道当中。服务器通过`select`机制进行不同通道之间的切换。



## 通道在限制并发当中的应用

