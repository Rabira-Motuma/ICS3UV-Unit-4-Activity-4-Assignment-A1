/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2026-01-13
 * Fileoverview: This program makes a random number 1-10 and prompts the user to guess it until they find it.
 */

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main () {
	// constants
	const MIN = 1
	const MAX = 10

	// variables
	var ranNumber int
	var guessString string
	var guess int

	reader := bufio.NewReader(os.Stdin)

	// make number
	ranNumber = rand.IntN(MAX-MIN+1) + MIN

	// loop
	fmt.Println("Welcome to the game!")
	fmt.Println("Guess a number between 1 and 10. Enter 0 to quit.")

	guess = 1
		for guess != 0 {
	fmt.Printf("Enter your guess: ") 
	guessString, _ = reader.ReadString('\n')
	guessString = strings.TrimSpace(guessString)
	guess, _ = strconv.Atoi(guessString)
			if guess == ranNumber {
					fmt.Printf("Congratulations! You guessed the correct number: %d\n", ranNumber)
					break
				} else {
					fmt.Println("Try again.")
				}
	}

	fmt.Println("\nDone.")
}
