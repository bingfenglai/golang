package main

import "fmt"

func main() {
	var firstname, lastname string
	fmt.Println("请输入您的姓名：")
	_, _ = fmt.Scanln(&firstname, &lastname)
	fmt.Printf("你好！%s · %s\n", lastname, firstname)

}
