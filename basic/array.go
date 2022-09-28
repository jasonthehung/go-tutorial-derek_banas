package main

import "fmt"

func main() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)

	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

}
