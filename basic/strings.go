package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {

	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Anotheroooo")         // replacer can do: Replace 'A' to 'Anotheroooo'
	sV2 := replacer.Replace(sV1)                                // use replacer to modify sV1 and replace string to another string
	pl("sv2 :", sV2)                                            // Anotheroooo word
	pl("Length :", len(sV2))                                    // 12
	pl("Contains Another :", strings.Contains(sV2, "Anothero")) // true
	pl("o index :", strings.Index(sV2, "o"))
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words\n"
	pl(sV3)
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)

	pl("Split :", strings.Split("a-b-c-d", "-"))
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))
}
