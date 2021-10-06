package main

import "fmt"

func main() {

	fmt.Println("发生程序崩溃前")
	panic("程序崩溃了")
	fmt.Println("已产生panic...")
}
