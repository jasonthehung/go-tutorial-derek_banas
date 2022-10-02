package main

import "fmt"

var pl = fmt.Println

func main() {
	// Method 1
	var heroes map[string]string
	heroes = make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Jason Wang"

	// Method 2
	villians := make(map[string]string)
	villians["Lex Luther"] = "Lex Luther KKT"

	// Method 3
	superPets := map[int]string{1: "Crypto", 2: "BTC", 3: "ETH"}
	pl("superPets 1 :", superPets[1])

	// Go through Maps
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// delete by key
	delete(heroes, "Batman")

	// Go through Maps
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
}
