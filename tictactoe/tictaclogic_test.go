package tictactoe

import (
	"boardgames/board"
	"fmt"
)

func ExamplePlayerWins() {
	bx_row := board.Board{
		[]string{"X", "X", "X"},
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
	}

	bo_diag := board.Board{
		[]string{"O", " ", " "},
		[]string{" ", "O", " "},
		[]string{" ", " ", "O"},
	}

	b_none := board.Board{
		[]string{"X", "O", "X"},
		[]string{"O", "X", "O"},
		[]string{"O", "X", "O"},
	}

	fmt.Println(PlayerWins(bx_row, "X"))
	fmt.Println(PlayerWins(bx_row, "O"))
	fmt.Println(PlayerWins(bo_diag, "X"))
	fmt.Println(PlayerWins(bo_diag, "O"))
	fmt.Println(PlayerWins(b_none, "X"))
	fmt.Println(PlayerWins(b_none, "O"))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// false
}

func ExamplePlayerAllowed() {
	b := board.Board{
		[]string{"X", "O", "X"},
		[]string{" ", " ", "O"},
		[]string{"O", "X", " "},
	}

	fmt.Println(PlayerAllowed(b, 0, 0))
	fmt.Println(PlayerAllowed(b, 0, 1))
	fmt.Println(PlayerAllowed(b, 0, 2))
	fmt.Println()

	fmt.Println(PlayerAllowed(b, 1, 0))
	fmt.Println(PlayerAllowed(b, 1, 1))
	fmt.Println(PlayerAllowed(b, 1, 2))
	fmt.Println()

	fmt.Println(PlayerAllowed(b, 2, 0))
	fmt.Println(PlayerAllowed(b, 2, 1))
	fmt.Println(PlayerAllowed(b, 2, 2))
	fmt.Println()

	fmt.Println(PlayerAllowed(b, 3, 0))
	fmt.Println(PlayerAllowed(b, 0, -1))

	// Output:
	// false
	// false
	// false
	//
	// true
	// true
	// false
	//
	// false
	// false
	// true
	//
	// false
	// false
}
