package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/h2non/filetype.v1"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("usage: %s <filename>", os.Args[0])
		os.Exit(1)
	}

	arg := os.Args[len(os.Args)-1]

	buf, _ := ioutil.ReadFile(arg)

	kind, unknown := filetype.Match(buf)
	if unknown != nil {
		fmt.Printf("Unknown: %s", unknown)
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}
