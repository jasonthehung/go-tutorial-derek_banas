package main

import "fmt"

var pl = fmt.Println

func changeVar(myPtr *int) {
	*myPtr = 20
}

func main() {
	myVar := 10

	pl("Before change :", myVar)
	changeVar(&myVar)
	pl("After change :", myVar)
}
