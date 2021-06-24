package main

import (
	"fmt"
	"time"
)

func main() {

	// arr := readfromfile.New("data.txt")
	// for _, val := range arr {
	// 	fmt.Println(val)
	// }
	// arr1 := make([]int, 10)
	// for i := range arr1 {
	// 	arr1 = append(arr1, i+1)
	// }
	// sum := varsum.Sum(arr1...)
	// fmt.Println(sum)

	// var area [2]inter.AreaCalculator
	// area[0] = &inter.Square{Height: 5}
	// area[1] = &inter.Circle{R: 3}

	// for i := range area {
	// 	fmt.Println(area[i].CalculateArea())
	// }

	// converted, ok := area[1].(*inter.Circle)
	// if ok {
	// 	fmt.Println(converted.CalculatePerimeter())
	// }

	// res, err := crawl.New()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(res)

	go printA()
	go printB()

	time.Sleep(time.Second / 10)
}

func printA() {
	for i := 0; i < 50; i++ {
		go fmt.Print("a")
	}
}

func printB() {
	for i := 0; i < 50; i++ {
		go fmt.Print("b")
	}
}
