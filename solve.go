package sudoku

import (
	"errors"
	"fmt"
	"strings"
)

const (
	boardSize     = 81
	boardSide     = 9
	boardHalfSize = 4
	boardBoxSize  = 3
	maxValue      = 9
)

var (
	errOutOfBoardIndex = errors.New("index is out of board")
	errOutOfBoundValue = errors.New("value is out of bounds, has to be between 1 - 9")
)

type Board struct {
	b []int
	e error
	s bool
}

type Game interface {
	SetClue(row, column, value int) Game
	Value(row, column int) int
	Row(row int) []int
	Column(column int) []int
	Box(boxIndex int) [][]int
	Board() [][]int
	IsEmpty(row, column int) bool
	IsValid() bool
	Solve()
	Error() error
	String() string
}

func NewBoard() Game {
	return &Board{
		b: make([]int, boardSize, boardSize), // board
		s: false,                             // solved
		e: nil,                               // error
	}
}

func (b *Board) SetClue(row, column, value int) Game {
	// do nothing when any error occurred
	if b.e != nil {
		return b
	}

	// check for board index value
	idx := b.index(row, column)
	if b.e != nil {
		return b
	}

	// check for value
	if value < 1 || value > maxValue {
		b.e = errOutOfBoundValue
		return b
	}

	b.b[idx] = value
	return b
}

func (b *Board) IsEmpty(row, column int) bool {
	// do nothing when any error occurred
	if b.e != nil {
		return false
	}

	// check for board index
	idx := b.index(row, column)
	if b.e != nil {
		return false
	}

	return b.b[idx] == 0
}

func (b Board) IsValid() bool {
	if b.e != nil {
		return false
	}

	for i := 0; i < boardSide; i++ {
		rowIdx := b.index(i, 0)
		columnIdx := b.index(0, i)
		if !b.isValidBox(i) || !b.isValidRow(rowIdx) || !b.isValidColumn(columnIdx) {
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

	idx := b.index(row, column)
	if b.e != nil {
		// non existing value + err already set
		return -1
	}

	return b.b[idx]
}

func (b Board) Row(row int) []int {
	if b.e != nil {
		return nil
	}

	idx := b.index(row, 0)
	if b.e != nil {
		return nil
	}

	return b.row(idx)
}

func (b Board) Column(column int) []int {
	if b.e != nil {
		return nil
	}

	idx := b.index(0, column)
	if b.e != nil {
		return nil
	}

	return b.column(idx)
}

func (b Board) Board() [][]int {
	board := make([][]int, boardSide, boardSide)
	for r := 0; r < boardSide; r++ {
		board[r] = make([]int, boardSide, boardSide)
		for c := 0; c < boardSide; c++ {
			board[r][c] = b.b[b.index(r, c)]
		}
	}

	return board
}

func (b *Board) Box(boxIndex int) [][]int {
	if boxIndex < 0 || boxIndex > boardSide {
		b.e = errOutOfBoardIndex
		return nil
	}

	box := make([][]int, boardBoxSize, boardBoxSize)
	bs := b.box(boxIndex)
	for i := 0; i < boardBoxSize; i++ {
		box[i] = make([]int, boardBoxSize, boardBoxSize)
		for j := 0; j < boardBoxSize; j++ {
			box[i][j] = bs[i*boardBoxSize+j]
		}
	}

	return box
}

func (b Board) Solve() {
	if b.e != nil || b.s {
		return
	}
	b.solve()
}

func (b Board) String() string {
	sb := strings.Builder{}

	for idx, i := 0, 1; idx < boardSize; i, idx = i+1, idx+1 {
		sb.WriteString(fmt.Sprintf("|%d", b.b[idx]))
		if i == boardSide {
			sb.WriteString("|\n")
			i = 0
		}
	}

	return sb.String()
}

// ------------------------------------------------- PRIVATE METHODS -------------------------------------------------

func (b *Board) index(row, column int) int {

	idx := row*boardSide + column
	if idx < 0 || idx > boardSize {
		b.e = errOutOfBoardIndex
	}

	return idx
}

func (b Board) row(idx int) []int {
	r := make([]int, boardSide, boardSide)
	for i, j := idx, 0; i < idx+boardSide; i, j = i+1, j+1 {
		r[j] = b.b[i]
	}

	return r
}

func (b Board) column(idx int) []int {
	c := make([]int, boardSide, boardSide)
	for i := 0; i < boardSide; i++ {
		c[i] = b.b[i*boardSide+idx]
	}

	return c
}

func (b Board) box(boxIndex int) []int {
	row := (boxIndex / boardBoxSize) * boardBoxSize
	column := (boxIndex % boardBoxSize) * boardBoxSize

	box := make([]int, boardBoxSize*boardBoxSize, boardBoxSize*boardBoxSize)
	i := 0
	for r := row; r < row+boardBoxSize; r++ {
		for c := column; c < column+boardBoxSize; c++ {
			idx := b.index(r, c)
			box[i] = b.b[idx]
			i++
		}
	}

	return box
}

func (b Board) isValidRow(idx int) bool {
	row := b.row(idx)
	return b.isValidSlice(row)
}

func (b Board) isValidColumn(idx int) bool {
	column := b.column(idx)
	return b.isValidSlice(column)
}

func (b Board) isValidBox(boxNumber int) bool {
	box := b.box(boxNumber)
	return b.isValidSlice(box)
}

func (b Board) emptyValueIndex() int {
	for i := 0; i < len(b.b); i++ {
		if b.b[i] == 0 {
			return i
		}
	}
	return -1
}

func (b *Board) solve() {
	// no need to check for error
	if b.s {
		return
	}

	idx := b.emptyValueIndex()
	if idx < 0 {
		b.s = true
		return
	}

	for !b.s {
		b.b[idx] += 1
		if b.b[idx] > 9 {
			b.b[idx] = 0
			return
		}

		if b.IsValid() {
			b.solve()
		}
	}
}

// limited functionality to Sudoku where values could be within a limit
func (b Board) isValidSlice(values []int) bool {
	if values == nil {
		return true
	}

	m := make([]int, maxValue+1, maxValue+1)
	for _, v := range values {
		if m[v] > 0 {
			return false
		}
		m[v] = v
	}

	return true
}
