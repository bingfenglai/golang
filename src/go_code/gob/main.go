package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// 修仙者
type Immortal struct {
	Name   string
	Age    int
	Gender string
}

type SimpleImmortal struct {
	Name   string
	Age    int
}

var buf  bytes.Buffer

func main() {
	
	var hanli = Immortal{
		Name:   "韩立",
		Age:    18,
		Gender: "男性",
	}

	fmt.Println("发送数据: ",hanli)
	sendMsg(&hanli)
	fmt.Println("buf中的数据：",buf)
	var i SimpleImmortal
	msg, _ := receiveMsg(i)

	fmt.Println("接收到数据：",msg)
}

func sendMsg(immortal *Immortal) error {
	enc :=gob.NewEncoder(&buf)
	return enc.Encode(immortal)
}

func receiveMsg(immortal SimpleImmortal) (SimpleImmortal,error) {
	dec := gob.NewDecoder(&buf)
	return immortal,dec.Decode(&immortal)
}
