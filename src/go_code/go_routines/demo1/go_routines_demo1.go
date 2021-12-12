package main

import (
	"runtime"
	"time"
)

func main() {
	i := make([]int, 0)
	i = append(i, 2, 4, 6)
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

		total = num[i] + total
	}

	println("求和结果为：", total)

}

// 乘积
func multiply(num []int) {
	var total int
	total = 1
	for i, _ := range num {

		total = num[i] * total
	}
	println("乘积结果为：", total)
}
