package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Intn(101)
	choice := 0
	for {
		fmt.Println("Write your guess:")
		fmt.Scanln(&choice)
		if choice > num {
			fmt.Println("Too high!")
		} else if choice < num {
			fmt.Println("Too low!")
		} else {
			fmt.Println("You guessed it!")
			break
		}
	}
}
