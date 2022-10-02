package main

import "fmt"

var pl = fmt.Println

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func main() {
	iSlice := []float64{11, 13, 17}
	pl(getAverage(iSlice...))
}
