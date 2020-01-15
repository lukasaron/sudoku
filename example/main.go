package main

import "github.com/lukasaron/sudoku"

func main() {
	b := sudoku.NewBoard()
	b.SetClue(4, 3, 7)
	b.SetClue(2, 1, 4)
}
