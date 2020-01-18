package main

import (
	"fmt"
	"github.com/lukasaron/sudoku"
)

func main() {
	b := sudoku.NewBoard().
		SetClue(1, 5, 3).
		SetClue(1, 7, 8).
		SetClue(1, 8, 5).
		SetClue(2, 2, 1).
		SetClue(2, 4, 2).
		SetClue(3, 3, 5).
		SetClue(3, 5, 7).
		SetClue(4, 2, 4).
		SetClue(4, 6, 1).
		SetClue(5, 1, 9).
		SetClue(6, 0, 5).
		SetClue(6, 7, 7).
		SetClue(6, 8, 3).
		SetClue(7, 2, 2).
		SetClue(7, 4, 1).
		SetClue(8, 4, 4).
		SetClue(8, 8, 9)

	b.Solve()
	fmt.Printf("%+v\n", b.Board())
}
