package listprops

// ContainsN liefert true, falls die Liste l
// String x mindestens n mal enthält.
func ContainsN(l []string, x string, n int) bool {
	// TODO
	var wieoft int

	for i := 0; i < len(l); i++ {
		if l[i] == x {
			wieoft++
		}
	}
	if wieoft >= n {
		return true
	}
	return false
}
