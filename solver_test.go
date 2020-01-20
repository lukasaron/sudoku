package sudoku

import (
	"reflect"
	"testing"
)

func TestBoard_SetValue(t *testing.T) {
	g := NewBoard().SetValue(-1, 0, 0)
	if g.Error() == nil {
		t.Error("error expected when set value outside the board")
	}

	g = NewBoard().SetValue(9, 0, 0)
	if g.Error() == nil {
		t.Error("error expected when set value outside the board")
	}

	g = NewBoard().SetValue(0, -1, 0)
	if g.Error() == nil {
		t.Error("error expected when set value outside the board")
	}

	g = NewBoard().SetValue(0, 9, 0)
	if g.Error() == nil {
		t.Error("error expected when set value outside the board")
	}

	g = NewBoard().SetValue(4, 3, 10)
	if g.Error() == nil {
		t.Error("error expected when set value bigger than allowed")
	}

	g = NewBoard().SetValue(4, 3, 7)
	if g.Error() != nil {
		t.Errorf("error not expected, got: %v", g.Error())
	}

	if g.Value(4, 3) != 7 {
		t.Error("value set, but not retrieved")
	}
}

func TestBoard_SetRow(t *testing.T) {

	g := NewBoard().SetRow(0, nil)
	if g.Error() == nil {
		t.Error("error expected when set row has no value")
	}

	g = NewBoard().SetRow(0, []int{})
	if g.Error() == nil {
		t.Error("error expected when set row has no value")
	}

	g = NewBoard().SetRow(0, []int{0})
	if g.Error() == nil {
		t.Error("error expected when set row has wrong length")
	}

	g = NewBoard().SetRow(0, []int{0, 0, 0, 0, 0, 0, 0, 0, 10})
	if g.Error() == nil {
		t.Error("error expected when set row has wrong value")
	}

	g = NewBoard().SetRow(0, []int{1, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Error("error expected when set row has a duplicated not empty value")
	}

	g = NewBoard().SetRow(-1, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Errorf("error expected when the row is out of a board")
	}

	g = NewBoard().SetRow(9, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Errorf("error expected when the row is out of a board")
	}

	g = NewBoard().SetRow(0, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() != nil {
		t.Errorf("error not expected, got: %v", g.Error())
	}

	if !reflect.DeepEqual(g.Row(0), []int{0, 0, 0, 0, 0, 0, 0, 0, 1}) {
		t.Error("row set, but not retrieved")
	}
}

func TestBoard_SetColumn(t *testing.T) {
	g := NewBoard().SetColumn(0, nil)
	if g.Error() == nil {
		t.Error("error expected when set column has no value")
	}

	g = NewBoard().SetColumn(0, []int{})
	if g.Error() == nil {
		t.Error("error expected when set column has empty value")
	}

	g = NewBoard().SetColumn(0, []int{0})
	if g.Error() == nil {
		t.Error("error expected when set column has wrong length")
	}

	g = NewBoard().SetColumn(0, []int{0, 0, 0, 0, 0, 0, 0, 0, 10})
	if g.Error() == nil {
		t.Error("error expected when set column has wrong value")
	}

	g = NewBoard().SetColumn(0, []int{1, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Error("error expected when set column has a duplicated not empty value")
	}

	g = NewBoard().SetColumn(-1, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Errorf("error expected when the column is out of a board")
	}

	g = NewBoard().SetColumn(9, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Errorf("error expected when the column is out of a board")
	}

	g = NewBoard().SetColumn(0, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() != nil {
		t.Errorf("error not expected, got: %v", g.Error())
	}

	if !reflect.DeepEqual(g.Column(0), []int{0, 0, 0, 0, 0, 0, 0, 0, 1}) {
		t.Error("column set, but not retrieved correctly")
	}
}

func TestBoard_SetBox(t *testing.T) {
	g := NewBoard().SetBox(0, nil)
	if g.Error() == nil {
		t.Error("error expected when set box has no value")
	}

	g = NewBoard().SetBox(0, []int{})
	if g.Error() == nil {
		t.Error("error expected when set box has empty")
	}

	g = NewBoard().SetBox(0, []int{0})
	if g.Error() == nil {
		t.Error("error expected when set box has wrong length")
	}

	g = NewBoard().SetBox(0, []int{0, 0, 0, 0, 0, 0, 0, 0, 10})
	if g.Error() == nil {
		t.Error("error expected when set box has wrong value")
	}

	g = NewBoard().SetBox(0, []int{1, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Error("error expected when set box has a duplicated not empty value")
	}

	g = NewBoard().SetBox(-1, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Errorf("error expected when the box index is out of a board")
	}

	g = NewBoard().SetBox(9, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() == nil {
		t.Errorf("error expected when the box index is out of a board")
	}

	g = NewBoard().SetBox(0, []int{0, 0, 0, 0, 0, 0, 0, 0, 1})
	if g.Error() != nil {
		t.Errorf("error not expected, got: %v", g.Error())
	}

	if !reflect.DeepEqual(g.Box(0), []int{0, 0, 0, 0, 0, 0, 0, 0, 1}) {
		t.Error("box set, but not retrieved correctly")
	}
}

func TestBoard_SetBoard(t *testing.T) {
	g := NewBoard().SetBoard(nil)
	if g.Error() == nil {
		t.Error("error expected when set board has no value")
	}

	g = NewBoard().SetBoard([][]int{})
	if g.Error() == nil {
		t.Error("error expected when set board has empty value")
	}

	g = NewBoard().SetBoard([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	if g.Error() == nil {
		t.Error("error expected when set board has wrong number of rows")
	}

	// 10 rows
	g = NewBoard().SetBoard([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	if g.Error() == nil {
		t.Error("error expected when set board has wrong number of rows")
	}

	g = NewBoard().SetBoard([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	})
	if g.Error() == nil {
		t.Error("error expected when set board has wrong number of columns")
	}

	row := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	g = NewBoard().SetBoard([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		row,
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	if g.Error() != nil {
		t.Errorf("error not expected and got: %v", g.Error())
	}

	if !reflect.DeepEqual(g.Row(3), row) {
		t.Error("row set via set board, but not retrieved correctly")
	}

	column := g.Column(3)
	if g.Error() != nil {
		t.Errorf("error not expected and got: %v", g.Error())
	}

	if len(column) != 9 || column[3] != 4 {
		t.Error("column set via set board, but not retrieved correctly")
	}

	board := [][]int{
		{0, 0, 0, 0, 0, 0, 1, 4, 8},
		{0, 1, 0, 0, 2, 6, 0, 0, 3},
		{0, 0, 0, 0, 1, 0, 6, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 0, 2},
		{1, 0, 0, 3, 6, 2, 0, 0, 7},
		{5, 0, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 5, 0, 3, 0, 0, 0, 0},
		{3, 0, 0, 1, 9, 0, 0, 7, 0},
		{4, 7, 0, 0, 0, 0, 0, 0, 0},
	}
	g = NewBoard().SetBoard(board)
	if !reflect.DeepEqual(g.Board(), board) {
		t.Error("board set, but not retrieved correctly")
	}
}

func TestBoard_Error(t *testing.T) {
	g := NewBoard()
	if g.Error() != nil {
		t.Errorf("error not expected after initialisation, got: %v", g.Error())
	}
}

func TestBoard_IsEmpty(t *testing.T) {
	g := NewBoard()
	if g.IsEmpty(-1, 0) {
		t.Error("out of board index shouldn't be empty")
	}

	if g.Error() == nil {
		t.Error("out of board index should set an error")
	}

	g = NewBoard()
	if g.IsEmpty(9, 0) {
		t.Error("out of board index shouldn't be empty")
	}

	if g.Error() == nil {
		t.Error("out of board index should set an error")
	}

	g = NewBoard()
	if g.IsEmpty(0, -1) {
		t.Error("out of board index shouldn't be empty")
	}

	if g.Error() == nil {
		t.Error("out of board index should set an error")
	}

	g = NewBoard()
	if g.IsEmpty(0, 9) {
		t.Error("out of board index shouldn't be empty")
	}

	if g.Error() == nil {
		t.Error("out of board index should set an error")
	}

	g = NewBoard()
	if !g.IsEmpty(0, 0) {
		t.Error("new board should have all values empty")
	}

	g.SetValue(0, 0, 5)
	if g.IsEmpty(0, 0) {
		t.Error("set value shouldn't be empty")
	}
}

func TestBoard_IsValid(t *testing.T) {
	g := easyGame()
	if !g.IsValid() {
		t.Error("basic sudoku should be valid")
	}
}

func TestBoard_IsValidRow(t *testing.T) {
	g := easyGame()

	g.SetValue(3, 0, 9)
	if g.IsValid() {
		t.Error("there is a duplicated value [9] in the row 3, column 0")
	}
}

func TestBoard_IsValidColumn(t *testing.T) {
	g := easyGame()

	g.SetValue(0, 4, 6)
	if g.IsValid() {
		t.Error("there is a duplicated value [6] in the row 0, column 4")
	}
}

func TestBoard_IsValidBox(t *testing.T) {
	g := easyGame()

	g.SetValue(1, 7, 1)
	if g.IsValid() {
		t.Error("there is a duplicated value [1] in the box 2")
	}
}

func TestBoard_SolveEasy(t *testing.T) {
	g := easyGame()
	g.Solve()
	err := g.Error()
	if err != nil {
		t.Error(err)
	}

	s := easyGameSolved()
	if !reflect.DeepEqual(g.Board(), s.Board()) {
		t.Errorf("sudoku solved expected:\n %s, got:\n %s", s, g)
	}
}

func TestBoard_SolveHard(t *testing.T) {
	g := hardGame()
	g.Solve()
	err := g.Error()
	if err != nil {
		t.Error(err)
	}

	s := hardGameSolved()
	if !reflect.DeepEqual(g.Board(), s.Board()) {
		t.Errorf("sudoku solved expected:\n%s, got:\n%s", s, g)
	}
}

// ------------------------------------------------------ DATA ------------------------------------------------------

func easyGame() Game {
	return NewBoard().SetBoard([][]int{
		{0, 0, 0, 0, 0, 0, 1, 4, 8},
		{0, 1, 0, 0, 2, 6, 0, 0, 3},
		{0, 0, 0, 0, 1, 0, 6, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 0, 2},
		{1, 0, 0, 3, 6, 2, 0, 0, 7},
		{5, 0, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 5, 0, 3, 0, 0, 0, 0},
		{3, 0, 0, 1, 9, 0, 0, 7, 0},
		{4, 7, 0, 0, 0, 0, 0, 0, 0},
	})
}

func easyGameSolved() Game {
	return NewBoard().SetBoard([][]int{
		{2, 5, 6, 9, 7, 3, 1, 4, 8},
		{8, 1, 4, 5, 2, 6, 7, 9, 3},
		{7, 3, 9, 4, 1, 8, 6, 2, 5},
		{6, 4, 3, 7, 5, 1, 9, 8, 2},
		{1, 9, 8, 3, 6, 2, 4, 5, 7},
		{5, 2, 7, 8, 4, 9, 3, 6, 1},
		{9, 6, 5, 2, 3, 7, 8, 1, 4},
		{3, 8, 2, 1, 9, 4, 5, 7, 6},
		{4, 7, 1, 6, 8, 5, 2, 3, 9},
	})
}

func hardGame() Game {
	return NewBoard().SetBoard([][]int{
		{4, 0, 0, 0, 0, 7, 5, 0, 0},
		{0, 1, 7, 0, 5, 0, 4, 0, 0},
		{9, 5, 0, 0, 4, 3, 7, 2, 0},
		{0, 0, 0, 4, 0, 0, 6, 0, 5},
		{7, 0, 5, 0, 0, 0, 1, 0, 2},
		{0, 0, 1, 0, 0, 5, 0, 0, 0},
		{5, 8, 3, 6, 0, 0, 0, 0, 9},
		{0, 0, 9, 0, 8, 0, 0, 5, 0},
		{0, 0, 0, 5, 0, 0, 0, 0, 6},
	})
}

func hardGameSolved() Game {
	return NewBoard().SetBoard([][]int{
		{4, 2, 8, 1, 9, 7, 5, 6, 3},
		{3, 1, 7, 2, 5, 6, 4, 9, 8},
		{9, 5, 6, 8, 4, 3, 7, 2, 1},
		{8, 9, 2, 4, 7, 1, 6, 3, 5},
		{7, 3, 5, 9, 6, 8, 1, 4, 2},
		{6, 4, 1, 3, 2, 5, 9, 8, 7},
		{5, 8, 3, 6, 1, 4, 2, 7, 9},
		{1, 6, 9, 7, 8, 2, 3, 5, 4},
		{2, 7, 4, 5, 3, 9, 8, 1, 6},
	})
}

func game2() Game {
	return NewBoard().
		SetValue(1, 5, 3).
		SetValue(1, 7, 8).
		SetValue(1, 8, 5).
		SetValue(2, 2, 1).
		SetValue(2, 4, 2).
		SetValue(3, 3, 5).
		SetValue(3, 5, 7).
		SetValue(4, 2, 4).
		SetValue(4, 6, 1).
		SetValue(5, 1, 9).
		SetValue(6, 0, 5).
		SetValue(6, 7, 7).
		SetValue(6, 8, 3).
		SetValue(7, 2, 2).
		SetValue(7, 4, 1).
		SetValue(8, 4, 4).
		SetValue(8, 8, 9)
}
