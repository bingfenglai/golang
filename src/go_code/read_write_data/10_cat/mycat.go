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

	open, err := os.Open(filename)
	defer open.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	myCat(bufio.NewReader(open))
}

func myCat(reader *bufio.Reader) {

	for {
		buf, err := reader.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s\n", buf)
		if err == io.EOF {
			break
		}
	}
	return
}
