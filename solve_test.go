package sudoku

import (
	"fmt"
	"testing"
)

func TestBoard_IsValid(t *testing.T) {
	g := easyGame()
	if !g.IsValid() {
		t.Error("basic sudoku should be valid")
	}
}

func TestBoard_IsValidRow(t *testing.T) {
	g := easyGame()
	if !g.IsValid() {
		t.Error("basic sudoku should be valid")
	}

	g.SetValue(8, 7, 1)
	if g.IsValid() {
		t.Error("there is a duplicated value in the row 8, column 7")
	}
}

func TestBoard_IsValidColumn(t *testing.T) {
	g := easyGame()
	if !g.IsValid() {
		t.Error("basic easyGame should be valid")
	}

	g.SetValue(4, 6, 2)
	if g.IsValid() {
		t.Error("there is a duplicated value in the row 4, column 6")
	}
}

func TestBoard_IsValidBox(t *testing.T) {
	g := easyGame()
	if !g.IsValid() {
		t.Error("basic sudoku should be valid")
	}

	g.SetValue(7, 7, 6)
	if g.IsValid() {
		t.Error("there is a duplicated value in the box 8, value 6")
	}
}

func TestBoard_Solve(t *testing.T) {
	g := game2()
	g.Solve()
	err := g.Error()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(g)
}

// ------------------------------------------------------ DATA ------------------------------------------------------

func easyGame() Game {
	return NewBoard().
		// first row - 0
		SetValue(0, 0, 4).
		SetValue(0, 2, 8).
		SetValue(0, 4, 9).
		// second row - 1
		SetValue(1, 1, 6).
		SetValue(1, 2, 2).
		SetValue(1, 4, 8).
		SetValue(1, 7, 1).
		// third row - 2
		SetValue(2, 0, 1).
		SetValue(2, 1, 9).
		SetValue(2, 5, 6).
		SetValue(2, 6, 7).
		// fourth row - 3
		SetValue(3, 0, 9).
		SetValue(3, 1, 3).
		SetValue(3, 4, 2).
		SetValue(3, 5, 8).
		SetValue(3, 7, 7).
		// fifth row - 4
		SetValue(4, 2, 7).
		SetValue(4, 3, 9).
		SetValue(4, 4, 4).
		SetValue(4, 8, 3).
		// sixth row - 5
		SetValue(5, 1, 1).
		SetValue(5, 3, 7).
		SetValue(5, 5, 3).
		// seventh row - 6
		SetValue(6, 1, 7).
		SetValue(6, 3, 8).
		SetValue(6, 6, 4).
		SetValue(6, 8, 6).
		// eight row - 7
		SetValue(7, 0, 3).
		SetValue(7, 1, 2).
		SetValue(7, 3, 4).
		SetValue(7, 4, 7).
		// ninth row - 8
		SetValue(8, 0, 8).
		SetValue(8, 4, 1).
		SetValue(8, 6, 2).
		SetValue(8, 8, 7)
}

func hardGame() Game {
	return NewBoard().
		SetValue(0, 0, 8).
		SetValue(0, 2, 1).
		SetValue(0, 7, 4).
		SetValue(0, 8, 5).
		SetValue(1, 6, 7).
		SetValue(1, 8, 6).
		SetValue(2, 1, 5).
		SetValue(2, 2, 6).
		SetValue(2, 6, 8).
		SetValue(3, 1, 9).
		SetValue(3, 3, 7).
		SetValue(3, 6, 1).
		SetValue(4, 4, 8).
		SetValue(5, 3, 2).
		SetValue(5, 6, 5).
		SetValue(5, 7, 3).
		SetValue(5, 8, 8).
		SetValue(6, 4, 4).
		SetValue(6, 7, 8).
		SetValue(7, 0, 4).
		SetValue(7, 1, 2).
		SetValue(7, 2, 7).
		SetValue(7, 7, 1).
		SetValue(8, 4, 9).
		SetValue(8, 8, 4)
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
