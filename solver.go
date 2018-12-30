package sudoku

import (
	"sync"
)

// Solve will find any/all possible solutions and print them out.
func Solve(b Board) {
	handler := NewSolutionHandler()

	findSolutions(b, handler)

	handler.Flush()
}

func findSolutions(b Board, handler SolutionHandler) {

	b, _ = b.CompleteSingleOptions()

	if b.IsSolved() {
		handler.Submit(b)
		return
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] != Blank {
				continue
			}

			options := b.Options(i, j)
			if len(options) == 0 {
				// if any blank on the board has no options, there is no
				// solution down this path.
				return
			}

			var wg sync.WaitGroup

			for _, opt := range options {
				b[i][j] = opt
				wg.Add(1)
				go func(b Board) {
					findSolutions(b, handler)
					wg.Done()
				}(b)
			}

			wg.Wait()
		}
	}
}
