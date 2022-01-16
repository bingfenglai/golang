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
		fmt.Println("服务器启动完成 " + time.Now().Format("2006-01-02 15:04:05"))
	}

	for {

		accept, err := listen.Accept()
		if err != nil {
			println("出现错误:\n", err.Error())
			return

		}

		go doServerStuff(accept)

	}

}

func doServerStuff(conn net.Conn) {

	defer func() {
		println("关闭连接")
		conn.Close()
	}()

	for {
		buf := make([]byte, 256)

		read, err := conn.Read(buf)

		if err != nil {
			if "EOF" == err.Error() {
				println("数据读取结束")
				return
			}
			println("出现错误:", err.Error())
			return

		}

		println("接收到数据：")
		println(string(buf[:read]))
		_, err = conn.Write([]byte(" hello client " + conn.RemoteAddr().String()))

		if err != nil {
			println("出现错误\n", err.Error())
		}

	}

}
