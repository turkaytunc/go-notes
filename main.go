package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter your number: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var num, _ = strconv.ParseInt(input, 10, 64)
	if num > 50 {
		fmt.Println("More than 50")
	}

	fmt.Println("your input is: " + input)
}
