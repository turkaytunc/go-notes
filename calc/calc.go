package calc

type Number int

func (n Number) Add(otherNumber int) int {
	return int(n) + otherNumber
}

func (n Number) Subtract(otherNumber int) int {
	return int(n) - otherNumber
}
