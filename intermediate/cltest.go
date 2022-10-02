package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var iArgs = []int{}

	for _, val := range args {
		val, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		iArgs = append(iArgs, val)
	}

	// Print MAX
	max := 0
	for _, val := range iArgs {
		if max < val {
			max = val
		} else {
			// continue
		}
	}
	fmt.Println("Max value :", max)
}
