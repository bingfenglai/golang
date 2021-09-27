package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("古丹丹方.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("打开文件出错")
		return
	}
	defer file.Close()

	// 创建一个缓冲区
	writer := bufio.NewWriter(file)

	s := "Hello World\n"

	for i := 0; i < 10; i++ {
		// 写入缓冲区
		_, _ = writer.WriteString(s)
	}
	// 将缓冲区的数据写入文件
	_ = writer.Flush()
}
