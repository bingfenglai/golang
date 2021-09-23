package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("请输入一段文字")
	inputReader := bufio.NewReader(os.Stdin)
	s, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("你输入的是：")
		fmt.Println(s)
	}

}
