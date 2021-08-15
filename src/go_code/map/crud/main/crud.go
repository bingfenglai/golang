package main

import "fmt"

func main() {
	//create()
	//update()
	//deleteMap()
	//find()
	traverse2()
}

func create() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	delete(m, "小王")

	fmt.Println(m)
}

func update() {

	var m =  map[string]string {
		"小王":"一听家就住在不远",
	}

	fmt.Println("修改前：",m)
	m["小王"] = "家住在隔壁"

	fmt.Println("修改后",m)
}

func deleteMap() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	fmt.Println("删除前： ",m)

	delete(m, "小王")

	fmt.Println("删除后： ", m)

}

func find () {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	value,ok := m["老王"]

	if ok {
		fmt.Println(value)
	}

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