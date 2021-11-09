package main

import "fmt"

func main() {


}

func check(e error) {
	if e!=nil {
		panic(e)
	}
}


func errorHanlder(f1 func(s string))  {
	defer func() {
		if r := recover();r!=nil {
			fmt.Println(r)
		}
	}()

}


