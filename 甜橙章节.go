package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//甜橙
func main() {
	fi, err := os.Open("甜橙/爱后余生：霍先生，请温柔.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
}
