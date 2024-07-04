package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"tetris-optimizer/tetrominosolver"
)

// Read Tetromino shapes from the specified input file.
// Validate the Tetromino shapes
// Calculate the initial size of the board required to fit all Tetrominos.
// Trim any unnecessary rows and columns from the Tetrominos.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	firstArg := os.Args[1]

	Tetrominos, err := tetrominosolver.ReadInputFile(firstArg)
	if err != nil {
		errMsg := err.Error()
		fmt.Println(errMsg)
		return
	}

	if !tetrominosolver.IsValid(Tetrominos) || len(Tetrominos) == 0 {
		fmt.Println("ERROR")
		return
	}

	boardSize := int(math.Ceil(math.Sqrt(float64(len(Tetrominos) * 4))))

	// Create a new board of the calculated size and initialize it with empty cells represented by ".".
	board := make([][]string, boardSize)
	for i := range board {
		board[i] = make([]string, boardSize)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	trimmedTetrominos := tetrominosolver.TrimUnusedLines(Tetrominos)

	resolvedBoard := tetrominosolver.Resolve(trimmedTetrominos, board)
	fmt.Println("Time Taken: ", time.Since(time.Now()))

	for _, row := range resolvedBoard {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}
