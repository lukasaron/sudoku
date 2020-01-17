package main

import "github.com/lukasaron/sudoku"

func main() {
	b := sudoku.NewBoard().
		SetClue(0, 0, 4).
		SetClue(0, 2, 8).
		SetClue(0, 4, 9).
		SetClue(1, 1, 6).
		SetClue(1, 2, 2).
		SetClue(1, 4, 8).
		SetClue(1, 7, 1).
		SetClue(2, 0, 1).
		SetClue(2, 1, 9).
		SetClue(2, 5, 6).
		SetClue(2, 6, 7).
		SetClue(3, 0, 9).
		SetClue(3, 1, 3).
		SetClue(3, 4, 2).
		SetClue(3, 5, 8).
		SetClue(3, 7, 7).
		SetClue(4, 2, 7).
		SetClue(4, 3, 9).
		SetClue(4, 4, 4).
		SetClue(4, 8, 3).
		SetClue(5, 1, 1).
		SetClue(5, 3, 7).
		SetClue(5, 5, 3).
		SetClue(6, 1, 7).
		SetClue(6, 3, 8).
		SetClue(6, 6, 4).
		SetClue(6, 8, 6).
		SetClue(7, 0, 3).
		SetClue(7, 1, 2).
		SetClue(7, 3, 4).
		SetClue(7, 4, 7).
		SetClue(8, 0, 8).
		SetClue(8, 4, 1).
		SetClue(8, 6, 2).
		SetClue(8, 8, 7)
	b.Solve()
}
