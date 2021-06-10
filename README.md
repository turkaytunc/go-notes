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
brokenStr := "Hell# W#rld!"
replacer := strings.NewReplacer("#", "o")
fixedString := replacer.Replace(brokenStr)
fmt.Println(fixedString)
```
