package varsum

func Sum(nums ...int) (result int) {

	var sum int
	for _, val := range nums {
		sum += val
	}

	return sum
}
