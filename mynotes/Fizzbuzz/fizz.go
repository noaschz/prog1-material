package Fizzbuzz

import "fmt"

//Fizz gibt die Zahlen von 1 bis 30 aus und ersetzt jede durch 3 teilbare Zahl durch "Fizz", jede durch 5 teilbare Zahl durch "Buzz" und jede durch 3 und 5 teilbare Zahl durch "FizzBuzz".
func Fizz() {
	for i := 1; i <= 30; i++ {
		//wenn i durch 3 teilbar ist, gib "Fizz" aus
		//wenn i durch 5 teilbar ist, gib "Buzz" aus
		//wenn i durch 3 und 5 teilbar ist, gib "FizzBuzz" aus
		//ansonsten gib i aus
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
