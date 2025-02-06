package board

import "fmt"

func ExampleBoard_RowEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", "Q", " "}
	board[2] = []string{" ", " ", " "}

	fmt.Println(board.RowEmpty(0))
	fmt.Println(board.RowEmpty(1))
	fmt.Println(board.RowEmpty(2))
	fmt.Println(board.RowEmpty(-1))
	fmt.Println(board.RowEmpty(3))

	// Output:
	// true
	// false
	// true
	// true
	// true
}

func ExampleBoard_ColumnEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{" ", " ", " "}

	fmt.Println(board.ColumnEmpty(0))
	fmt.Println(board.ColumnEmpty(1))
	fmt.Println(board.ColumnEmpty(2))
	fmt.Println(board.ColumnEmpty(-1))
	fmt.Println(board.ColumnEmpty(3))

	// Output:
	// true
	// true
	// false
	// true
	// true
}

func ExampleBoard_DiagDownRightEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(board.DiagDownRightEmpty(-3))
	fmt.Println(board.DiagDownRightEmpty(-2))
	fmt.Println(board.DiagDownRightEmpty(-1))
	fmt.Println(board.DiagDownRightEmpty(0))
	fmt.Println(board.DiagDownRightEmpty(1))
	fmt.Println(board.DiagDownRightEmpty(2))
	fmt.Println(board.DiagDownRightEmpty(3))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// true
	// true
}

func ExampleBoard_DiagUpRightEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(board.DiagUpRightEmpty(-3))
	fmt.Println(board.DiagUpRightEmpty(-2))
	fmt.Println(board.DiagUpRightEmpty(-1))
	fmt.Println(board.DiagUpRightEmpty(0))
	fmt.Println(board.DiagUpRightEmpty(1))
	fmt.Println(board.DiagUpRightEmpty(2))
	fmt.Println(board.DiagUpRightEmpty(3))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// true
	// true
}
