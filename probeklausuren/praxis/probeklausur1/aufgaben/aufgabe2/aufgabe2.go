package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	// TODO
	firstpostion := -1
	lastposition := -1

	for i := 0; i < len(list); i++ {
		if list[i] == first {
			firstpostion = i
		}
		if list[i] == last {
			lastposition = i
		}
	}

	if firstpostion == -1 || lastposition == -1 || lastposition <= firstpostion {
		return []string{}
	}

	restlist := append(list[:firstpostion], list[lastposition+1:]...)
	return restlist
}
