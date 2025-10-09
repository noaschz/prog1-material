package main

import (
	//"fmt"
	"fmt"
	"prog1-material/mynotes/greet"
)

func main() {
	fmt.Println(greet.Greet("Kurs"))
	fmt.Println("Alle " + greet.Greet("Kursteilnehmer: "))
	//greet.Greet("Kurs") // Aufruf der Greet-Funktion aus dem greet-Paket. Main und Greet sind im gleichen repo.
	//return fmt.Sprintf("Hello, %s!", name)
}
