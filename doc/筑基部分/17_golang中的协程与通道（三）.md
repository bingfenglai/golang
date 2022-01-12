# golang中的协程与通道（三）

## 写在前面

在前一篇文章《[golang中的协程与通道（二） ](https://bingfenglai.github.io/2021/12/15/golang/16_golang中的协程与通道（二）/)》当中，我们初步学习了通道的相关内容，并结合协程进行了简单的应用，接下来，我们将通过具体的练习来进一步学习该知识点。



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

我们知道，在硬件资源不变的情况下，计算机处理的最大请求数是固定。当某些极限情况下，可能会出现大量的请求打到计算机上，这时，可能会出现请求过载，导致计算机宕机。因此，我们需要将请求最大值限制在计算机能够处理的范围之内，也就是`限流`。在`Golang`中使用带有缓冲的通道就很容易实现这一点，**通道的缓冲大小就是同时处理请求的最大数量。**

我们将上面的例子当中的`service`通道修改一下，在初始化时指定缓冲大小

```go
// 启动服务器方法
// resp: 请求通道、退出信号通道
func startServer(do sayHello) (seivice chan *request, quit chan int) {
	seivice = make(chan *request,4)
	quit = make(chan int)
	go server(do, seivice, quit)
	return seivice, quit

}
```

这时，我们再次运行一下demo,就可以看到此时服务器函数同时接收到请求就扩大到了4.

## 并行计算的应用

现代计算机绝大多数都是多核心CPU。并行计算可以理解为：利用多个处理器协同求解同一问题的过程。接下来，我们通过一个demo了解Golang在并行计算上面的应用。

```go
package main

import (
	"runtime"
	"strconv"
)

// 数列求和函数
func sum(list []int, ch chan int) error {
	i := 0

	for _, j := range list {
		i = i + j
	}
	println("分段结果:" + strconv.Itoa(i))
	ch <- i
	return nil
}

// 参与运算的cpu数
const cpu_num = 6

func main() {
	runtime.GOMAXPROCS(cpu_num)
	var list = []int{1, 22, 31, 34, 52, 46, 87, 18, 91, 101, 161, 182}
	ch := make(chan int, 2)
	// 分为六等份
	num := cap(list) / cpu_num
	println("每份大小：" + strconv.Itoa(num))
	for i := 0; i < cpu_num; i++ {
		sub := list[i*num : i*num+num]
		
		go sum(sub, ch)

	}
	var total []int
	for i := 0; i < cpu_num; i++ {
		sum := <-ch
		println("收到结果" + strconv.Itoa(sum))
		total = append(total, sum)
	}
	println("结果长度", len(total))
	for i := 0; i < len(total)-1; i++ {
		println(total[i])
	}

	println("==========")
	go sum(total, ch)

	println(<-ch)

}

```

输出：

```go
每份大小：2
分段结果:23
收到结果23
分段结果:98
收到结果98
分段结果:343
分段结果:192
分段结果:65
收到结果343
收到结果192
收到结果65
分段结果:105
收到结果105
结果长度 6
23
98
343
192
65
==========
分段结果:826
826
```

在这个例子当中，存在一个长度为N的数组，我们需要对数组内的元素（元素是无规律的）进行求和。我们将其分为n（n为N的公约数且n小于计算机CPU核心数）等份，并通过协程去同时计算n等分的分段和，最后进行再对分段和组成的数组（可以再次切分为m等分，依次类推）进行求和。

在这个例子当中，我们可以充分的利用资源进行协同求解，使得计算效率大大提升。

## 通过通道来访问共享资源

在这之前，我们为了安全地访问共享资源，一般通过加锁来实现。

```go
type Info struct {
    mu sync.Mutex
    // ... other fields, e.g.: Str string
}
```

我们知道，对于通道内的元素，我们将其取出是具备顺序性的，因此，我们可不可利用此特性来实现对共享资源的访问呢？请看下面的例子：

```go
package main

import (
	"runtime"
	"time"
)
// 共享资源
type Count struct {
	count int
	funCh chan func()
}

// 工厂函数
func NewCount(i int) *Count {

	count := &Count{

		count: i,
		funCh: make(chan func()),
	}
	go count.backend()

	return count
}

// 后台协程方法
func (receiver *Count) backend() {
	for {
		f := <-receiver.funCh
		f()

	}
}

// 访问资源的方法
func (receiver *Count) AddCount(count int) {
	f := func() {
		receiver.count = receiver.count + count
	}
	receiver.funCh <- f
}

func main() {

	runtime.GOMAXPROCS(2)
	count := NewCount(0)

	total := 0

	go func(count *Count) {
		for i := 0; i < 500; i++ {
			count.AddCount(1)
			total = total + 1
		}
	}(count)

	go func(count *Count) {
		for i := 0; i < 500; i++ {
			count.AddCount(-1)
			total = total - 1

		}
	}(count)

	time.Sleep(1 * 1e9)
	println("通过通道访问的资源：", count.count)
	println("直接访问的资源：", total)

}
```

程序输出：

```go
通过通道访问的资源： 0
直接访问的资源： -2
```



在这个例子当中，我们分别通过通道和直接访问的方式操作资源，最后打印结果。可以看到，通过通道去组织对资源的访问，可以起到对资源加锁的作用。当然，这仅仅是一个简化的demo，虽然不能直接用于实际开发，但是这种方式给我们提供了在实际场景中并发编程对于资源访问方面提供了思路。

## 写在最后

在这篇文章当中，我们通过几个小案例跟大家探讨了关于协程、通道的应用，使得我们可以更好地掌握这些知识点。

本文当中涉及到的例子可以[点击此处下载](https://github.com/bingfenglai/golang)。如果我的学习笔记能够给你带来帮助，还请多多点赞鼓励。文章如有错漏之处还请各位小伙伴帮忙斧正。







