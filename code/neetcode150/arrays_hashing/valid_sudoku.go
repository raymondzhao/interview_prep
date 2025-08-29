package arrayshashing

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to
the following rules:

    Each row must contain the digits 1-9 without repetition.
    Each column must contain the digits 1-9 without repetition.
    Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:

    A Sudoku board (partially filled) could be valid but is not necessarily solvable.
    Only the filled cells need to be validated according to the mentioned rules.
*/

func isValidSudoku(board [][]byte) bool {
	// go through each number and check the row, column, box
	// slices/lists of maps to check that row/column/box
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// instantiate each map (go things)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// iterate through the board checking each place
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			curr := board[r][c]

			// ignore the . bc those are empty
			if curr == '.' {
				continue
			}

			// to figure out which box, we can do the below
			boxIndex := (r/3)*3 + (c / 3)

			// check each map to see if there are any duplicates
			if rows[r][curr] || columns[c][curr] || boxes[boxIndex][curr] {
				return false
			}

			// we have now seen the value so update all the maps accordingly
			rows[r][curr] = true
			columns[c][curr] = true
			boxes[boxIndex][curr] = true
		}
	}

	return true
}
