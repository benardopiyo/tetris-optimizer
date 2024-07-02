package tetrominosolver

// Pos represents the coordinates on the board
type Pos struct {
	X, Y int
}

// Save stores the state of the board and the positions of tetrominos
type Save struct {
	Board [][]string
	Pos   []Pos
}

// GetBoard creates a deep copy of the board
func GetBoard(board [][]string) [][]string {
	newBoard := make([][]string, len(board))
	for i := range board {
		newBoard[i] = append([]string(nil), board[i]...)
	}
	return newBoard
}

// CanPlaceTetromino checks if a tetromino can be placed at the specified position on the board
func CanPlaceTetromino(tetromino []string, board [][]string, row, col int) bool {
	for rowIdx, line := range tetromino {
		for colIdx, char := range line {
			if char == '#' {
				if row+rowIdx >= len(board) || col+colIdx >= len(board[0]) || board[row+rowIdx][col+colIdx] != "." {
					return false
				}
			}
		}
	}
	return true
}

// PlaceTetromino places a tetromino on the board at the specified position
func PlaceTetromino(tetromino []string, board [][]string, row, col, index int) [][]string {
	for rowIdx, line := range tetromino {
		for colIdx, char := range line {
			if char == '#' {
				board[row+rowIdx][col+colIdx] = string(rune('A' + index))
			}
		}
	}
	return board
}

// Backtrack resets the board to the previous state
func Backtrack(boardSaves []Save, board [][]string) ([]Save, [][]string) {
	if len(boardSaves) > 0 {
		boardSaves = boardSaves[:len(boardSaves)-1]
		if len(boardSaves) > 0 {
			board = GetBoard(boardSaves[len(boardSaves)-1].Board)
		} else {
			board = CreateNewBoard(len(board) + 1)
		}
	}
	return boardSaves, board
}

// CreateNewBoard initializes a new board with given size
func CreateNewBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// Resolve tries to place tetrominos on the board
func Resolve(tetrominos [][]string, board [][]string) [][]string {
	var boardSaves []Save

	for i := 0; i < len(tetrominos); i++ {
		placed := false

		for row := 0; row < len(board) && !placed; row++ {
			for col := 0; col < len(board[0]) && !placed; col++ {
				if CanPlaceTetromino(tetrominos[i], board, row, col) {
					localBoard := GetBoard(board)
					board = PlaceTetromino(tetrominos[i], board, row, col, i)
					var positions []Pos
					if len(boardSaves) > 0 {
						positions = append(positions, boardSaves[len(boardSaves)-1].Pos...)
					}
					positions = append(positions, Pos{row, col})
					boardSaves = append(boardSaves, Save{localBoard, positions})
					placed = true
				}
			}
		}

		if !placed {
			i-- // Go back to previous tetromino
			boardSaves, board = Backtrack(boardSaves, board)
			if len(boardSaves) == 0 {
				i = -1 // Restart placement process with a larger board
			}
		}
	}

	return board
}
