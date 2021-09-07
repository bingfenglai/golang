package main

import "fmt"

type immortal struct {
	// 姓名
	string
	// 年龄
	int
	// 修仙境界
	level
}

type level struct {
	// 境界名称
	string
	// 灵气值
	float32
}

func main() {
	var im = immortal{
		string: "韩立",
		int:    500,
		level: level{
			"练气七层",
			7800.0,
		},
	}
	fmt.Println("======修仙者资料卡======")
	fmt.Println("姓名：", im.string)
	fmt.Println("年龄：", im.int)
	fmt.Println("------境界信息---------")
	fmt.Println("境界名称：", im.level.string)
	fmt.Println("境界灵气值：", im.level.float32)
	fmt.Println("------境界信息---------")
	fmt.Println("======修仙者资料卡======")
}
