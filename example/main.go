package main

import (
	"fmt"
	"github.com/lukasaron/sudoku"
)

func main() {
	b := sudoku.NewBoard().
		// []int{0,0,0,0,0,0,0,0,0}
		SetRow(1, []int{0, 0, 0, 0, 0, 3, 0, 8, 5}).
		SetRow(2, []int{0, 0, 1, 0, 2, 0, 0, 0, 0}).
		SetRow(3, []int{0, 0, 0, 5, 0, 7, 0, 0, 0}).
		SetRow(4, []int{0, 0, 4, 0, 0, 0, 1, 0, 0}).
		SetRow(5, []int{0, 9, 0, 0, 0, 0, 0, 0, 0}).
		SetRow(6, []int{5, 0, 0, 0, 0, 0, 0, 7, 3}).
		SetRow(7, []int{0, 0, 2, 0, 1, 0, 0, 0, 0}).
		SetRow(8, []int{0, 0, 0, 0, 4, 0, 0, 0, 9})
	b.Solve()
	fmt.Printf("%+v\n", b)
}
