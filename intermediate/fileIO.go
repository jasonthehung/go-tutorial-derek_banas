package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	} else {
		// ...
	}

	defer f.Close()

	iPrimeArr := []int{11, 13, 17}
	var sPrimeArr = []string{}

	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		pl("Prime :", scan.Text())
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	defer f.Close()

}
