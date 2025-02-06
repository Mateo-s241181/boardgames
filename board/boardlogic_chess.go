package board

// RowContainsQueen erwartet ein Spielfeld und eine Zeilennummer row.
// Liefert true, falls die row-te Zeile mindestens ein "Q" enthält,
// sonst false.
// Liefert false, falls die Zeile nicht existiert.
func RowContainsQueen(board Board, row int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsAny und GetRow.
	 */
	return ContainsAny(GetRow(board, row), "Q")
}

// ColumnContainsQueen erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die col-te Spalte mindestens ein "Q" enthält,
// sonst false.
// Liefert false, falls die Spalte nicht existiert.
func ColumnContainsQueen(board Board, col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsAny und GetColumn.
	 */
	return ContainsAny(GetColumn(board, col), "Q")
}

// DiagDownRightContainsQueen erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und Zeile 0 beginnt,
// mindestens ein "Q" enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagDownRightContainsQueen(board Board, col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsAny und GetDiagDownRight.
	 */
	return ContainsAny(GetDiagDownRight(board, col), "Q")
}

// DiagUpRightContainsQueen erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und der letzten Zeile beginnt,
// mindestens ein "Q" enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagUpRightContainsQueen(board Board, col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsAny und GetDiagUpRight.
	 */
	return ContainsAny(GetDiagUpRight(board, col), "Q")
}
