package board

// RowContainsOnly erwartet eine Zeilennummer row und einen String s.
// Liefert true, wenn die row-te Zeile nur aus dem gegebenen String s besteht.
func (board Board) RowContainsOnly(row int, s string) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetRow.
	 */
	// TODO
	return false
}

// ColumnContainsOnly erwartet eine Spaltennummer col und einen String s.
// Liefert true, wenn die col-te Spalte nur aus dem gegebenen String s besteht.
func (board Board) ColumnContainsOnly(col int, s string) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetColumn.
	 */
	// TODO
	return false
}

// DiagDownRightContainsOnly erwartet einen String s.
// Liefert true, wenn die Diagonale, die in Spalte 0 und Zeile 0 beginnt,
// nur aus dem gegebenen String s besteht.
func (board Board) DiagDownRightContainsOnly(s string) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetDiagDownRight.
	 */
	// TODO
	return false
}

// DiagUpRightContainsOnly erwartet einen String s.
// Liefert true, wenn die Diagonale, die in Spalte 0 und der letzten Zeile beginnt,
// nur aus dem gegebenen String s besteht.
func (board Board) DiagUpRightContainsOnly(s string) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetDiagUpRight.
	 */
	// TODO
	return false
}

// RowEmpty erwartet ein Spielfeld und eine Zeilennummer row.
// Liefert true, wenn die row-te Zeile leer ist, sonst false.
// Liefert true, wenn die Zeile nicht existiert.
func (board Board) RowEmpty(row int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktion RowContainsOnly.
	 */
	// TODO
	return false
}

// ColumnEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, wenn die col-te Spalte leer ist, sonst false.
// Liefert true, wenn die Spalte nicht existiert.
func (board Board) ColumnEmpty(col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktion ColumnContainsOnly.
	 */
	// TODO
	return false
}

// DiagDownRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und Zeile 0 beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func (board Board) DiagDownRightEmpty(col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen GetDiagDownRight und ContainsOnly.
	 */
	// TODO
	return false
}

// DiagUpRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und der letzten Zeile beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func (board Board) DiagUpRightEmpty(col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen GetDiagUpRight und ContainsOnly.
	 */
	// TODO
	return false
}
