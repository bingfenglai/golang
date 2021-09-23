package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("G:\\06_golangProject\\golang\\doc\\筑基部分\\10_golang中的读写数据.md")
	defer inputFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	input := bufio.NewReader(inputFile)
	for {

		readString, err := input.ReadString('\n')
		fmt.Println(readString)
		if err == io.EOF {
			fmt.Println(err)
			return
		}
	}

}
