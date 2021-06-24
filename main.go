package main

import (
	"fmt"

	"github.com/turkaytunc/go-notes/inter"
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

	var area [2]inter.AreaCalculator
	area[0] = &inter.Square{Height: 5}
	area[1] = &inter.Circle{R: 3}

	for i := range area {
		fmt.Println(area[i].CalculateArea())
	}

	converted, ok := area[1].(*inter.Circle)
	if ok {
		fmt.Println(converted.CalculatePerimeter())
	}
}
