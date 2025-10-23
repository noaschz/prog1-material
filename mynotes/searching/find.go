package searching

// Liefert die Position von x in der Liste l oder -1, falls x nicht vorkommt.
// Voraussetzung: l ist aufsteigend sortiert.
func BinFind1(l []int, x int) int {
	links := 0 // persistenter Offset nach rechts

	// Vergleiche x mit dem Element in der Mitte. Wenn x == l[mitte], dann fertig
	for len(l) > 0 {
		mitte := len(l) / 2

		if x == l[mitte] {
			return mitte + links
		}
		// Wenn x < l[mitte], dann suche in der linken Hälfte weiter
		// Wenn x > l[mitte], dann suche in der rechten Hälfte weiter
		if x < l[mitte] {
			// lasse nur den linken Teil der Liste übrig. Alles von 0 bis exklusive Mitte
			l = l[:mitte]
		} else {
			// lasse nur den rechten Teil der Liste übrig. Alles ab (Mitte + 1)
			l = l[mitte+1:]
			links += mitte + 1
		}
	}

	return -1
}

func BinFind2(l []int, x int) int {
	links := 0       //Hier beginnt der interessante Teil
	rechts := len(l) // Hier endet der interessante Teil
	for rechts > links {
		mitte := (rechts-links)/2 + links

		if x == l[mitte] {
			return mitte
		}

		if x < l[mitte] {
			//suche in der linken Hälfte weiter
			rechts = mitte
		} else {
			//suche in der rechten Hälfte weiter
			links = mitte + 1
		}

	}
	return -1
}
