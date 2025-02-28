package knights

import "boardgames/board"

// BoardPos repräsentiert eine Position auf dem Spielfeld.
type BoardPos struct {
	Row int
	Col int
}

// KnightAllowed erwartet ein Spielfeld und eine Position.
// Liefert true, falls auf dieses Feld ein Springer gesetzt werden darf.
func KnightAllowed(b board.Board, pos BoardPos) bool {

	accessibleNeighbors := []BoardPos{}

	//Filter neighbors => Positionen müssen im Board enthalten sein
	for i := range KnightNeighbours(pos) {

		if KnightNeighbours(pos)[i].Col >= 0 && KnightNeighbours(pos)[i].Col < len(b[0]) && KnightNeighbours(pos)[i].Row >= 0 && KnightNeighbours(pos)[i].Row < len(b) {
			accessibleNeighbors = append(accessibleNeighbors, KnightNeighbours(pos)[i])
		}
	}

	for _, el := range accessibleNeighbors {

		if b[el.Row][el.Col] != " " {
			return false
		}
	}

	return true
}

// KnightNeighbours erwartet eine Spielfeldposition und liefert eine Liste mit
// allen von dort mittels einer Springer-Bewegung erreichbaren Positionen,
func KnightNeighbours(pos BoardPos) []BoardPos {

	UpTwoLeft := BoardPos{pos.Row - 1, pos.Col - 2}
	TwoUpLeft := BoardPos{pos.Row - 2, pos.Col - 1}
	DownTwoLeft := BoardPos{pos.Row + 1, pos.Col - 2}
	TwoDownLeft := BoardPos{pos.Row + 2, pos.Col - 1}

	UpTwoRight := BoardPos{pos.Row - 1, pos.Col + 2}
	TwoUpRight := BoardPos{pos.Row - 2, pos.Col + 1}
	DownTwoRight := BoardPos{pos.Row + 1, pos.Col + 2}
	TwoDownRight := BoardPos{pos.Row + 2, pos.Col + 1}

	neighbors := []BoardPos{UpTwoLeft, UpTwoRight, DownTwoLeft, DownTwoRight, TwoUpLeft, TwoUpRight, TwoDownLeft, TwoDownRight}

	return neighbors
}
