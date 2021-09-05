package main

import (
	"fmt"
)

func main() {
	sayHello("厉飞羽")
	s1 := getPhoneByName("南宫婉")
	fmt.Println(s1)
	s2 := getPhoneByName2("南宫婉")
	fmt.Println(s2)
	sayHello3("厉飞羽","韩立")
}

func sayHello(name string) {
	fmt.Println("你好！",name)
}

func getPhoneByName(name string) string {

	return name+" 的手机号是 " + "139-6379-3306"
}


func getPhoneByName2(name string) (str string){
	str = name+" 的手机号是 " + "139-6379-3306"
	return
}



func sayHello2(name1 ,name2 string)  {
	fmt.Println("你们好！",name1,name2)
}

func sayHello3(name ...string) {
	for _, s := range name {
		fmt.Println("你好！",s)
	}
}
