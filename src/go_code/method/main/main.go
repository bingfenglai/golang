package main

import (
	"fmt"
	"go_code/method/model"
)

func main() {
	i := model.NewImmortal(18,"韩立","男")
	name :=i.GetName()
	fmt.Println(name)
	

}
