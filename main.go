package main

import (
	"fmt"
	"math/rand"
	"strings"
)


func main() {
	fmt.Print("\nWelcome to the Number Guessing Game!\n")
	fmt.Print("I'm thinking of a number between 1 and 100.\n\n")

	for true {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Print("3. Hard (3 chances)\n\n")

	var choice int
	highscore := 100
	fmt.Print("Enter your choice: ")
	fmt.Scanf("%d", &choice)

	if choice < 1 || choice > 3 {
		fmt.Println("Invalid Choice ")
		return
		
	}
	tries := 0

	switch choice {
	case 1:
		tries = 10
		fmt.Println("Great! You have selected the Easy difficulty level. You have 10 chances to guess the number.")

	case 2:
		tries = 5
		fmt.Println("Great! You have selected the Medium difficulty level. You have 5 chances to guess the number.")

	case 3:
		tries = 3
		fmt.Println("Great! You have selected the Hard difficulty level. You have 3 chances to guess the number.")

	}

	number := rand.Intn(100) + 1
	var guess int

	turn := 0

	guessed := false
	
	for tries > 0 {
		fmt.Printf("\nYou have %d tries left. Enter your guess: ", tries)
		fmt.Scanf("%d", &guess)

		if guess < 1 || guess > 100 {
			fmt.Println("Please enter a number between 1 and 100.")
			continue
		}

		turn += 1
		
		if guess < number {
			fmt.Println("Too low!")
		} else if guess > number {
			fmt.Println("Too high!")
		} else {
			guessed = true
			break
		}
		tries--
	}
	if guessed {
		fmt.Printf("\nCongratulations! You've guessed the number %d correctly in %d tries!\n", number, turn)
		if turn < highscore {
		highscore = turn
		fmt.Printf("New high score! You've guessed the number in %d tries.\n", highscore)
	} else {
		fmt.Printf("Your high score remains at %d tries.\n", highscore)
	}
	} else {
		fmt.Printf("\nSorry, you've used all your tries. The number was %d. Better luck next time!\n", number)
	}

	


	fmt.Print("\nDo you want to play again? (yes/no): ")
	var playAgain string
	fmt.Scanf("%s", &playAgain)

	if strings.ToLower(playAgain) != "yes" && strings.ToLower(playAgain) != "no" {
		fmt.Println("Invalid input. Exiting the game.")
		return
	}

	if strings.ToLower(playAgain) != "yes" {
		fmt.Println("Thank you for playing! Goodbye!")
		break
	}
	}
}
