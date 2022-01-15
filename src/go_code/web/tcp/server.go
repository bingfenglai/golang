package main

import "net"

func main() {

	listen, err := net.Listen("tcp", "localhost:9527")

	if err != nil {
		println("监听端口出错", err)
		return

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

		println("接收到数据：", string(buf[:read]))
	}

}
