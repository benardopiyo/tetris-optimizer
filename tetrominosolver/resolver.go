package tetrominosolver

// Position of a Tetromino on the board with X and Y coordinates.
type pos struct {
	X int
	Y int
}

// Holds the state of the board and the positions of the Tetrominos
type save struct {
	board [][]string
	pos   [][]pos
}

// Creates a deep copy of the given board to ensure that modifications do not affect the original board.
func createBoardCopy(board [][]string) [][]string {
	newBoard := [][]string{}

	for _, row := range board {
		newRow := []string{}

		newRow = append(newRow, row...)

		newBoard = append(newBoard, newRow)
	}

	return newBoard
}

// Attempts to place all Tetrominos on the board and finds a valid configuration for the given Tetrominos.
func Resolve(tetrominos [][]string, board [][]string) [][]string {
	boardSaves := []save{} // Stack to keep track of different board states and Tetromino positions

	for i := 0; i < len(tetrominos); i++ {
		successfullyPlaced := false
		isLastPossiblePos := true
		duplicateFound := false
		possibleToPlaceCurrentTetromino := false
		Positions := [][]pos{}

		// Try placing the Tetromino in every position on the board.
		for rowIndex := 0; rowIndex < len(board[0]) && !successfullyPlaced; rowIndex++ {
			for colIndex := 0; colIndex < len(board) && !successfullyPlaced; colIndex++ {
				if CanPlaceTetromino(tetrominos[i], board, rowIndex, colIndex) && !successfullyPlaced {
					possibleToPlaceCurrentTetromino = true
					posistionTriedBefore := false // Flag to indicate if the position has been tried before

					// Check for duplicate position by comparing with previous saves.
					if len(boardSaves) > 0 && len(boardSaves) == i+1 {
						for _, pos := range boardSaves[len(boardSaves)-1].pos[len(boardSaves[len(boardSaves)-1].pos)-1] {
							if pos.X == rowIndex && pos.Y == colIndex {
								posistionTriedBefore = true // The position is a duplicate
								duplicateFound = true
							}
						}
					}

					if len(boardSaves) > 0 { // Retrieve the previous Tetromino positions if there are any saves.
						Positions = boardSaves[len(boardSaves)-1].pos
					}

					if posistionTriedBefore { // Skip if the position is a duplicate.
						continue
					}

					if !duplicateFound {
						Positions = append(Positions, []pos{}) // Add the current position to the list of positions.
					}

					successfullyPlaced = true
					localBoard := createBoardCopy(board)
					Positions[len(Positions)-1] = append(Positions[len(Positions)-1], pos{rowIndex, colIndex})

					if !duplicateFound {
						isLastPossiblePos = false
						boardSaves = append(boardSaves, save{localBoard, Positions}) // Save the current state of the board and Tetromino positions.
					}

					board = PlaceTetromino(tetrominos[i], board, rowIndex, colIndex, i)
				}
			}
		}

		if !possibleToPlaceCurrentTetromino {
			isLastPossiblePos = false // No found for the current Tetromino, reset and try a new configuration.
		}

		if !successfullyPlaced {
			i-- // Go back in next cycle to the current Tetromino

			// If there are no more saves and we are out of space, reset the board.
			if len(boardSaves) <= 1 {
				i = -1 // Go back in the next cycle to the first Tetromino

				boardSize := len(board) + 1
				board = make([][]string, boardSize)
				for i := range board {
					board[i] = make([]string, boardSize)
					for j := range board[i] {
						board[i][j] = "."
					}
				}

				boardSaves = []save{} // Clear the save stack as we are starting a new configuration.
			}

			// If there are previous saves, revert to the last saved state.
			if len(boardSaves) > 1 {
				i-- // Go back in the next cycle to the previous Tetromino

				if isLastPossiblePos {
					{
						boardSaves = boardSaves[:len(boardSaves)-1]                  // Remove the last save
						board = createBoardCopy(boardSaves[len(boardSaves)-1].board) // Restore the board from the last save
					}
				} else {
					board = createBoardCopy(boardSaves[len(boardSaves)-1].board)
				}
			}
		}
	}

	return board
}

// Checks if a Tetromino: can be placed on the board at a specific position;
// Tetromino is out of bounds & if the position on the board is empty.
// Returns true if a Tetromino can be placed at the specified position.
func CanPlaceTetromino(tetromino []string, board [][]string, indexVertical int, indexHorizontal int) bool {
	for indexVerticalTetromino, line := range tetromino {
		for indexHorizontalTetromino, char := range line {
			if char == '#' {
				if indexVertical+indexVerticalTetromino > len(board)-1 || indexHorizontal+indexHorizontalTetromino > len(board)-1 {
					return false
				}
				if board[indexVertical+indexVerticalTetromino][indexHorizontal+indexHorizontalTetromino] != "." {
					return false
				}
			}
		}
	}
	return true
}

// Places a Tetromino on the board at a specific position and marks it with chars (A ... Z)
func PlaceTetromino(tetromino []string, board [][]string, colIndex int, rowIndex int, tetrominoIndex int) [][]string {
	for colIndexTetromino, row := range tetromino {
		for rowIndexTetromino, char := range row {
			if char == '#' {

				startLetter := 'A' + tetrominoIndex // Adds the Tetromino index to get the next letter.
				if startLetter > 'Z' {
					board[colIndex+colIndexTetromino][rowIndex+rowIndexTetromino] = "X"
				} else {
					board[colIndex+colIndexTetromino][rowIndex+rowIndexTetromino] = string(rune(startLetter)) // Use 'A' to 'Z'.
				}
			}
		}
	}
	return board
}
