package sudoku

import "errors"

const (
	rows    = 9
	columns = 9
)

var (
	errOutOfBoardIndex = errors.New("index is out of board")
	errOutOfBoundValue = errors.New("value is out of bounds, has to be between 1 - 9")
)

type Board struct {
	b [rows][columns]int
	e error
}

type Game interface {
	SetClue(row, column, value int)
	IsEmpty(row, column int) bool
	IsValid() bool
}

func NewBoard() Game {
	return &Board{
		b: [rows][columns]int{},
	}
}

func (b *Board) SetClue(row, column, value int) {
	// do nothing when any error occurred
	if b.e != nil {
		return
	}

	// check for board index value
	if b.outOfBoard(row, column) {
		b.e = errOutOfBoardIndex
		return
	}

	// check for value
	if value < 1 || value > 9 {
		b.e = errOutOfBoundValue
		return
	}

	b.b[row][column] = value
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

	//TODO isValid - rows, boxes, rows
	return true
}

// ------------------------------------------------- PRIVATE METHODS -------------------------------------------------

func (b Board) outOfBoard(row, column int) bool {
	return row < 0 || row > (rows-1) || column < 0 || column > (columns-1)
}

func (b Board) isValidRow(row int) bool {
	m := [rows + 1]int{}
	for _, value := range b.b[row] {
		if m[value] > 0 {
			return false
		}
		m[value] = value
	}

	return true
}

func (b Board) isValidColumn(column int) bool {
	m := [columns + 1]int{}
	for row := 0; row < rows; row++ {
		value := b.b[row][column]
		if m[value] > 0 {
			return false
		}
		m[value] = value
	}

	return true
}

func (b Board) isValidBox(box int) bool {
	return true
}
