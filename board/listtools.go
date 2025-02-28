package board

// ContainsOnly erwartet eine Liste von Strings und liefert true,
// wenn all diese Strings gleich dem gegebenen String s sind, sonst false.
func ContainsOnly(list []string, s string) bool {

	for i := range list {

		if list[i] != s {
			return false
		}
	}
	return true
}

// ContainsAny erwartet eine Liste von Strings und liefert true,
// wenn darin mindestens ein String gleich dem gegebenen String s ist, sonst false.
func ContainsAny(list []string, s string) bool {

	for i := range list {

		if list[i] == s {
			return true
		}
	}

	return false
}
