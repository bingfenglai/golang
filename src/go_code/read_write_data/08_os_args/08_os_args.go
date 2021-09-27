package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			sayHello(os.Args[i])
		}
	} else {
		fmt.Println("参数为空")
	}

}

func sayHello(name ...string) {

	fmt.Println("hello!", name)
}
