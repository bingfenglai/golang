package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case i := <-ch1:
		fmt.Println(i)
	case j := <-ch2:
		fmt.Println(j)
	default:
		fmt.Println(" default")

	}
}
