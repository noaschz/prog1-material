package listprops

// ContainsNRow liefert true, falls die Liste l
// String x n mal hintereinander enthält.
func ContainsNRow(l []string, x string, n int) bool {
	// TODO
	var counter int
	for i := 0; i < len(l); i++ {
		for Folger := i; Folger < len(l); Folger++ {
			if l[Folger] == x {
				counter++
				if counter == n {
					return true
				}
			} else {
				counter = 0
				break
			}
		}
	}
	return false
}
