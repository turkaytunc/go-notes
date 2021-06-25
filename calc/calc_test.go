package calc

import "testing"

type number struct {
	A Number
	B Number
}
type testData struct {
	number
	want int
}

func TestAdd(t *testing.T) {
	testCases := []testData{
		{number: number{3, 5}, want: 8},
		{number: number{5, 5}, want: 10},
		{number: number{-2, 5}, want: 3},
	}

	for _, v := range testCases {
		got := v.A.Add(int(v.B))

		if got != v.want {
			t.Errorf("Add(%#v) = \"%d\", want \"%d\"", v.number, got, v.want)
		}
	}
}

func TestSubtact(t *testing.T) {
	testCases := []testData{
		{number: number{3, 5}, want: -2},
		{number: number{5, 5}, want: 0},
		{number: number{-2, 5}, want: -7},
	}

	for _, v := range testCases {
		got := v.A.Subtract(int(v.B))

		if got != v.want {
			t.Errorf("Add(%#v) = \"%d\", want \"%d\"", v.number, got, v.want)
		}
	}
}
