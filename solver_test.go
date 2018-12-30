package sudoku_test

import (
	"testing"

	"github.com/pdk/sudoku"
)

func TestSolver(t *testing.T) {

	// This has 2 solutions:
	puzzle := sudoku.Board{
		{9, 2, 6, 5, 7, 1, 4, 8, 3},
		{3, 5, 1, 4, 8, 6, 2, 7, 9},
		{8, 7, 4, 9, 2, 3, 5, 1, 6},
		{5, 8, 2, 3, 6, 7, 1, 9, 4},
		{1, 4, 9, 2, 5, 8, 3, 6, 7},
		{7, 6, 3, 1, 0, 0, 8, 2, 5},
		{2, 3, 8, 7, 0, 0, 6, 5, 1},
		{6, 1, 7, 8, 3, 5, 9, 4, 2},
		{4, 9, 5, 6, 1, 2, 7, 3, 8},
	}

	sudoku.Solve(puzzle)
}