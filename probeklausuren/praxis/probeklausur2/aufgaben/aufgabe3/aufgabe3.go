package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.
func CountSquares(list []int) int {
	counter := 0
	if len(list) == 0 {
		return counter
	}
	zahl := list[0]
	wurzel := int(math.Sqrt(float64(zahl)))
	if wurzel*wurzel == zahl {
		counter++
	}
	return counter + CountSquares(list[1:])

	/* counter := 0
	if len(list) == 0 {
		return counter
	}
	zahl := list[0]
	teiler := 0
	for i := 1; i <= zahl; i++ {
		if zahl%i == 0 {
			teiler++
		}
	}
	if teiler == 2 {
		counter++
		teiler = 0
	}
	return counter + CountSquares(list[1:]) */
}
