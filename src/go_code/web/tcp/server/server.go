package main

import (
	"fmt"
	"go_code/web/tcp/constants"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen(constants.Protocol, constants.Addr)

	if err != nil {
		println("监听端口出错", err)
		return

	} else {
		fmt.Printf("服务器启动完成 " + time.Now().Format("2006-01-02 15:04:05"))
	}

	for {

		accept, err := listen.Accept()
		if err != nil {
			println(err)
			return

		}

		go doServerStuff(accept)

	}

}

func doServerStuff(conn net.Conn) {

	for {
		buf := make([]byte, 1024)

		read, err := conn.Read(buf)

		if err != nil {
			println(err)
			return

		}

		println("\n 接收到数据：\n", string(buf[:read]))

		conn.Write([]byte(" hello client " + conn.RemoteAddr().String()))
	}

}
