package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copyFile("古丹丹方.txt", "给韩立的玉简.txt")
	fmt.Println("拷贝完成！")

}

func copyFile(source, target string) {

	openFile, err := os.Open(source)
	if err != nil {
		fmt.Println("打开源文件出错")
		return
	}
	defer openFile.Close()

	createFile, err := os.Create(target)
	if err != nil {
		fmt.Println("创建目标文件出错")
		return
	}
	defer createFile.Close()

	written, err := io.Copy(createFile, openFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(written)

}
