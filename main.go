package main

import (
	"fmt"

	readfromfile "github.com/turkaytunc/go-notes/read-from-file"
)

func main() {

	arr := readfromfile.New("data.txt")

	for _, val := range arr {
		fmt.Println(val)
	}
}
