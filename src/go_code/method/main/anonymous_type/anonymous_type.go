package main

import (
	"fmt"
	"go_code/method/model"
)

func main() {

	im := model.NewImmortal2(18,"韩立","男",
		"练气九层",9200,"木灵根","水灵根","火灵根","土灵根")
	im.SetLevel("练气大圆满")
	fmt.Println("境界：",im.LevelName())
	fmt.Println("灵根：",im.LingGenNames())
}
