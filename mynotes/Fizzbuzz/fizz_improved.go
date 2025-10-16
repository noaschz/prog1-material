package Fizzbuzz

import "fmt"

func FizzImproved(x, y, n int) {

	for i := 1; i <= n; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Println("fizzbuzz")

		} else if i%y == 0 {
			fmt.Println("buzz")
		} else if i%x == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}

}
