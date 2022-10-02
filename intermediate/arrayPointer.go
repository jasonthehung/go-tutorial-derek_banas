package main

import "fmt"

var pl = fmt.Println

func doubleArrayVals(arr *[4]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

func main() {
	var arr = [4]int{1, 2, 3, 4}
	pl(arr)
	doubleArrayVals(&arr)
	pl(arr)
}
