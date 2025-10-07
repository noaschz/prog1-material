package main //erste Zeile in jeder Go-Datei (package). Später werden wir nicht mehr main verwenden.

import "fmt" //Importieren des fmt-Pakets für Ein- und Ausgabe. Formatted I/O with functions.

func main() { //Hauptfunktion, der Einstiegspunkt des Programms
	fmt.Println("Hello world!") //Ausgabe auf der Konsole: "Hello world!"
}

//mit go run . wird die aktuelle Datei ausgeführt
