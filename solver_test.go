package sudoku_test

import (
	"testing"

	"github.com/pdk/sudoku"
)

func TestSolver(t *testing.T) {

	// this has 2 solutions
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

	failCount := 0
	failCounter := func(b sudoku.Board) {
		failCount++
	}

	solutionCount := 0
	solutionCounter := func(b sudoku.Board) {
		solutionCount++
	}

	sudoku.Solve(puzzle, solutionCounter, failCounter)

	if solutionCount != 2 {
		t.Errorf("there should be two solutions, got %d", solutionCount)
	}

	if failCount != 0 {
		t.Errorf("there should not be failures, got %d", failCount)
	}
}

func TestSolver2(t *testing.T) {

	puzzle := sudoku.Board{
		{0, 2, 0, 0, 0, 5, 0, 0, 0},
		{0, 1, 5, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 8, 7, 0, 3},
		{0, 5, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 9, 7, 0, 0, 0, 1, 0},
		{0, 0, 0, 3, 0, 0, 0, 4, 6},
		{0, 0, 0, 0, 8, 0, 0, 0, 1},
		{7, 0, 0, 9, 3, 0, 0, 6, 0},
		{0, 0, 0, 0, 0, 0, 4, 0, 8},
	}

	failCount := 0
	failCounter := func(b sudoku.Board) {
		failCount++
	}

	solutionCount := 0
	solutionCounter := func(b sudoku.Board) {
		solutionCount++
	}

	sudoku.Solve(puzzle, solutionCounter, failCounter)

	if solutionCount != 1 {
		t.Errorf("there should be one solution, got %d", solutionCount)
	}

	if failCount != 201 {
		t.Errorf("there should be 201 failures, got %d", failCount)
	}
}
