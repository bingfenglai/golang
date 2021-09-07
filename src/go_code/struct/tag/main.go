package main

import "go_code/struct/model/tag"

func main() {
	var immortal = tag.Immortal{Name: "南宫婉", Age: 18, Gender: "女"}
	for i := 0; i < 3; i++ {
		tag.PrintTag(immortal, i)
	}
}
