package main

import (
	"errors"
	"fmt"
)

var f1 = func(name string) error {
	if name != "" {
		println("Hello! ", name)

	} else {
		return errors.New("姓名不能为空")
	}

	return nil
}

func main() {

	errorHandler("", f1)

}

func errorHandler(name string, f1 func(s string) error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	err := f1(name)
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
