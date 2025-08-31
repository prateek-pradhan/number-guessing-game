# Number Guessing Game (Go)

This is a simple **Number Guessing Game** written in Go.\
The game generates a random number between 1 and 100, and the player has
to guess it within a limited number of tries based on the selected
difficulty level.

## Features

-   Choose difficulty level:
    -   Easy: 10 chances
    -   Medium: 5 chances
    -   Hard: 3 chances
-   Provides hints whether your guess is too high or too low.
-   Keeps track of the number of tries taken.
-   Maintains a high score (fewest guesses).
-   Option to play again after each game.

## How to Run

1.  Make sure you have Go installed on your system.
2.  Save the provided code into a file named `main.go`.
3.  Open your terminal in the directory containing `main.go`.
4.  Run the following commands:

``` bash
go run main.go
```

## Gameplay Example

    Welcome to the Number Guessing Game!
    I'm thinking of a number between 1 and 100.

    Please select the difficulty level:
    1. Easy (10 chances)
    2. Medium (5 chances)
    3. Hard (3 chances)

    Enter your choice: 2
    Great! You have selected the Medium difficulty level. You have 5 chances to guess the number.

    You have 5 tries left. Enter your guess: 50
    Too low!

    You have 4 tries left. Enter your guess: 75
    Too high!

    You have 3 tries left. Enter your guess: 63
    Congratulations! You've guessed the number correctly!

## 🔗 Reference
This project is inspired by the [Number Guessing Game](https://roadmap.sh/projects/number-guessing-game).