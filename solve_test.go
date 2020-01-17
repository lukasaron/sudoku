package sudoku

import "testing"

func TestBoard_IsValid(t *testing.T) {
	g := game()
	if !g.IsValid() {
		t.Error("basic sudoku should be valid")
	}
}

func TestBoard_IsValidRow(t *testing.T) {
	g := game()
	if !g.IsValid() {
		t.Error("basic sudoku should be valid")
	}

	g.SetClue(8, 7, 1)
	if g.IsValid() {
		t.Error("there is a duplicated value in the row 8, column 7")
	}
}

func TestBoard_IsValidColumn(t *testing.T) {
	g := game()
	if !g.IsValid() {
		t.Error("basic game should be valid")
	}

	g.SetClue(4, 6, 2)
	if g.IsValid() {
		t.Error("there is a duplicated value in the row 4, column 6")
	}
}

func TestBoard_IsValidBox(t *testing.T) {
	g := game()
	if !g.IsValid() {
		t.Error("basic sudoku should be valid")
	}

	g.SetClue(7, 7, 6)
	if g.IsValid() {
		t.Error("there is a duplicated value in the box 8, value 6")
	}
}

// ------------------------------------------------------ DATA ------------------------------------------------------

func game() Game {
	return NewBoard().
		// first row - 0
		SetClue(0, 0, 4).
		SetClue(0, 2, 8).
		SetClue(0, 4, 9).
		// second row - 1
		SetClue(1, 1, 6).
		SetClue(1, 2, 2).
		SetClue(1, 4, 8).
		SetClue(1, 7, 1).
		// third row - 2
		SetClue(2, 0, 1).
		SetClue(2, 1, 9).
		SetClue(2, 5, 6).
		SetClue(2, 6, 7).
		// fourth row - 3
		SetClue(3, 0, 9).
		SetClue(3, 1, 3).
		SetClue(3, 4, 2).
		SetClue(3, 5, 8).
		SetClue(3, 7, 7).
		// fifth row - 4
		SetClue(4, 2, 7).
		SetClue(4, 3, 9).
		SetClue(4, 4, 4).
		SetClue(4, 8, 3).
		// sixth row - 5
		SetClue(5, 1, 1).
		SetClue(5, 3, 7).
		SetClue(5, 5, 3).
		// seventh row - 6
		SetClue(6, 1, 7).
		SetClue(6, 3, 8).
		SetClue(6, 6, 4).
		SetClue(6, 8, 6).
		// eight row - 7
		SetClue(7, 0, 3).
		SetClue(7, 1, 2).
		SetClue(7, 3, 4).
		SetClue(7, 4, 7).
		// ninth row - 8
		SetClue(8, 0, 8).
		SetClue(8, 4, 1).
		SetClue(8, 6, 2).
		SetClue(8, 8, 7)
}
