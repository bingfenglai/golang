package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("done...")
		if err := recover(); err != nil {
			fmt.Println(err)
			func() {
				fmt.Println("end...")
			}()
		}
	}()

	fmt.Println("start...")
	panic("this is a error")
}
