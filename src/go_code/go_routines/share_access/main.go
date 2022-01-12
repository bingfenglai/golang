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
