package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	for _, arg := range flag.Args() {
		sayHello(arg)
	}

}

func sayHello(name ...string) {

	fmt.Println("hello!", name)
}
