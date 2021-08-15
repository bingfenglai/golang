package main

import "fmt"

func main() {

	map1()
	map2()
	map3()
}

func map1() {
	// 通过map关键字生命map
	var peopleMap map[string]string

	// 初始化
	peopleMap = make(map[string]string, 4)

	peopleMap["01"] = "喜小乐"
	peopleMap["02"] = "东小贝"
	peopleMap["03"] = "北小楠"
	peopleMap["04"] = "楠小茽"

	fmt.Println(peopleMap)


}

func map2() {

	m := make(map[string]string, 4)
	m["no1"] = "向北"
	m["no2"] = "向南"
	m["no3"] = "向东"
	m["no4"] = "向西"

	fmt.Println(m)

}

func map3() {

	m2 := map[string]string{
		"小明":"小朋友",
		"小张": "是个大人",
		"小李": "是个司机",
		"小王": "家住隔壁",
	}

	fmt.Print(m2)
}
