package main

import (
	"fmt"
	"math"
	"os"

	"tetris-optimizer/tetrominosolver"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <input_file>")
		return
	}

	inputFile := os.Args[1]

	// Read Tetromino shapes from the specified input file
	tetrominos, err := tetrominosolver.ReadInputFile(inputFile)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	// Validate the read Tetromino shapes
	if !tetrominosolver.IsValid(tetrominos) || len(tetrominos) == 0 {
		fmt.Println("ERROR")
		return
	}

	// Calculate the initial board size based on the number of Tetrominos
	boardSize := int(math.Ceil(math.Sqrt(float64(len(tetrominos) * 4))))
	board := make([][]string, boardSize)

	// Initialize the board with empty cells represented by "."
	for i := range board {
		board[i] = make([]string, boardSize)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	trimmedTetrominos := tetrominosolver.TrimUnusedLines(tetrominos) // Trim unnecessary rows and columns

	solvedBoard := tetrominosolver.Resolve(trimmedTetrominos, board) // Solve the board by placing tetrominos

	for _, row := range solvedBoard {
		fmt.Println(row)
	}
}
