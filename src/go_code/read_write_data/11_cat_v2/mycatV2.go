package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	if filename == "" {
		fmt.Println("Command Usage Format: mycat filename")
		return
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	myCatV2(bufio.NewReader(file))
}

func myCatV2(reader *bufio.Reader) {
	buf := make([]byte, 512)

	for {
		n, err := reader.Read(buf)
		fmt.Fprintf(os.Stdout, "%s", buf[0:n])
		if err == io.EOF {
			break
		}
	}
	return
}
