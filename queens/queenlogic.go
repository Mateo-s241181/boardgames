package queens

import "boardgames/board"

// QueenAllowed erwartet ein Spielfeld und eine Position.
// Liefert true, falls es erlaubt ist, dort eine Dame zu setzen.
func QueenAllowed(b board.Board, row, col int) bool {

	//Moves einer Dame: Diagonal, Horizontal, Vertikal, Nur Damen sind auf dem Feld
	//Wenn die Horizontale, Vertikale und Diagonale leer sind, ist die Position erlaubt

	if b.RowEmpty(row) && b.ColumnEmpty(col) && b.DiagDownRightEmpty(col-row) && b.DiagUpRightEmpty(col-(len(b)-1-row)) {
		return true
	}

	//Alle Restlichen sind nicht verboten

	return false
}
