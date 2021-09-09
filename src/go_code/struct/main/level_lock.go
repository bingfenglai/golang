package main

import "go_code/method/model"

func main() {
	level := model.NewLevelLock("练气九层",9200)
	// 获取锁
	level.Lock.Lock()
	//修改值
	level.SetLevel("练气圆满")
	// 释放锁
	defer level.Lock.Unlock()
}
