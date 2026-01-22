/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2026-01-13
* @fileoverview This program makes a random number 1-10 and prompts the user to guess it until they find it.
*/

// constants
const MAX = 10;
const MIN = 1;

// variable
let ranNumber: number;
let guessString: string;
let guess: number;

// make number
ranNumber = Math.floor(Math.random() * (MAX - MIN + 1)) + MIN;

console.log("Guess a number between 1 and 10. Enter 0 to quit.")

// loop
do {
  guessString = prompt("Enter your guess: ") || "0";
  guess = parseInt(guessString);
  if (guess === ranNumber) {
    console.log(`Congratulations! You guessed the correct number: ${ranNumber}`);
    break;
  } else {
    console.log("Try again.");
  }
} while (ranNumber !== 0);

console.log("\nDone.")
