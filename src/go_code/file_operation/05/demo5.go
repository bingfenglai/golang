package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	buf := make([]byte, 1024)

	inputFile, err := os.Open("G:\\06_golangProject\\golang\\doc\\筑基部分\\10_golang中的读写数据.md")
	defer inputFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, err := inputFile.Read(buf)

		if err == io.EOF {
			return
		}
		fmt.Println(string(buf))
	}

}
