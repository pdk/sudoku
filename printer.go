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
