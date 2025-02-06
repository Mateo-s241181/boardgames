package main

import (
	"boardgames/board"
	"boardgames/knights"
)

func main() {
	b := board.EmptyBoard(8)
	SolveKnights(b, knights.BoardPos{Row: 0, Col: 0}, 1)
	board.PrintBoard(b)
}

// SolveKnights erwartet ein Spielfeld und eine Position, sowie eine laufende Nummer n.
// Versucht, das Spiel ab der Nummer n zu lösen, indem es n an die aktuelle Position
// schreibt und dann rekursiv mit allen für den Springer erreichbaren Feldern weitermacht.
func SolveKnights(b board.Board, pos knights.BoardPos, n int) bool {
	/* Hinweis:
	 * 1. Rekursionsanker: Bei negativem oder zu großem n ist das Spiel per Definitionem gelöst.
	 * 2. Rekursionsanker: Bei ungültiger oder bereits vergebener Position (verbotener Zug)
	 *    ist das Spiel nicht lösbar. Verwenden Sie die Funktion KnightAllowed.
	 */

	/* Hinweis:
	 * n auf das aktuelle Feld schreiben (verwenden Sie fmt.Sprintf).
	 */

	/* Hinweis:
	 * Verwenden Sie KnightNeighbours, um alle Nachbarpositionen aufzuzählen.
	 * Prüfen Sie in einer Schleife für jede dieser Positionen, ob das Spiel ab dort mit n+1 lösbar ist.
	 */

	/* Hinweis:
	 * Falls das Spiel durch die Schleife über die Nachbarpositionen nicht gelöst werden kann,
	 * ist es nicht lösbar. Entfernen Sie den aktuellen Schritt wieder aus dem Spielfeld.
	 */

	return false
}
