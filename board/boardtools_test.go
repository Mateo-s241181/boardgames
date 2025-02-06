package board

import "fmt"

func ExampleRowEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", "Q", " "}
	board[2] = []string{" ", " ", " "}

	fmt.Println(RowEmpty(board, 0))
	fmt.Println(RowEmpty(board, 1))
	fmt.Println(RowEmpty(board, 2))
	fmt.Println(RowEmpty(board, -1))
	fmt.Println(RowEmpty(board, 3))

	// Output:
	// true
	// false
	// true
	// true
	// true
}

func ExampleColumnEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{" ", " ", " "}

	fmt.Println(ColumnEmpty(board, 0))
	fmt.Println(ColumnEmpty(board, 1))
	fmt.Println(ColumnEmpty(board, 2))
	fmt.Println(ColumnEmpty(board, -1))
	fmt.Println(ColumnEmpty(board, 3))

	// Output:
	// true
	// true
	// false
	// true
	// true
}

func ExampleDiagDownRightEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagDownRightEmpty(board, -3))
	fmt.Println(DiagDownRightEmpty(board, -2))
	fmt.Println(DiagDownRightEmpty(board, -1))
	fmt.Println(DiagDownRightEmpty(board, 0))
	fmt.Println(DiagDownRightEmpty(board, 1))
	fmt.Println(DiagDownRightEmpty(board, 2))
	fmt.Println(DiagDownRightEmpty(board, 3))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// true
	// true
}

func ExampleDiagUpRightEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagUpRightEmpty(board, -3))
	fmt.Println(DiagUpRightEmpty(board, -2))
	fmt.Println(DiagUpRightEmpty(board, -1))
	fmt.Println(DiagUpRightEmpty(board, 0))
	fmt.Println(DiagUpRightEmpty(board, 1))
	fmt.Println(DiagUpRightEmpty(board, 2))
	fmt.Println(DiagUpRightEmpty(board, 3))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// true
	// true
}
