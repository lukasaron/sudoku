package sudoku

import "errors"

var (
	errOutOfBoardIndex = errors.New("index is out of board")
	errOutOfBoundValue = errors.New("value is out of bounds, has to be between 1 - 9")
)

type Board struct {
	b [9][9]int
	e error
}

type Game interface {
	SetClue(x, y, value int)
	IsEmpty(x, y int) bool
	IsValid() bool
}

func NewBoard() Game {
	return &Board{
		b: [9][9]int{},
	}
}

func (b *Board) SetClue(x, y, value int) {
	// do nothing when any error occurred
	if b.e != nil {
		return
	}

	// check for board index value
	if b.outOfBoard(x, y) {
		b.e = errOutOfBoardIndex
		return
	}

	// check for value
	if value < 1 || value > 9 {
		b.e = errOutOfBoundValue
		return
	}

	b.b[x][y] = value
}

func (b *Board) IsEmpty(x, y int) bool {
	// do nothing when any error occurred
	if b.e != nil {
		return false
	}

	// check for board index
	if b.outOfBoard(x, y) {
		b.e = errOutOfBoardIndex
		return false
	}

	return b.b[x][y] == 0
}

func (b *Board) IsValid() bool {
	if b.e != nil {
		return false
	}

	//TODO isValid - rows, boxes, rows
	return true
}

func (b Board) outOfBoard(x, y int) bool {
	return x < 0 || x > 8 || y < 0 || y > 8
}
