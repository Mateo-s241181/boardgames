package board

// ContainsOnly erwartet eine Liste von Strings und liefert true,
// wenn all diese Strings gleich dem gegebenen String s sind, sonst false.
func ContainsOnly(list []string, s string) bool {
	/* Hinweis:
	 * Prüfen Sie mit einer Schleife, ob irgend einer der
	 * Listeneinträge ungleich s ist.
	 */
	for _, v := range list {
		if v != s {
			return false
		}
	}
	return true
}

// ContainsAny erwartet eine Liste von Strings und liefert true,
// wenn darin mindestens ein String gleich dem gegebenen String s ist, sonst false.
func ContainsAny(list []string, s string) bool {
	/* Hinweis:
	 * Prüfen Sie mit einer Schleife, ob irgend einer der
	 * Listeneinträge gleich s ist.
	 */
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
