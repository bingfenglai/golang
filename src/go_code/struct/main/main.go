package main

import (
	"go_code/struct/model"
	"strconv"
)

func main() {

	var i *model.Immortal
	i = new(model.Immortal)
	i.Age = 500
	i.Name = "韩立"
	i.Gender = "男"

	println(i.Name, strconv.Itoa(i.Age)+"岁", i.Gender)
	i2 := getImmortal()
	println(i2.Name, strconv.Itoa(i2.Age)+"岁", i2.Gender)

}

func getImmortal() *model.Immortal {

	var i = &model.Immortal{"南宫婉", 18, "女"}

	return i
}
