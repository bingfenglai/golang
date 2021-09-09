package main

import (
	"fmt"
	"go_code/method/model"
)

func main() {
	i := model.NewImmortal(18,"韩立","男")
	name :=i.GetName()
	fmt.Println(name)

	level := model.NewLevel("练气九层",9200)
	levelPointer := & level
	fmt.Println("晋级之前：",level.Level())
	levelPointer.SetLevel("炼气大圆满")
	fmt.Println("晋级之后：",level.Level())
}
