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
	i  int
	mu sync.Mutex
}

func main() {
	runtime.GOMAXPROCS(10)

	// 初始化通道
	resume = intergers()

	count := count{
		i: 0,
	}

	// 不断地获取下一个序列值并格式化打印
	// 为了验证所有的协程获取的序列号不重复，每获取一个记录数count+1
	for j := 10; j > 0; j-- {

		go func() {

			for {

				fmt.Println(fmt.Sprintf("%05d", generateInteger()))
				count.mu.Lock()
				count.i = count.i + 1
				count.mu.Unlock()
			}
		}()

	}

	time.Sleep(1e9 * 2)
	println("总共获取：", count.i)

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
				if err != nil {
					return
				}
			}

		}

	}()

	return yield
}
