package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	seedSec := time.Now().Unix()
	rand.Seed(seedSec)
	randNum := rand.Intn(1) + 1
	pl("Random :", randNum)

	// %t = boolean
	// %s = string
	// %o = Base 8
	// %x = Base 16
	// %v = (Guesses based in data type)
	// %T = Type of supplied value
}
