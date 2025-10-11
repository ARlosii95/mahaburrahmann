package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(5) + 1
	
	fmt.Print("Guess 1-5: ")
	var guess int
	fmt.Scan(&guess)
	
	if guess == number {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Wrong! It was %d\n", number)
	}
}