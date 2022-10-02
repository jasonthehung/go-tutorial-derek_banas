package main

import "fmt"

var pl = fmt.Println

func getQuotent(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

func getSum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func getSumByArray(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}

	return sum
}

func main() {
	// pl(getQuotent(2.3, 0))
	pl(getSum(1, 2, 3, 4, 5, 6, 7, 8))

	varArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
	pl(getSumByArray(varArray))
}
