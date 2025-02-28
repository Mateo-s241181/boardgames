package main

import (
	"boardgames/board"
	"boardgames/queens"
)

func main() {
	for n := 1; n <= 10; n++ {
		b := board.EmptyBoard(n)
		SolveQueens(b, 0)
		board.PrintBoard(b)
	}
}

// SolveQueens erwartet ein Spielfeld und eine Zeilennummer row.
// Versucht, das Spiel ab Zeile row zu lösen.
// Liefert true, falls dies gelingt und setzt die entsprechenden Damen.
func SolveQueens(b board.Board, row int) bool {

	//Durch row rangen und nach einer Erlaubten Position suchen
	for i := range b[row] {
		//An der ersten erlaubten Position eine Königin setzen und die Funktion eine Rekursionsebene tiefer callen
		if queens.QueenAllowed(b, row, i) {
			b[row][i] = "Q"

			if row == len(b)-1 {
				return true
			}

			//Falls das Spiel für diese Position von "Q" nicht lösbar ist, soll das Feld wieder leer werden
			if !SolveQueens(b, row+1) {
				b[row][i] = " "
			} else {
				return true
			}
		}
	}

	return false
}
