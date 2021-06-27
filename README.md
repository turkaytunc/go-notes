# Headfirst Go Book Notes

## Index

### Type of

```go

reflect.TypeOf(42);
reflect.TypeOf(3.14159)

fmt.Printf("%v", reflect.TypeOf(5).Kind() == reflect.Int) // true
```

### Change between

```go
var first = "Hello"
var second = "World"

first, second = second, first
fmt.Println(first + " " + second)
```

### time package

```go
var currentTime = time.Now()

fmt.Println(currentTime.Year(), currentTime.Month(), currentTime.Day())
```

### Replace chars

```go
oldStr := "Hell# W#rld!"
replacer := strings.NewReplacer("#", "o")
newStr := replacer.Replace(oldStr)
fmt.Println(newStr)
```

### Get input from terminal

```go
fmt.Println("Enter your input: ")
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')

if err != nil {
  log.Fatal(err)
}

fmt.Println("your input is: " + input)
```

### String to number conversion

```go
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
```

### Guess the number game

```go
func main() {

	target := generateRandomNumber(1, 100)
	var endText string
	for i := 0; i < 10; i++ {
		inputNum := getInputFromUser()

		if inputNum > int64(target) {
			fmt.Println("Opps. Your guess was HIGH")
		} else if inputNum < int64(target) {
			fmt.Println("Opps. Your guess was LOW")
		} else {
			endText = "Good job! You guessed it!"
			fmt.Println(endText)
			return
		}
	}
	if endText == "" {
		fmt.Printf("Sorry. You didn't guess my number. It was: %v", target)
	}

}

func getInputFromUser() int64 {

	fmt.Println("Enter your number: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var inputNum, parseError = strconv.ParseInt(strings.Trim(input, "\n\r"), 10, 64)
	if parseError != nil {

		fmt.Println("error: ", parseError)
	}

	return inputNum
}

func generateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(max-min+1) + min
	return target
}

```

### Formatting Strings

```go

fmt.Printf("%s %s!", "Hello", "World")
fmt.Printf("%d", 5);

str := fmt.Sprintf("%s %s!", "Hello", "World") // Returns string

Verb 	Output

%f 		Floating-point number
%d 		Decimal integer
%s 		String
%t 		Boolean (true or false)
%v 		Any value (chooses an appropriate format based on the supplied value’s type)
%#v 	Any value, formatted as it would appear in Go program code
%T 		Type of the supplied value (int, string, etc.)
%% 		A literal percent sign
```

### Return Error and Named Return Values

```go
func squareRoot(number float64) (result float64, err error) {
	if number < 0 {
		return 0, fmt.Errorf("Can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}
```

### Pointers

```go
var myInt int
var myIntPointer *int
myInt = 42
myIntPointer = &myInt

fmt.Print(*myIntPointer)
```

### Go Doc

In terminal

```bash
go doc errors
go doc strconv
go doc strconv ParseInt
```

### Arrays

```go
primes := [5]int{2, 3, 5, 7, 11}
```

### Read Data From File

```go
file, err := os.Open("data.txt")
if err != nil {
	log.Fatal(err)
}

scanner := bufio.NewScanner(file)
for scanner.Scan() {
	fmt.Printf("%v\n", scanner.Text())
}
err = file.Close()
if err != nil {
	log.Fatal(err)
}
if scanner.Err() != nil {
	log.Fatal(scanner.Err())
}
```

### Slice an array

```go

primes := [5]int{2, 3, 5, 7, 11}

sliced := primes[1:3] // first parameter is inclusive, second is exclusive

fmt.Println(sliced) // [3, 5]

```

### Read Terminal Arguments

```go
arguments := os.Args[1:]
```

### Variadic Functions

```go
func myFunc(param1 string, param2 ...int){

}
```

### Passing Slice To Variadic Function

```go
func Sum(nums ...int) (result int) {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

arr1 := make([]int, 10)
for i, _ := range arr1 {
	arr1 = append(arr1, i+1)
}
sum := varsum.Sum(arr1...)
fmt.Println(sum)
```

### Creating Maps

```go
freq := make(map[string]int)

strArr := strings.Split(arr, "")

for v := range strArr {
	_, ok := freq[strArr[v]]
	if ok {
		freq[strArr[v]]++
		continue
	}
	freq[strArr[v]] = 1
}

freq["h"]
// values can be removed using delete function.   delete(mapName, value)  delete(freq, "h")

```

### Sorting

```go
arr2 := []string{"hello", "bottle", "car", "table"}

for _, v := range arr2 {
	fmt.Println(v)
}

sort.Strings(arr2)

for _, v := range arr2 {
	fmt.Println(v)
}
```

### Structs

```go
type user struct {
	name string
	age int
	isAdmin bool
}

// can be used like this
user1 := user{"mike", 15, false}
// or this
user2 := user{name: "John", age: 74, isAdmin: false}
```

### Receiver and Defined Types

```go
package calc

type Number int // Defined Type

func (n Number) Add(otherNumber int) int {
	return int(n) + otherNumber
}
func (n Number) Subtract(otherNumber int) int {
	return int(n) - otherNumber
}

//

num := calc.Number(5)
fmt.Println(num.Add(2)) // prints out 7

```

### interface

```go
package inter

import (
	"math"
)

type Double float64

type AreaCalculator interface {
	CalculateArea() Double
}

type Circle struct {
	R float64
}

func (c Circle) CalculateArea() Double {
	return Double(math.Pow(c.R, 2) * math.Pi)
}

type Square struct {
	Height float64
}

func (s Square) CalculateArea() Double {
	return Double(math.Pow(s.Height, 2))
}

// in main

var area [2]inter.AreaCalculator
area[0] = inter.Square{Height: 5}
area[1] = inter.Circle{R: 2}

for i := range area {
	fmt.Println(area[i].CalculateArea())
}

// if methods have pointer receivers then we must use like this

var area [2]inter.AreaCalculator
area[0] = &inter.Square{Height: 5}
area[1] = &inter.Circle{R: 2}

for i := range area {
	fmt.Println(area[i].CalculateArea())
}

// if we add another method that not satisfy interface
func (c *Circle) CalculatePerimeter() Double {
	return Double(c.R * 2 * math.Pi)
}
// then we have to convert "type assertion" it to use method above
converted, ok := area[1].(*inter.Circle) // return value and ok bool
if ok {
	fmt.Println(converted.CalculatePerimeter())
}
```

### Goroutines

```go
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

//

go printA()
go printB()

time.Sleep(time.Second / 10)
```

### channels

```go
func SiteBodyLength(url string, c chan string) {
	res, httpErr := http.Get(url)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
	defer res.Body.Close()

	body, parseErr := ioutil.ReadAll(res.Body)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	c <- string(body)
}

//

c := make(chan string)
go crawl.SiteBodyLength("https://google.com/", c)
go crawl.SiteBodyLength("https://amazon.com/", c)
go crawl.SiteBodyLength("https://aws.com/", c)
go crawl.SiteBodyLength("https://amazon.com/", c)
go crawl.SiteBodyLength("https://google.com/", c)

for v := range c {
	fmt.Println(len(v))
}
```

### Compose interface

```go

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type ReadCloser interface {
    Reader
    Closer
}
```
