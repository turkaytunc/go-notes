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
