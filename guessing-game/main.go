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

func main() {
	fmt.Println("*** Welcome to the Guessing Game! ***")
	fmt.Println("I'm thinking of a number between 0-100. Guess what it is!")
	gameLoop()
}

func gameLoop() {
	num := randomNum()
	for i := 0; i < 10; i++ {
		fmt.Printf("You have %v guesses remaining.\n", 10-i)
		guess := getGuessInput()

		if guess < num {
			fmt.Println("Your guess was too low!")
		} else if guess > num {
			fmt.Println("Your guess was too high!")
		} else if guess == num {
			fmt.Println("You guessed it! Good job!")
			return
		}
	}

	fmt.Println("Oh no, you've run out of guesses. Game over!")
}

// Return random number between 0-100
func randomNum() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}

func getGuessInput() int {
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdout)
	guessInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	guessTrimmed := strings.TrimSpace(guessInput)
	guess, err := strconv.Atoi(guessTrimmed)
	if err != nil {
		log.Fatal(err)
	}
	return guess
}
