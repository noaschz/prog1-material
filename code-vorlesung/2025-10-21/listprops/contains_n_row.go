package listprops

// ContainsNRow liefert true, falls die Liste l
<<<<<<< HEAD
// String x n mal hintereinander enthält.
=======
// den String x n mal hintereinander enthält.
>>>>>>> 9a22e09e105f7d8388ac0794ab0dcb95e29cca77
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
