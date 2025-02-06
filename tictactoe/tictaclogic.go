package tictactoe

import "boardgames/board"

// PlayerWins erwartet ein Spielfeld und ein Zeichen,
// das einen Spieler repräsentiert (meist "X" oder "O").
// Liefert true, falls dieser Spieler gewonnen hat.
func PlayerWins(b board.Board, player string) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen Row
	 */
	for i := 0; i < len(b); i++ {
		if b.RowContainsOnly(i, player) {
			return true
		}
		if b.ColumnContainsOnly(i, player) {
			return true
		}
		if b.DiagDownRightContainsOnly(player) {
			return true
		}
		if b.DiagUpRightContainsOnly(player) {
			return true
		}
	}
	return false
}

// PlayerAllowed erwartet ein Spielfeld und eine Position
// in Form einer Zeilen- und einer Spaltennummer.
// Liefert true, falls auf dieses Feld ein Zug erlaubt ist.
// Ein Zug ist erlaubt, wenn das Feld leer ist
// und wenn die Position gültig ist.
func PlayerAllowed(b board.Board, row, col int) bool {
	/* Hinweis:
	 * Prüfen Sie ob die Zeilen- und Spaltennummer innerhalb des Spielfelds liegen.
	 * Prüfen Sie ob das Feld leer ist.
	 */
	return row >= 0 && row < len(b) && col >= 0 && col < len(b[0]) && b[row][col] == " "
}
