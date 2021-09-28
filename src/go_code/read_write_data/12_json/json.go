package main

import (
	"encoding/json"
	"fmt"
)

// 修仙者
type Immortal struct {
	Name   string
	Age    int
	Gender string
}

func main() {

	immortal := &Immortal{
		Name:   "韩立",
		Age:    18,
		Gender: "男性",
	}

	jsonImmortal, _ := json.Marshal(immortal)

	fmt.Printf("%s\n", jsonImmortal)
}
