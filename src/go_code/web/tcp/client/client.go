package main

import (
	"fmt"
	"go_code/web/tcp/constants"
	"net"
	"time"
)

func main() {

	// 建立连接

	conn, err := net.Dial(constants.Protocol, constants.Addr)

	if err != nil {
		println("建立连接失败", err)
		return

	} else {
		fmt.Println("建立连接成功")
	}
	conn.Write([]byte("hello server " + time.Now().Format("2006-01-02 15:04:05")))

	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)

		println("收到服务器发送来的消息")

		if err != nil {
			println("读取数据失败 ", err)
			return
		}

		fmt.Println(string(buf[:read]))
	}
}
