package main

import (
	"fmt"

	readfromfile "github.com/turkaytunc/go-notes/read-from-file"
	varsum "github.com/turkaytunc/go-notes/var-sum"
)

func main() {

	arr := readfromfile.New("data.txt")
	for _, val := range arr {
		fmt.Println(val)
	}
	arr1 := make([]int, 10)
	for i := range arr1 {
		arr1 = append(arr1, i+1)
	}
	sum := varsum.Sum(arr1...)
	fmt.Println(sum)
}
