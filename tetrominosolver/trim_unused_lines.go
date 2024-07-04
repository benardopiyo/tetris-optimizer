package tetrominosolver

// Remove empty rows and columns from the Tetromino shapes; parts that do not contain '#' chars.
func TrimUnusedLines(tetrominos [][]string) [][]string {
	trimmedTetrominos := [][]string{} // Hold the Tetrominos after removing unused lines.

	// Trim unused lines horizontally; remove rows that do not contain a '#' character.
	for _, tetromino := range tetrominos {
		var trimmedTetromino []string // Temporarily holds the rows of the current Tetromino that contain at least one '#'.
		for _, row := range tetromino {
			rowHashCount := 0
			for _, char := range row {
				if char == '#' {
					rowHashCount++
				}
			}

			if rowHashCount > 0 {
				trimmedTetromino = append(trimmedTetromino, row)
			}
		}
		trimmedTetrominos = append(trimmedTetrominos, trimmedTetromino)
	}

	// Trim unused lines vertically; remove any columns that do not contain a '#' character.
	for i := range trimmedTetrominos {
		for j := len(trimmedTetrominos[i][0]) - 1; j >= 0; j-- {
			columnHashCount := 0
			for k := range trimmedTetrominos[i] {
				if trimmedTetrominos[i][k][j] == '#' {
					columnHashCount++
				}
			}
			// If the current column does not contain any '#', remove it from all rows.
			if columnHashCount == 0 {
				for colIndex := range trimmedTetrominos[i] {
					trimmedTetrominos[i][colIndex] = removeCharAtIndex(trimmedTetrominos[i][colIndex], j)
				}
			}
		}
	}

	return trimmedTetrominos
}

// Removes the character at a specific index in a string; remove a column from each row of the Tetromino.
func removeCharAtIndex(input string, index int) string {
	if index < 0 || index >= len(input) {
		return input // return the original string if index is out of bounds

	}

	runes := []rune(input)
	runes = append(runes[:index], runes[index+1:]...)
	return string(runes)
}
