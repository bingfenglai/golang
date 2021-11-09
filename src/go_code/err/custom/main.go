package main

import (
	"fmt"
)

func main() {
	var name []string
	name = append(name, "韩立","","南宫婉")

	for i, _ := range name {

		err := SayHello(name[i])
		if err != nil {
			fmt.Println(err)
		}
		continue
	}




}

func SayHello(name string) (err error) {
	defer func()  {
		if r := recover(); r != nil {
			//var ok bool
				err = fmt.Errorf("%v",r)
		}
	}()
	// 注意： 抛出panic的函数必须在defer之后调用
	doSayHello(name)
	return nil

}

func doSayHello(name string) {
	if len(name)==0 {
		panic("名字不能是一个空字符串")

	}
	fmt.Printf("hello %s\n", name)


}
