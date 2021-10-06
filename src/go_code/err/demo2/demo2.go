package main

import (
	"fmt"
)

func main() {

	hello := sayHello("")

	fmt.Println(hello)
}

func sayHello(name string) string {

	if "" == name {
		panic("name 不能是一个空串")
	}

	return "Hello " + name
}
