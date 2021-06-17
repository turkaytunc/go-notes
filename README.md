# Headfirst Go Book Notes

## Index

### Type of

```go

reflect.TypeOf(42);
reflect.TypeOf(3.14159)
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
