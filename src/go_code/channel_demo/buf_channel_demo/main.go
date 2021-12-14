package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	// 缓冲元素个数
	buf := 3
	ch := make(chan string, buf)

	ch1 := make(chan string)
	go sendData(ch)
	go receivingData(ch, ch1)

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
	fmt.Println("开始发送数据", time.Now().Format("2006-01-02 15:04:05"))
	for i, _ := range data {

		ch <- data[i]
		fmt.Println("发送数据：", data[i])
		if data[i] == "EOF" {
			fmt.Println("数据发送完毕", time.Now().Format("2006-01-02 15:04:05"))
			break

		}

	}

}

func receivingData(ch, ch1 chan string) {
	// 为了演示缓冲效果，先让接收者函数休眠3秒
	time.Sleep(3 * 1e9)
	fmt.Println("开始接收数据", time.Now().Format("2006-01-02 15:04:05"))
	for {
		input := <-ch
		if input != "EOF" {
			fmt.Println("接收到数据：", input)
			fmt.Println("数据处理中...")
			// 模拟数据处理耗时
			time.Sleep(1 * 1e9)
		} else {
			fmt.Println("数据接收完毕", time.Now().Format("2006-01-02 15:04:05"))
			break
		}

	}

	ch1 <- "EOF"
}
