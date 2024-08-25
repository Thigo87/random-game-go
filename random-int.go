package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	num := rand.Intn(101)
	choice := 0
	playAgain := ""
	for {
		fmt.Println("Write your guess: 0-100")
		fmt.Scanln(&choice)
		if choice > num {
			fmt.Println("Too high!")
		} else if choice < num {
			fmt.Println("Too low!")
		} else {
			fmt.Println("You guessed it!")
			fmt.Println("Do you want to play again? Y/N")
			fmt.Scanln(&playAgain)
			if strings.ToUpper(playAgain) == "Y" {
				num = rand.Intn(101)
			} else {
				break
			}
		}
	}
}
