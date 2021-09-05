package main

import "fmt"

func main() {
	sayHello("南宫婉")
}

func sayHello(name string) {
	fmt.Println("before1")
	defer fun1()
	fmt.Println("before2")
}

func fun1() {
	fmt.Println("hello world!")
}
