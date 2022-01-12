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
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ch := make(chan int, 2)
	// 分为六等份
	num := cap(list) / cpu_num
	println("每份大小：" + strconv.Itoa(num))
	for i := 0; i < cpu_num; i++ {
		sub := list[i*num : i*num+num]
		//println("第"+ strconv.Itoa(i) +"等份")
		//for _, i2 := range sub {
		//	println(i2)
		//}

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
