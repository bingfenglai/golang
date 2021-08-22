package main

import (
	"fmt"
	"sort"
)

func main() {

	orderlyTraversal()

}


func traverse1() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"


	for key, value := range m {
		fmt.Println(key,value)
	}

}

func traverse2() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	for s := range m {
		fmt.Println(s)
	}

}


func orderlyTraversal(){

	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	var keyArray []string

	for key := range m{
		keyArray = append(keyArray, key)
	}
	sort.Strings(keyArray)

	for i := 0; i < len(keyArray); i++ {
		fmt.Println(keyArray[i],m[keyArray[i]])
	}




}
