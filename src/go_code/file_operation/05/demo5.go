package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	buf := make([]byte, 1024)

	inputFile, err := os.Open("古丹丹方.txt")
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
