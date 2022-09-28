package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	for x := 1; x <= 5; x++ {
		pl(x)
	}

	seedSec := time.Now().Unix()
	rand.Seed(seedSec)
	randNum := rand.Intn(50) + 1

	for {
		fmt.Print("Guess a number between 0 and 50 :\n")
		pl("The answer is :", randNum)

		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		} else {
			guess = strings.TrimSpace(guess)
			iGuess, err := strconv.Atoi(guess)

			if err != nil {
				log.Fatal(err)
			}

			if iGuess > randNum {
				pl("Pick a lower value")
			} else if iGuess < randNum {
				pl("Pick a higher value")
			} else {
				pl("You guessed it!")
				break
			}
		}
	}
}
