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

func generateRand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func main() {
	randomInt := generateRand()
	tries, limit := 0, 10
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess a num between 1 and 100..")

	for {
		if tries >= limit {
			fmt.Println("Sorry you didn't guess my number, it was ", randomInt)
			return
		}
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guessedNum, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guessedNum == randomInt {
			fmt.Println("Good job!, you guessed it")
			return
		}
		if guessedNum > randomInt {
			fmt.Println("Oops, your guess was too HIGH! ", limit-tries, "tries left.")
		} else {
			fmt.Println("Oops, your guess was too LOW! ", limit-tries, "tries left.")
		}
		tries++

	}
}
