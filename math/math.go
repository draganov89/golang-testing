package math

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Compare(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumEven(nums ...int) int {
	sum := 0
	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}
