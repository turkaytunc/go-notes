package main

import (
	"fmt"
	"strings"
)

func main() {

	brokenStr := "Hell# W#rld!"
	replacer := strings.NewReplacer("#", "o")
	fixedString := replacer.Replace(brokenStr)
	fmt.Println(fixedString)
}
