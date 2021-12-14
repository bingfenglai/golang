package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	ch := make(chan interface{}, 3)

	for i := 0; i < 10; i++ {
		go buyGoods(ch)
	}

	time.Sleep(3 * 1e9)

}

func buyGoods(ch chan interface{}) {
	println("准备下单")

	ch <- 1
	println("购买成功！库存剩余", cap(ch)-len(ch))
}
