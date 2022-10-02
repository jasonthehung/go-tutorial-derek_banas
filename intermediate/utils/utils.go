package utils

import "strconv"

var Name string = "Jason"

func IntArrayToStringArr(intArr []int) []string {
	var strArr []string

	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}

	return strArr
}
