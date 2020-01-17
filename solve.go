package sudoku

import (
	"errors"
)

const (
	size    = 9
	boxSize = 3
)

var (
	errOutOfBoardIndex = errors.New("index is out of board")
	errOutOfBoundValue = errors.New("value is out of bounds, has to be between 1 - 9")
)

type Board struct {
	b [][]int
	e error
	s bool
}

type Game interface {
	SetClue(row, column, value int) Game
	GetValue(row, column int) int
	GetRow(row int) []int
	GetColumn(column int) []int
	GetBoard() [][]int
	IsEmpty(row, column int) bool
	IsValid() bool
	Solve()
	Error() error
}

func NewBoard() Game {
	board := make([][]int, size, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size, size)
	}

	return &Board{
		b: board, // board
		s: false, // solved
	}
}

func (b *Board) SetClue(row, column, value int) Game {
	// do nothing when any error occurred
	if b.e != nil {
		return b
	}

	// check for board index value
	if b.outOfBoard(row, column) {
		b.e = errOutOfBoardIndex
		return b
	}

	// check for value
	if value < 1 || value > 9 {
		b.e = errOutOfBoundValue
		return b
	}

	b.b[row][column] = value
	return b
}

func (b *Board) IsEmpty(row, column int) bool {
	// do nothing when any error occurred
	if b.e != nil {
		return false
	}

	// check for board index
	if b.outOfBoard(row, column) {
		b.e = errOutOfBoardIndex
		return false
	}

	return b.b[row][column] == 0
}

func (b *Board) IsValid() bool {
	if b.e != nil {
		return false
	}

	for i := 0; i < size; i++ {
		if !b.isValidBox(i) || !b.isValidRow(i) || !b.isValidColumn(i) {
			return false
		}
	}

	return true
}

func (b Board) Error() error {
	return b.e
}

func (b Board) GetValue(row, column int) int {
	if b.e != nil {
		return -1
	}

	if b.outOfBoard(row, column) {
		b.e = errOutOfBoardIndex
		return -1
	}

	return b.b[row][column]
}

func (b Board) GetRow(row int) []int {
	if b.e != nil {
		return nil
	}

	if b.outOfBoard(row, 0) {
		b.e = errOutOfBoardIndex
		return nil
	}

	r := make([]int, size, size)
	copy(r, b.b[row])
	return r
}

func (b Board) GetColumn(column int) []int {
	if b.e != nil {
		return nil
	}

	if b.outOfBoard(0, column) {
		b.e = errOutOfBoardIndex
		return nil
	}

	c := make([]int, size, size)
	for i := 0; i < size; i++ {
		c[i] = b.b[i][column]
	}

	return c
}

func (b Board) GetBoard() [][]int {
	board := make([][]int, size, size)
	for i := 0; i < size; i++ {
		copy(board[i], b.b[i])
	}

	return board
}

func (b Board) Solve() {
	if b.e != nil {
		return
	}

	panic("implement me")
}

// ------------------------------------------------- PRIVATE METHODS -------------------------------------------------

func (b Board) outOfBoard(row, column int) bool {
	return row < 0 || row > (size-1) || column < 0 || column > (size-1)
}

func (b Board) isValidRow(row int) bool {
	m := [size + 1]int{}
	for _, value := range b.b[row] {
		if m[value] > 0 {
			return false
		}
		m[value] = value
	}

	return true
}

func (b Board) isValidColumn(column int) bool {
	m := [size + 1]int{}
	for row := 0; row < size; row++ {
		value := b.b[row][column]
		if m[value] > 0 {
			return false
		}
		m[value] = value
	}

	return true
}

func (b Board) isValidBox(box int) bool {
	if box < 0 || box > 8 {
		return false
	}

	m := [size + 1]int{}
	row := (box / boxSize) * boxSize
	column := (box % boxSize) * boxSize

	for r := row; r < row+boxSize; r++ {
		for c := column; c < column+boxSize; c++ {
			if m[b.b[r][c]] > 0 {
				return false
			}
			m[b.b[r][c]] = b.b[r][c]
		}
	}

	return true
}

func (b Board) getNextEmptyIndex() (int, int) {
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if b.b[r][c] == 0 {
				return r, c
			}
		}
	}

	return -1, -1
}
