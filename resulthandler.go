package sudoku

import (
	"fmt"
	"strconv"
)

// Print prints out a Board
func (b Board) Print() {
	for i := 0; i < 9; i++ {
		fmt.Printf("%s%s%s %s%s%s %s%s%s\n",
			b[i][0].String(),
			b[i][1].String(),
			b[i][2].String(),
			b[i][3].String(),
			b[i][4].String(),
			b[i][5].String(),
			b[i][6].String(),
			b[i][7].String(),
			b[i][8].String(),
		)
		if i == 2 || i == 5 {
			fmt.Printf("\n")
		}
	}
}

func (v Value) String() string {
	if v == 0 {
		return "·"
	}

	return strconv.Itoa(int(v))
}

// SolutionHandler receives and prints found solutions.
type SolutionHandler struct {
	solutions chan Board
	wait      chan struct{}
	seen      map[Board]bool
}

// NewSolutionHandler creates and returns a new SolutionHandler
func NewSolutionHandler() SolutionHandler {
	h := SolutionHandler{
		solutions: make(chan Board),
		wait:      make(chan struct{}),
		seen:      make(map[Board]bool),
	}

	go func() {
		c := 0
		for solved := range h.solutions {

			if h.seen[solved] {
				continue
			}
			h.seen[solved] = true

			c++
			fmt.Printf("\nSolution %d:\n\n", c)
			solved.Print()
		}

		// signal that we're done processing
		close(h.wait)
	}()

	return h
}

// Submit hands off a solution to be printed.
func (h SolutionHandler) Submit(solution Board) {
	h.solutions <- solution
}

// Flush should be called after all solutions have been found & submitted.
func (h SolutionHandler) Flush() {
	close(h.solutions)

	// wait until we've finished printing all solutions
	<-h.wait
}
