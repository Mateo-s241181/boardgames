package board

import "fmt"

func ExampleRowContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", "Q", " "}
	board[2] = []string{" ", " ", " "}

	fmt.Println(RowContainsQueen(board, 0))
	fmt.Println(RowContainsQueen(board, 1))
	fmt.Println(RowContainsQueen(board, 2))
	fmt.Println(RowContainsQueen(board, -1))
	fmt.Println(RowContainsQueen(board, 3))

	// Output:
	// false
	// true
	// false
	// false
	// false
}

func ExampleColumnContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{" ", " ", " "}

	fmt.Println(ColumnContainsQueen(board, 0))
	fmt.Println(ColumnContainsQueen(board, 1))
	fmt.Println(ColumnContainsQueen(board, 2))
	fmt.Println(ColumnContainsQueen(board, -1))
	fmt.Println(ColumnContainsQueen(board, 3))

	// Output:
	// false
	// false
	// true
	// false
	// false
}

func ExampleDiagDownRightContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagDownRightContainsQueen(board, -3))
	fmt.Println(DiagDownRightContainsQueen(board, -2))
	fmt.Println(DiagDownRightContainsQueen(board, -1))
	fmt.Println(DiagDownRightContainsQueen(board, 0))
	fmt.Println(DiagDownRightContainsQueen(board, 1))
	fmt.Println(DiagDownRightContainsQueen(board, 2))
	fmt.Println(DiagDownRightContainsQueen(board, 3))

	// Output:
	// false
	// true
	// true
	// false
	// true
	// false
	// false
}

func ExampleDiagUpRightContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagUpRightContainsQueen(board, -3))
	fmt.Println(DiagUpRightContainsQueen(board, -2))
	fmt.Println(DiagUpRightContainsQueen(board, -1))
	fmt.Println(DiagUpRightContainsQueen(board, 0))
	fmt.Println(DiagUpRightContainsQueen(board, 1))
	fmt.Println(DiagUpRightContainsQueen(board, 2))
	fmt.Println(DiagUpRightContainsQueen(board, 3))

	// Output:
	// false
	// false
	// false
	// true
	// true
	// false
	// false
}
