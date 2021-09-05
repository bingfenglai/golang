package main

import "fmt"

func main() {

	fun := func (name string){
		fmt.Println("Hello",name)
	}

	fun("向北")
}


