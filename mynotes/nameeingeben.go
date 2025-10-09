package main

import "fmt"

func nameEingeben() string {
	var name string
	fmt.Print("Bitte geben Sie Ihren Namen ein: ")
	fmt.Scanln(&name)
	return name
}
