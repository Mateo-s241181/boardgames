package board

import "fmt"

func ExampleBoard_GetRow() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(board.GetRow(0))
	fmt.Println(board.GetRow(1))
	fmt.Println(board.GetRow(2))

	fmt.Println(board.GetRow(-1))
	fmt.Println(board.GetRow(3))

	// Output:
	// [A B C]
	// [D E F]
	// [G H I]
	// []
	// []
}

func ExampleBoard_GetColumn() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(board.GetColumn(0))
	fmt.Println(board.GetColumn(1))
	fmt.Println(board.GetColumn(2))

	fmt.Println(board.GetColumn(-1))
	fmt.Println(board.GetColumn(3))

	// Output:
	// [A D G]
	// [B E H]
	// [C F I]
	// []
	// []
}

func ExampleBoard_GetDiagDownRight() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(board.GetDiagDownRight(-3))
	fmt.Println(board.GetDiagDownRight(-2))
	fmt.Println(board.GetDiagDownRight(-1))
	fmt.Println(board.GetDiagDownRight(0))
	fmt.Println(board.GetDiagDownRight(1))
	fmt.Println(board.GetDiagDownRight(2))
	fmt.Println(board.GetDiagDownRight(3))

	// Output:
	// []
	// [G]
	// [D H]
	// [A E I]
	// [B F]
	// [C]
	// []
}

func ExampleBoard_GetDiagUpRight() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(board.GetDiagUpRight(-3))
	fmt.Println(board.GetDiagUpRight(-2))
	fmt.Println(board.GetDiagUpRight(-1))
	fmt.Println(board.GetDiagUpRight(0))
	fmt.Println(board.GetDiagUpRight(1))
	fmt.Println(board.GetDiagUpRight(2))
	fmt.Println(board.GetDiagUpRight(3))

	// Output:
	// []
	// [A]
	// [D B]
	// [G E C]
	// [H F]
	// [I]
	// []
}
