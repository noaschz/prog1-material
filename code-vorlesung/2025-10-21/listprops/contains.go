package listprops

// Contains liefert true, falls die Liste l
// den String x enthält.
func Contains(l []string, x string) bool {
<<<<<<< HEAD
	// TODO
=======
>>>>>>> 9a22e09e105f7d8388ac0794ab0dcb95e29cca77
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return true
		}
	}
	return false
}
