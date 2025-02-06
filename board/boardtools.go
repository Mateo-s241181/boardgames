package board

// RowEmpty erwartet ein Spielfeld und eine Zeilennummer row.
// Liefert true, wenn die row-te Zeile leer ist, sonst false.
// Liefert true, wenn die Zeile nicht existiert.
func RowEmpty(board Board, row int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetRow.
	 */
	return ContainsOnly(GetRow(board, row), " ")
}

// ColumnEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, wenn die col-te Spalte leer ist, sonst false.
// Liefert true, wenn die Spalte nicht existiert.
func ColumnEmpty(board Board, col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetColumn.
	 */
	return ContainsOnly(GetColumn(board, col), " ")
}

// DiagDownRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und Zeile 0 beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagDownRightEmpty(board Board, col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetDiagDownRight.
	 */
	return ContainsOnly(GetDiagDownRight(board, col), " ")
}

// DiagUpRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und der letzten Zeile beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagUpRightEmpty(board Board, col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen ContainsOnly und GetDiagUpRight.
	 */
	return ContainsOnly(GetDiagUpRight(board, col), " ")
}
