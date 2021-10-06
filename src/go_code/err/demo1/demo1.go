package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	hello, err := sayHello("")
	if err != nil {
		log.Default().Println(err)
		return
	}
	fmt.Println(hello)
}

func sayHello(name string) (string, error) {

	if "" == name {
		return "", errors.New("name 不能是一个空字符串")
	}

	return "Hello " + name, nil
}
