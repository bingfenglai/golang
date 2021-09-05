package main

import "fmt"

func main() {
	sayHello("南宫婉")
}

func sayHello(name string) {
	fmt.Println("before")
	defer fun1()
	fmt.Println("after")
}

func fun1() {
	fmt.Println("hello world!")
}
