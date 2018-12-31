package sudoku

// Value is what can be in a single spot on the board
type Value int

// Valid Value choices
const (
	Blank Value = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

// Board is a 9x9 grid of Values
type Board [9][9]Value

// Segment is a subgroup. Either row, column or block.
type Segment [9]Value

// Mark places a value on a row, column location in a Board.
func Mark(b Board, i, j int, v Value) Board {
	b[i][j] = v
	return b
}

// Row returns the Segment of a particular row.
func (b Board) Row(row int) Segment {
	s := Segment{}
	for i := 0; i < 9; i++ {
		s[i] = b[row][i]
	}
	return s
}

// Column returns the Segment of a particular column.
func (b Board) Column(col int) Segment {
	s := Segment{}
	for i := 0; i < 9; i++ {
		s[i] = b[i][col]
	}
	return s
}

// Block returns the Segment of a particular Block given row and column position
// in the overal grid.
func (b Board) Block(row, col int) Segment {
	s := Segment{}

	blockRow := row / 3
	blockCol := col / 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			tPos := i*3 + j

			s[tPos] = b[blockRow*3+i][blockCol*3+j]
		}
	}

	return s
}

// IsComplete is true IFF there are no blanks in the Segment
func (s Segment) IsComplete() bool {
	for _, v := range s {
		if v == Blank {
			return false
		}
	}

	return true
}

// IsValid is true IFF no non-Blank Values are repeated
func (s Segment) IsValid() bool {
	seen := [10]bool{}

	for _, v := range s {
		if seen[v] {
			return false
		}
		seen[v] = true
	}

	return true
}

// IsComplete is true IFF the entire board has non-Blank values.
func (b Board) IsComplete() bool {

	_, _, gotBlank := b.FirstBlank()
	return !gotBlank
}

// FirstBlank returns the row,col position of the first blank on the board.
func (b Board) FirstBlank() (int, int, bool) {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == Blank {
				return i, j, true
			}
		}
	}

	return 0, 0, false
}

// IsValid is true IFF no row, column or block of the board contains duplicate
// values.
func (b Board) IsValid() bool {

	for i := 0; i < 9; i++ {
		if !b.Row(i).IsValid() {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if !b.Column(j).IsValid() {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !b.Block(i, j).IsValid() {
				return false
			}
		}
	}

	return true
}

// Options returns a slice of Values that could be placed at position row, col
// without invalidating the board.
func (b Board) Options(row, col int) []Value {
	counts := [10]int{}

	for _, v := range b.Row(row) {
		counts[v]++
	}
	for _, v := range b.Column(col) {
		counts[v]++
	}
	for _, v := range b.Block(row, col) {
		counts[v]++
	}

	vals := []Value{}

	for i := 1; i < 10; i++ {
		if counts[i] == 0 {
			vals = append(vals, Value(i))
		}
	}

	return vals
}

// CompleteSingleOptions will find all locations that have only a single option,
// and place that value. Repeats until there are no more single-option locations
// left, then returns. Returns the new Board, and true/false to indicate if any
// values were placed on the board.
func (b Board) CompleteSingleOptions() (Board, bool) {

	updatedBoard := false

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if b[i][j] != Blank {
				continue
			}

			options := b.Options(i, j)
			if len(options) != 1 {
				continue
			}

			// found one single-option position
			b[i][j] = options[0]
			updatedBoard = true
			// start looping again from the beginning
			i = 0
			j = -1
		}
	}

	return b, updatedBoard
}

// IsSolved returns true IFF the board is complete and valid.
func (b Board) IsSolved() bool {
	return b.IsComplete() && b.IsValid()
}
