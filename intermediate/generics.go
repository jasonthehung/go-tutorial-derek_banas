package main

import "fmt"

var pl = fmt.Println

type MyConstraint interface {
	int | float64 | uint
}

func getSumGen[T MyConstraint](x T, y T) T {
	var sum T
	sum = x + y

	return sum
}

func main() {
	pl("uint + uint :", getSumGen(1, 2))
	pl("uint + int :", getSumGen(1, -2))
	pl("int + uint :", getSumGen(-1, 2))
	pl("uint + float64 :", getSumGen(1, float64(2.998)))
	pl("float64 + uint :", getSumGen(float64(2.998), 5))
}
