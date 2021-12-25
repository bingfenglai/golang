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



