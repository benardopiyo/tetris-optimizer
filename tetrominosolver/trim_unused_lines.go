package tetrominosolver

// Remove empty rows and columns from Tetromino shapes
func TrimUnusedLines(tetrominos [][]string) [][]string {
	trimmedTetrominos := [][]string{}

	// Remove empty rows from each Tetromino
	for _, tetromino := range tetrominos {
		var trimmedTetromino []string
		for _, row := range tetromino {
			containsBlock := false
			for _, char := range row {
				if char == '#' {
					containsBlock = true
					break
				}
			}
			if containsBlock {
				trimmedTetromino = append(trimmedTetromino, row)
			}
		}
		trimmedTetrominos = append(trimmedTetrominos, trimmedTetromino)
	}

	// Remove empty columns from each Tetromino
	for i := range trimmedTetrominos {
		for col := len(trimmedTetrominos[i][0]) - 1; col >= 0; col-- {
			containsBlock := false
			for row := range trimmedTetrominos[i] {
				if trimmedTetrominos[i][row][col] == '#' {
					containsBlock = true
					break
				}
			}
			if !containsBlock {
				for row := range trimmedTetrominos[i] {
					trimmedTetrominos[i][row] = removeCharAt(trimmedTetrominos[i][row], col)
				}
			}
		}
	}

	return trimmedTetrominos
}

// Removes the character at a specified index in a string
func removeCharAt(input string, index int) string {
	if index < 0 || index >= len(input) {
		return input
	}

	runes := []rune(input)
	return string(append(runes[:index], runes[index+1:]...))
}
