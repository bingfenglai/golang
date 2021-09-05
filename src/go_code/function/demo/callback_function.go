package main

import "fmt"

func main() {
	name := "喜小乐"
	sayHello(name,callback)
}

func callback(name string) {
	fmt.Println("hello ",name)
}

func sayHello(name string,f func(name string)) {
	fmt.Println("我是",name)
	f(name)
}
