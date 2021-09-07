package tag

import (
	"fmt"
	"reflect"
)

type Immortal struct {
	Name   string "休仙者的名字"
	Age    int    "修仙者的年龄"
	Gender string "修仙者的性别"
}

func PrintTag(im Immortal, i int) {
	imm := reflect.TypeOf(im)
	value := imm.Field(i)
	fmt.Println(value.Tag)
}
