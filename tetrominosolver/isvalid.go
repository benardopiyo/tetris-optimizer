package tetrominosolver

// Checks if a list of Tetromino shapes is valid.
// Each Tetromino is represented by a 2D slice of strings.
// The function returns true if all Tetrominos are valid, otherwise false.
func IsValid(tetrominos [][]string) bool {
	for _, tetromino := range tetrominos {
		connections := 0 // Count of valid connections between '#' characters
		hashCount := 0   // Count of '#' characters in the Tetromino

		for rowIndex, row := range tetromino {
			for colIndex, char := range row {
				adjacentHashes := 0 // Count of connections for the current '#' character

				if char != '#' && char != '.' {
					return false
				} else if char == '#' {
					hashCount++

					// Check for adjacent '#' characters in all four directions
					if rowIndex > 0 && tetromino[rowIndex-1][colIndex] == '#' {
						adjacentHashes++
					}
					if rowIndex < len(tetromino)-1 && tetromino[rowIndex+1][colIndex] == '#' {
						adjacentHashes++
					}
					if colIndex > 0 && tetromino[rowIndex][colIndex-1] == '#' {
						adjacentHashes++
					}
					if colIndex < len(row)-1 && tetromino[rowIndex][colIndex+1] == '#' {
						adjacentHashes++
					}

					// If no connections found for a '#', it's an isolated block
					if adjacentHashes == 0 {
						return false
					} else {
						connections += adjacentHashes // Add connections to total count
					}
				}
			}
		}

		// Check if the Tetromino has valid connection count and number of '#' characters
		if connections < 6 || hashCount > 4 {
			return false
		}
	}

	return true // All Tetrominos are valid
}
