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
	Value(row, column int) int
	Row(row int) []int
	Column(column int) []int
	Board() [][]int
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

func (b Board) Value(row, column int) int {
	if b.e != nil {
		return -1
	}

	if b.outOfBoard(row, column) {
		b.e = errOutOfBoardIndex
		return -1
	}

	return b.b[row][column]
}

func (b Board) Row(row int) []int {
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

func (b Board) Column(column int) []int {
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

func (b Board) Board() [][]int {
	board := make([][]int, size, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size, size)
		copy(board[i], b.b[i])
	}

	return board
}

func (b Board) Solve() {
	if b.e != nil || b.s {
		return
	}
	b.solve()
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
	//n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(size * size)
	rowSeed := size / 2
	columnSeed := size / 2

	row, column := b.getNextEmpty(rowSeed, columnSeed)
	if row < 0 || column < 0 {
		row, column = b.getPreviousEmpty(rowSeed, columnSeed)
	}

	return row, column
}

func (b Board) getPreviousEmpty(row, column int) (int, int) {
	if b.b[row][column] == 0 {
		return row, column
	}

	if column > 0 {
		return b.getPreviousEmpty(row, column-1)
	}

	if row > 0 {
		return b.getPreviousEmpty(row-1, size-1)
	}

	return -1, -1
}

func (b Board) getNextEmpty(row, column int) (int, int) {
	if b.b[row][column] == 0 {
		return row, column
	}

	if column < size-1 {
		return b.getNextEmpty(row, column+1)
	}

	if row < size-1 {
		return b.getNextEmpty(row+1, column)
	}

	return -1, -1
}

func (b *Board) solve() {
	// no need to check for error
	if b.s {
		return
	}

	r, c := b.getNextEmptyIndex()
	if c < 0 || r < 0 {
		b.s = true
		return
	}

	for !b.s {
		b.b[r][c] += 1
		if b.b[r][c] > 9 {
			b.b[r][c] = 0
			return
		}

		if b.IsValid() {
			b.solve()
		}
	}
}
