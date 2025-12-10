package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}
	// TODO
	if len(l1) == 0 || len(l2) == 0 {
		return result
	}
	last1 := len(l1) - 1
	last2 := len(l2) - 1
	maxlength := len(l1)
	if len(l2) > maxlength {
		maxlength = len(l2)
	}
	for i := 0; i < maxlength; i++ {
		if i < len(l1) {
			last1 = l1[i]
		}
		if i < len(l2) {
			last2 = l2[i]
		}
		result = append(result, last1+last2)
	}
	return result
}
