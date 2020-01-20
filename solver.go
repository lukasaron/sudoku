package sudoku

import (
	"errors"
	"fmt"
	"strings"
)

const (
	BoardSize    = 81
	BoardSide    = 9
	BoardBoxSize = 3
	MaxValue     = 9
)

// Errors that could occur during user communication with the Game.
var (
	errOutOfBoardIndex = errors.New("index is out of board")
	errWrongInput      = errors.New("wrong input value(s)")
)

// Board is the implementation of the Game interface.
type Board struct {
	b []uint
	e error
	s bool
}

// Game interface defines basic methods that could be useful when interacting with the Sudoku.
type Game interface {
	SetValue(row, column, value int) Game
	SetRow(row int, values []int) Game
	SetColumn(column int, values []int) Game
	SetBox(boxIndex int, values []int) Game
	SetBoard(values [][]int) Game
	Value(row, column int) int
	Row(row int) []int
	Column(column int) []int
	Box(boxIndex int) []int
	Board() [][]int
	IsEmpty(row, column int) bool
	IsValid() bool
	Solve()
	Error() error
}

// NewBoard method creates Game with the Board instance.
func NewBoard() Game {
	return &Board{
		b: make([]uint, BoardSize, BoardSize), // board
		s: false,                              // solved
		e: nil,                                // error
	}
}

// SetValue method sets the value in the specific board coordinate. When there is a state error this method returns
// Game, but there is no behavior.
func (b *Board) SetValue(row, column, value int) Game {
	// do nothing when any error occurred
	if b.e != nil {
		return b
	}

	// check for board index value
	idx, err := b.index(row, column)
	if err != nil {
		b.e = err
		return b
	}

	// check for value
	if value < 0 || value > MaxValue {
		b.e = errWrongInput
		return b
	}

	b.b[idx] = uint(value)
	return b
}

// SetRow method sets the entire row. When there is a state error this method returns Game, but there is no behavior.
func (b *Board) SetRow(row int, values []int) Game {
	// do nothing when any error occurred
	if b.e != nil {
		return b
	}

	if !b.isValidSlice(values) {
		b.e = errWrongInput
		return b
	}

	idx, err := b.index(row, 0)
	if err != nil {
		b.e = err
		return b
	}

	for _, v := range values {
		b.b[idx] = uint(v)
		idx++
	}

	return b
}

// SetColumn method sets the entire column. When there is a state error this method returns Game, but there is no
// behavior.
func (b *Board) SetColumn(column int, values []int) Game {
	// do nothing when any error occurred
	if b.e != nil {
		return b
	}

	if !b.isValidSlice(values) {
		b.e = errWrongInput
		return b
	}

	idx, err := b.index(0, column)
	if err != nil {
		b.e = err
		return b
	}

	for _, v := range values {
		b.b[idx] = uint(v)
		idx += BoardSide
	}

	return b
}

// SetBox method sets the box of 3x3 size into the specific place - boxIndex that has to be between 0 - 8.
// When there is a state error this method returns Game, but there is no behavior.
func (b *Board) SetBox(boxIndex int, values []int) Game {
	// do nothing when any error occurred
	if b.e != nil {
		return b
	}

	if boxIndex < 0 || boxIndex >= BoardSide || !b.isValidSlice(values) {
		b.e = errWrongInput
		return b
	}

	idx, err := b.indexBox(boxIndex)
	if err != nil {
		b.e = err
		return b
	}

	// i is initialised in the first (outside) cycle, however incremented in the inner one
	for i, r := 0, 0; r < BoardBoxSize; r, idx = r+1, idx+BoardSide {
		for c := 0; c < BoardBoxSize; c, i = c+1, i+1 {
			b.b[idx+c] = uint(values[i])
		}
	}

	return b
}

// SetBoard method sets the entire board. When there is a state error this method has no behavior.
func (b *Board) SetBoard(values [][]int) Game {
	// do nothing when any error occurred
	if b.e != nil {
		return b
	}

	if len(values) != BoardSide {
		b.e = errWrongInput
		return b
	}

	for i, row := range values {
		b.SetRow(i, row)
	}

	// rows checked, but columns not -> validate
	if !b.IsValid() {
		b.e = errWrongInput
	}

	return b
}

// IsEmpty method checks if the value in the specific coordinates is empty (below 1).
// When there is a state error this method returns false.
func (b *Board) IsEmpty(row, column int) bool {
	// do nothing when any error occurred
	if b.e != nil {
		return false
	}

	// check for board index
	idx, err := b.index(row, column)
	if err != nil {
		b.e = err
		return false
	}

	return b.b[idx] < 1
}

// IsValid method checks if the board is valid, which means all values in the row, column and/or box are not duplicated.
// When there is a state error this method returns false.
func (b Board) IsValid() bool {
	// do nothing when any error occurred
	if b.e != nil {
		return false
	}

	for i := 0; i < BoardSide; i++ {
		rowIdx, _ := b.index(i, 0)
		columnIdx, _ := b.index(0, i)
		if !b.isValidBox(i) || !b.isValidRow(rowIdx) || !b.isValidColumn(columnIdx) {
			return false
		}
	}

	return true
}

// Error method returns status error if there is any.
func (b Board) Error() error {
	return b.e
}

// Value method returns the value in the specific coordinates. If there is a state error the value is equal -1.
func (b Board) Value(row, column int) int {
	// do nothing when any error occurred
	if b.e != nil {
		return -1
	}

	idx, err := b.index(row, column)
	if err != nil {
		b.e = err
		// non existing value
		return -1
	}

	return int(b.b[idx])
}

// Row method returns the entire row on a specific coordinate, which starts at 0 and the maximal value is 8.
// When there is a state error this method returns nil.
func (b Board) Row(row int) []int {
	// do nothing when any error occurred
	if b.e != nil {
		return nil
	}

	idx, err := b.index(row, 0)
	if err != nil {
		b.e = err
		return nil
	}

	return b.row(idx)
}

// Column method returns the entire column on a specific coordinate, which starts at 0 and the maximal value is 8.
// When there is a state error this method returns nil.
func (b Board) Column(column int) []int {
	// do nothing when any error occurred
	if b.e != nil {
		return nil
	}

	idx, err := b.index(0, column)
	if err != nil {
		b.e = err
		return nil
	}

	return b.column(idx)
}

// Board method returns the whole board values. When there is a state error this method returns nil.
func (b Board) Board() [][]int {
	// do nothing when any error occurred
	if b.e != nil {
		return nil
	}

	board := make([][]int, BoardSide, BoardSide)
	i := 0
	for r := 0; r < BoardSide; r++ {
		board[r] = make([]int, BoardSide, BoardSide)
		for c := 0; c < BoardSide; c++ {
			board[r][c] = int(b.b[i])
			i++
		}
	}

	return board
}

// Box method returns the box values based on the box index, which starts at 0 and the maximal value is 8.
// When there is a state error this method returns nil.
func (b *Board) Box(boxIndex int) []int {
	// do nothing when any error occurred
	if b.e != nil {
		return nil
	}

	if boxIndex < 0 || boxIndex >= BoardSide {
		b.e = errOutOfBoardIndex
		return nil
	}

	return b.box(boxIndex)
}

// Solve methods solve the Sudoku based on the set values. When there is a state error this method has no behavior.
func (b Board) Solve() {
	// do nothing when any error occurred
	if b.e != nil {
		return
	}
	b.s = false // set as not solved
	b.solve()
}

// String method provides the printable version of Sudoku board.
func (b Board) String() string {
	sb := strings.Builder{}

	for idx, i := 0, 1; idx < BoardSize; i, idx = i+1, idx+1 {
		sb.WriteString(fmt.Sprintf("|%d", b.b[idx]))
		if i == BoardSide {
			sb.WriteString("|\n")
			i = 0
		}
	}

	return sb.String()
}

// ------------------------------------------------- PRIVATE METHODS -------------------------------------------------

func (b Board) index(row, column int) (int, error) {
	if row < 0 || row >= BoardSide || column < 0 || column >= BoardSide {
		return -1, errOutOfBoardIndex
	}

	return row*BoardSide + column, nil
}

func (b Board) indexBox(boxIndex int) (int, error) {
	row := (boxIndex / BoardBoxSize) * BoardBoxSize
	column := (boxIndex % BoardBoxSize) * BoardBoxSize
	return b.index(row, column)
}

func (b Board) row(idx int) []int {
	r := make([]int, BoardSide, BoardSide)
	for i, j := idx, 0; i < idx+BoardSide; i, j = i+1, j+1 {
		r[j] = int(b.b[i])
	}

	return r
}

func (b Board) column(idx int) []int {
	c := make([]int, BoardSide, BoardSide)
	for i := 0; i < BoardSide; i++ {
		c[i] = int(b.b[i*BoardSide+idx])
	}

	return c
}

func (b Board) box(boxIndex int) []int {
	idx, _ := b.indexBox(boxIndex)

	box := make([]int, BoardSide, BoardSide)
	i := 0
	for r := 0; r < BoardBoxSize; r++ {
		for c := 0; c < BoardBoxSize; c++ {
			box[i] = int(b.b[idx+c])
			i++
		}
		idx += BoardSide
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
	// start searching from the end
	for i := len(b.b) - 1; i >= 0; i-- {
		if b.b[i] < 1 {
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
			b.b[idx] = 0 // empty value
			return
		}

		if b.IsValid() {
			b.solve()
		}
	}
}

// limited functionality to Sudoku where values could be within a limit
func (b Board) isValidSlice(values []int) bool {
	if len(values) != BoardSide {
		return false
	}

	m := make([]int, MaxValue+1, MaxValue+1)
	for _, v := range values {
		if v < 0 || v > MaxValue || m[v] > 0 {
			return false
		}
		m[v] = v
	}

	return true
}
