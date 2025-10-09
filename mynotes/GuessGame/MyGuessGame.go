package main

import (
	"fmt"
	"math/rand"
)

func main() {

	myNumber := rand.Intn(100) // Zufallszahl zwischen 0 und 99

	for i := 0; i < 3; i++ { // i++ ist i = i + 1
		guess := ReadNumber()

		if guess < myNumber {
			fmt.Println("Die gesuchte Zahl ist größer")
		} else if guess > myNumber {
			fmt.Println("Die gesuchte Zahl ist kleiner")
		} else {
			fmt.Println("Das war richtig!")
			return
		}
	}
	fmt.Printf("Game over! Die richtige Zahl war %d\n", myNumber)
}

func ReadNumber() int {
	var n int
	fmt.Print("Bitte gib eine Zahl ein: ")
	fmt.Scan(&n)
	return n
}
