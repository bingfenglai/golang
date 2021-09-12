package main

import (
	"fmt"
	"go_code/interface/interfaces"
	"go_code/interface/model"
)

func main() {

	 // 声明了一个SpiritualRootAble接口类型的变量
	 var sr interfaces.SpiritualRootAble

	 // 降生了一个凡人
	 mortal := model.NewMortal("韩立","男性",1)

	 // 接口变量指向凡人实例
	sr = mortal

	// 获取凡人的灵根
	fmt.Println(sr.SpiritualRoot())

	// 凡人开始修炼
	sr.Practice()
	
	

	if v,ok :=sr.(*model.Mortal);ok{
		fmt.Println(ok)
		fmt.Println(v)
	}
	var m interface{}
	if _,ok := m.(interfaces.SpiritualRootAble);ok {
		fmt.Println(ok)
	}else {
		fmt.Println(ok)
	}

}
