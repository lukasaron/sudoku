# Sudoku 

[![Build Status](https://travis-ci.com/lukasaron/sudoku.svg?branch=master)](https://travis-ci.com/lukasaron/sudoku)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

Package sudoku contains the solver for the well known and ubiquitous Sudoku game.
The sudoku package is just a library that can be used by anyone. Based on this fact, there is not an executable
part of this repository.

When the Sudoku Board is created there are many ways how to fill clues 
(the original values before the user input any other guesses). A clue can be set by a single value, row, column or 
entirely whole board at once.

Sudoku can be solved by calling the method `Solve` on created Sudoku instance. Moreover, the whole sudoku can be
printed in any state, because implements the `Stringer` interface

Example of basic usage:
```go
		package main

		import (
			"fmt"
			"github.com/lukasaron/sudoku"
			"log"
		)

		func main() {
			game := sudoku.NewBoard().SetBoard([][]int{
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

			err := game.Error()
			if err != nil {
				log.Fatal(err)
			}

			game.Solve()
			fmt.Println(game)
		}
		package main

		import (
			"fmt"
			"github.com/lukasaron/sudoku"
			"log"
		)

		func main() {
			game := sudoku.NewBoard().SetBoard([][]int{
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

			err := game.Error()
			if err != nil {
				log.Fatal(err)
			}

			game.Solve()
			fmt.Println(game)
		}
```