package main

import (
	utils "example/intermediate/utils"
	"fmt"
	"reflect"
)

/*
This main func is related to utils folder, about package and modules
*/
func main() {
	fmt.Println("Hello")
	fmt.Println(utils.Name)
	intArr := []int{1, 2, 4}
	strArr := utils.IntArrayToStringArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}
