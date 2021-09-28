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

	// 1. 事先知道json对应的数据类型时
	var jsonValue Immortal

	json.Unmarshal(jsonImmortal, &jsonValue)

	fmt.Println("name", jsonValue.Name)
	fmt.Println("age", jsonValue.Age)
	fmt.Println("gender", jsonValue.Gender)

	// 2. 不知道json对应的数据结构
	var m interface{}
	json.Unmarshal(jsonImmortal, &m)

	jsonMap := m.(map[string]interface{})
	for key, value := range jsonMap {
		printJson(key, value)
	}

}

func printJson(key string, value interface{}) {

	switch value.(type) {
	case string:
		fmt.Println(key, "value is a string: ", value)
	case float64:
		fmt.Println(key, "value is int type: ", value)
	case []interface{}:
		fmt.Println(key, "value is a array", value)
	case map[string]interface{}:
		m := value.(map[string]interface{})
		for k, v := range m {
			printJson(k, v)
		}

	}

}
