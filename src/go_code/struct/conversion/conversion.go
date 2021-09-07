package main

import "go_code/struct/model"

func main() {
	var a = myNumber{Value: 18.0}
	b := model.Number(a)
	println(b.Value)
}

type myNumber model.Number
