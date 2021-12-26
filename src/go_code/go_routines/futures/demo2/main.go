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
