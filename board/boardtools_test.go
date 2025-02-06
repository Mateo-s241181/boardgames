package board

import "fmt"

func ExampleBoard_RowContainsOnly() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", "Q", " "}
	board[2] = []string{" ", " ", " "}

	fmt.Println(board.RowContainsOnly(0, " "))
	fmt.Println(board.RowContainsOnly(1, " "))
	fmt.Println(board.RowContainsOnly(2, " "))
	fmt.Println(board.RowContainsOnly(-1, " "))
	fmt.Println(board.RowContainsOnly(3, " "))

	// Output:
	// true
	// false
	// true
	// true
	// true
}

func ExampleBoard_ColumnContainsOnly() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{" ", " ", " "}

	fmt.Println(board.ColumnContainsOnly(0, " "))
	fmt.Println(board.ColumnContainsOnly(1, " "))
	fmt.Println(board.ColumnContainsOnly(2, " "))
	fmt.Println(board.ColumnContainsOnly(-1, " "))
	fmt.Println(board.ColumnContainsOnly(3, " "))

	// Output:
	// true
	// true
	// false
	// true
	// true
}

func ExampleBoard_DiagDownRightContainsOnly() {
	board1 := EmptyBoard(3)
	board1[0] = []string{"Q", " ", " "}
	board1[1] = []string{" ", "Q", " "}
	board1[2] = []string{" ", " ", "Q"}

	board2 := EmptyBoard(3)
	board2[0] = []string{"Q", " ", " "}
	board2[1] = []string{" ", " ", " "}
	board2[2] = []string{" ", " ", "Q"}

	board3 := EmptyBoard(3)
	board3[0] = []string{" ", " ", " "}
	board3[1] = []string{" ", " ", " "}
	board3[2] = []string{" ", " ", " "}

	fmt.Println(board1.DiagDownRightContainsOnly("Q"))
	fmt.Println(board1.DiagDownRightContainsOnly(" "))

	fmt.Println(board2.DiagDownRightContainsOnly("Q"))
	fmt.Println(board2.DiagDownRightContainsOnly(" "))

	fmt.Println(board3.DiagDownRightContainsOnly("Q"))
	fmt.Println(board3.DiagDownRightContainsOnly(" "))

	// Output:
	// true
	// false
	// false
	// false
	// false
	// true

}

func ExampleBoard_DiagUpRightContainsOnly() {
	board1 := EmptyBoard(3)
	board1[0] = []string{" ", " ", "Q"}
	board1[1] = []string{" ", "Q", " "}
	board1[2] = []string{"Q", " ", " "}

	board2 := EmptyBoard(3)
	board2[0] = []string{" ", " ", "Q"}
	board2[1] = []string{" ", " ", " "}
	board2[2] = []string{"Q", " ", " "}

	board3 := EmptyBoard(3)
	board3[0] = []string{" ", " ", " "}
	board3[1] = []string{" ", " ", " "}
	board3[2] = []string{" ", " ", " "}

	fmt.Println(board1.DiagUpRightContainsOnly("Q"))
	fmt.Println(board1.DiagUpRightContainsOnly(" "))

	fmt.Println(board2.DiagUpRightContainsOnly("Q"))
	fmt.Println(board2.DiagUpRightContainsOnly(" "))

	fmt.Println(board3.DiagUpRightContainsOnly("Q"))
	fmt.Println(board3.DiagUpRightContainsOnly(" "))

	// Output:
	// true
	// false
	// false
	// false
	// false
	// true
}

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
