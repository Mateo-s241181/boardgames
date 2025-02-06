package tictactoe

import "boardgames/board"

// PlayerWins erwartet ein Spielfeld und ein Zeichen,
// das einen Spieler repräsentiert (meist "X" oder "O").
// Liefert true, falls dieser Spieler gewonnen hat.
func PlayerWins(b board.Board, player string) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen Row
	 */
	// TODO
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
	// TODO
	return false
}
