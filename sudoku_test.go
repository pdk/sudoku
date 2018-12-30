package sudoku_test

import (
	"testing"

	"github.com/pdk/sudoku"
)

func TestBlock(t *testing.T) {
	b := sudoku.Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 2, 3, 0, 0, 0},
		{0, 0, 0, 4, 5, 6, 0, 0, 0},
		{0, 0, 0, 7, 8, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 8, 7},
		{0, 0, 0, 0, 0, 0, 6, 5, 4},
		{0, 0, 0, 0, 0, 0, 3, 2, 1},
	}

	ZeroBlock := sudoku.Segment{
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	}

	OneBlock := sudoku.Segment{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	NineBlock := sudoku.Segment{
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
	}

	s := b.Block(0, 0)

	if s != ZeroBlock {
		t.Errorf("did not get a zero block, got: %s", s)
	}

	s = b.Block(3, 3)

	if s != OneBlock {
		t.Errorf("did not get a one block, got: %s", s)
	}

	s = b.Block(4, 5)

	if s != OneBlock {
		t.Errorf("did not get a one block, got: %s", s)
	}

	s = b.Block(3, 3)

	if s != OneBlock {
		t.Errorf("did not get a one block, got: %s", s)
	}

	s = b.Block(6, 6)

	if s != NineBlock {
		t.Errorf("did not get a nine block, got: %s", s)
	}

	s = b.Block(8, 8)

	if s != NineBlock {
		t.Errorf("did not get a nine block, got: %s", s)
	}
}

func TestColumn(t *testing.T) {

	puzzle := sudoku.Board{
		{0, 2, 6, 0, 0, 0, 8, 1, 0},
		{3, 0, 0, 7, 0, 8, 0, 0, 6},
		{4, 0, 0, 0, 5, 0, 0, 0, 7},
		{0, 5, 0, 1, 0, 7, 0, 9, 0},
		{0, 0, 3, 9, 0, 5, 1, 0, 0},
		{0, 4, 0, 3, 0, 2, 0, 5, 0},
		{1, 0, 0, 0, 3, 0, 0, 0, 2},
		{5, 0, 0, 2, 0, 4, 0, 0, 9},
		{0, 3, 8, 0, 0, 0, 4, 6, 0},
	}

	solved := sudoku.Board{
		{7, 2, 6, 4, 9, 3, 8, 1, 5},
		{3, 1, 5, 7, 2, 8, 9, 4, 6},
		{4, 8, 9, 6, 5, 1, 2, 3, 7},
		{8, 5, 2, 1, 4, 7, 6, 9, 3},
		{6, 7, 3, 9, 8, 5, 1, 2, 4},
		{9, 4, 1, 3, 6, 2, 7, 5, 8},
		{1, 9, 4, 8, 3, 6, 5, 7, 2},
		{5, 6, 7, 2, 1, 4, 3, 8, 9},
		{2, 3, 8, 5, 7, 9, 4, 6, 1},
	}

	c0 := puzzle.Column(0)

	if (c0 != sudoku.Segment{0, 3, 4, 0, 0, 0, 1, 5, 0}) {
		t.Errorf("c0 not as expected: %s", c0)
	}

	if c0.IsValid() {
		t.Errorf("c0 should not be valid, but is: %s", c0)
	}

	if c0.IsComplete() {
		t.Errorf("c0 should not be complete, but is: %s", c0)
	}

	c1 := solved.Column(0)

	if (c1 != sudoku.Segment{7, 3, 4, 8, 6, 9, 1, 5, 2}) {

		t.Errorf("c1 not as expected: %s", c1)
	}

	if !c1.IsValid() {
		t.Errorf("c1 should be valid, but isn't: %s", c1)
	}

	if !c1.IsComplete() {
		t.Errorf("c1 should be complete, but isn't: %s", c1)
	}
}

func TestRow(t *testing.T) {
	b := sudoku.Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 2, 3, 0, 0, 0},
		{0, 0, 0, 4, 5, 6, 0, 0, 0},
		{0, 0, 0, 7, 8, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 8, 7},
		{0, 0, 0, 0, 0, 0, 6, 5, 4},
		{0, 0, 0, 0, 0, 0, 3, 2, 1},
	}

	ZeroRow := sudoku.Segment{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	OneRow := sudoku.Segment{
		0, 0, 0, 1, 2, 3, 0, 0, 0,
	}
	LastRow := sudoku.Segment{
		0, 0, 0, 0, 0, 0, 3, 2, 1,
	}

	s := b.Row(0)

	if s != ZeroRow {
		t.Errorf("didn't get ZeroRow, got %s", s)
	}

	s = b.Row(3)

	if s != OneRow {
		t.Errorf("didn't get OneRow, got %s", s)
	}

	s = b.Row(8)

	if s != LastRow {
		t.Errorf("didn't get LastRow, got %s", s)
	}
}

func TestOptions(t *testing.T) {
	b := sudoku.Board{
		{1, 2, 3, 4, 5, 6, 7, 8},
	}

	b, changed := b.CompleteSingleOptions()

	if !changed {
		t.Errorf("should have changed, but did not")
	}

	newRow := sudoku.Segment{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	if b.Row(0) != newRow {
		t.Errorf("should have got new row, but got %s", b.Row(0))
	}
}

func TestBoard(t *testing.T) {

	puzzle := sudoku.Board{
		{0, 2, 6, 0, 0, 0, 8, 1, 0},
		{3, 0, 0, 7, 0, 8, 0, 0, 6},
		{4, 0, 0, 0, 5, 0, 0, 0, 7},
		{0, 5, 0, 1, 0, 7, 0, 9, 0},
		{0, 0, 3, 9, 0, 5, 1, 0, 0},
		{0, 4, 0, 3, 0, 2, 0, 5, 0},
		{1, 0, 0, 0, 3, 0, 0, 0, 2},
		{5, 0, 0, 2, 0, 4, 0, 0, 9},
		{0, 3, 8, 0, 0, 0, 4, 6, 0},
	}

	solved := sudoku.Board{
		{7, 2, 6, 4, 9, 3, 8, 1, 5},
		{3, 1, 5, 7, 2, 8, 9, 4, 6},
		{4, 8, 9, 6, 5, 1, 2, 3, 7},
		{8, 5, 2, 1, 4, 7, 6, 9, 3},
		{6, 7, 3, 9, 8, 5, 1, 2, 4},
		{9, 4, 1, 3, 6, 2, 7, 5, 8},
		{1, 9, 4, 8, 3, 6, 5, 7, 2},
		{5, 6, 7, 2, 1, 4, 3, 8, 9},
		{2, 3, 8, 5, 7, 9, 4, 6, 1},
	}

	isSolved := puzzle.IsSolved()

	if isSolved {
		t.Errorf("IsSolved returned true, but should not have")
	}

	isSolved = solved.IsSolved()

	if !isSolved {
		t.Errorf("IsSolved returned false, but should not have")
	}

	solution, changed := puzzle.CompleteSingleOptions()

	if !changed {
		t.Errorf("should have changed, but didn't")
	}

	if puzzle == solution {
		t.Errorf("should have gotten back a changed board, but got same")
	}

	if solved != solution {
		t.Errorf("should have gotten back solved, but got %s", solution)
	}

}

func TestMultiSolution(t *testing.T) {

	puzzle := sudoku.Board{
		{0, 8, 0, 0, 0, 9, 7, 4, 3},
		{0, 5, 0, 0, 0, 8, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 8, 0, 4, 0, 0, 0},
		{0, 0, 0, 3, 0, 0, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 0, 7, 0},
		{0, 3, 0, 5, 0, 0, 0, 8, 0},
		{9, 7, 2, 4, 0, 0, 0, 5, 0},
	}

	partialSolved := sudoku.Board{
		{2, 8, 6, 1, 5, 9, 7, 4, 3},
		{0, 5, 0, 0, 0, 8, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 8, 0, 4, 0, 0, 0},
		{0, 0, 0, 3, 0, 0, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 0, 7, 0},
		{0, 3, 0, 5, 0, 0, 0, 8, 0},
		{9, 7, 2, 4, 0, 0, 0, 5, 1},
	}

	solution, changed := puzzle.CompleteSingleOptions()

	if !changed {
		t.Errorf("should have changed, but didn't")
	}

	if solution == puzzle {
		t.Errorf("solution should be diff than puzzle, got: %s", solution)
	}

	if solution != partialSolved {
		t.Errorf("solution should be same as partialSolved, but got: %s", solution)
	}

}
