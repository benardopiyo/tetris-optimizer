package tetrominosolver

// Checks if the given Tetrominos are valid according to the game’s rules.
// Rules:
// Contains `#` and `.` only.
// Each `#` char must be connected to at least one other `#` char (horizontally or vertically; not diagonally).
// Tetromino must have exactly 4 `#` chars.
// Total number of connections between `#` chars must be at least 6.
// The function returns `true` if all Tetrominos are valid, and `false` otherwise.

func IsValid(Tetrominos [][]string) bool {
	for _, Tetromino := range Tetrominos {
		ConnectionCount := 0
		HashCount := 0

		// Iterate over each line (row) of the Tetromino.
		for rowIndex, line := range Tetromino {
			for colIndex, char := range line {
				CurrentHashConnectionCount := 0

				if char != '#' && char != '.' {
					return false
				} else if char == '#' {
					HashCount++

					// Check if the `#` has a neighboring `#`
					if rowIndex > 0 && Tetromino[rowIndex-1][colIndex] == '#' { // above it.
						CurrentHashConnectionCount++
					}

					if rowIndex < len(Tetromino)-1 && Tetromino[rowIndex+1][colIndex] == '#' { // below it.
						CurrentHashConnectionCount++
					}

					if colIndex > 0 && Tetromino[rowIndex][colIndex-1] == '#' { // to the left of it.
						CurrentHashConnectionCount++
					}

					if colIndex < len(line)-1 && Tetromino[rowIndex][colIndex+1] == '#' { // to the right of it.
						CurrentHashConnectionCount++
					}

					if CurrentHashConnectionCount == 0 { // Invalid Tetromino if no connections for the '#'
						return false
					} else {
						ConnectionCount += CurrentHashConnectionCount // Add up the total number of connections
					}
				}
			}
		}

		if ConnectionCount < 6 || HashCount > 4 { // Check if there are exactly 4 `#` characters and at least 6 connections.
			return false
		}
	}

	return true
}
