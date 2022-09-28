package main

import "fmt"

var pl = fmt.Println

func main() {
	sl1 := make([]string, 6)

	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	pl("Slice size :", len(sl1))

	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}
}
