package sudoku

import (
	"fmt"
	"log"
	"runtime"
)

var (
	maxGoroutines   int
	countGoroutines int
)

func setMaxGoroutines() {
	i := runtime.NumGoroutine()
	if i > maxGoroutines {
		maxGoroutines = i
	}
}

// PrintSolutions will run the solver and print the solutions. Also prints a
// count of search paths that did not produce a solution.
func PrintSolutions(b Board) {

	failCount := 0
	failCounter := func(b Board) {
		failCount++
	}

	solutionCount := 0
	solutionPrinter := func(b Board) {
		solutionCount++
		fmt.Printf("Solution %d:\n\n", solutionCount)
		b.Print()
		fmt.Println()
	}

	Solve(b, solutionPrinter, failCounter)

	fmt.Printf("there were %d failed search paths", failCount)
}

// ResultHandler is a thing that does something with a result.
type ResultHandler func(Board)

// Solve will find any/all possible solutions and print them out.
func Solve(b Board, solutionHandler, failureHandler ResultHandler) {

	done := make(chan interface{})
	extraResults := 0

	for result := range SearchSolutions(b, done) {
		if result.IsSolved() {
			solutionHandler(result)
			if done != nil {
				close(done)
				done = nil
			}
		} else {
			failureHandler(result)
		}
		if done == nil {
			extraResults++
		}
	}

	log.Printf("maxGoroutines = %d", maxGoroutines)
	log.Printf("countGoroutines = %d", countGoroutines)
	log.Printf("extraResults = %d", extraResults)
}

// SearchSolutions will search options trying to find solutions to the given
// board. Returns a channel that will produce all the search results. The caller
// must determine if the results are valid (a solution) or not (a failure).
func SearchSolutions(b Board, done <-chan interface{}) <-chan Board {

	setMaxGoroutines()

	select {
	case <-done:
		return NoValues()
	default:
		// proceed
	}

	b, _ = b.CompleteSingleOptions()

	i, j, foundBlank := b.FirstBlank()
	if !foundBlank {
		// no blanks: board is complete.
		return OneValue(b)
	}

	options := b.Options(i, j)
	if len(options) == 0 {
		// no options: there is no solution for this board.
		return OneValue(b)
	}

	var c <-chan Board
	for _, opt := range options {
		c = MergeValues(c, SearchOption(Mark(b, i, j, opt), done), done)
	}

	return c
}

// SearchOption explores one option in a new goroutine.
func SearchOption(b Board, done <-chan interface{}) <-chan Board {
	countGoroutines++

	c := make(chan Board)
	go func() {
		defer close(c)

		for result := range SearchSolutions(b, done) {
			c <- result
		}
	}()

	return c
}

// NoValues returns a channel that produces 0 values.
func NoValues() <-chan Board {

	c := make(chan Board)
	close(c)

	return c
}

// OneValue returns a channel that produces a single value.
func OneValue(b Board) <-chan Board {
	countGoroutines++

	c := make(chan Board)
	go func() {
		defer close(c)

		c <- b
	}()

	return c
}

// MergeValues coalesces values from two channels into a single channel.
func MergeValues(a <-chan Board, b <-chan Board, done <-chan interface{}) <-chan Board {
	countGoroutines++

	copying := true

	c := make(chan Board)
	go func() {
		defer close(c)

		for a != nil || b != nil {
			select {
			case <-done:
				copying = false
			case r, ok := <-a:
				if !ok {
					a = nil
					break
				}
				if copying {
					c <- r
				}
			case r, ok := <-b:
				if !ok {
					b = nil
					break
				}
				if copying {
					c <- r
				}
			}
		}
	}()

	return c
}
