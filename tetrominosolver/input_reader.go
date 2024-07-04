package tetrominosolver

import (
	"bufio"
	"fmt"
	"os"
)

// ReadInputFile reads a file containing Tetromino shapes and returns them as a slice of 2D string slices.
// The function will read the file, validate the format, and parse the Tetromino shapes into a slice of 2D string slices.

func ReadInputFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize a slice to hold the Tetromino grids.
	Tetromino := [][]string{}
	tetrominoLines := []string{}
	for scanner.Scan() {
		currentLine := scanner.Text()
		if len(currentLine) == 4 {
			tetrominoLines = append(tetrominoLines, currentLine)
		} else if currentLine != "" && len(currentLine) != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		if currentLine == "" { // Indicate the end of a Tetromino grid
			Tetromino = append(Tetromino, tetrominoLines)
			tetrominoLines = []string{} // Reset the temporary slice for the next Tetromino grid.
		} else if currentLine != "" && len(tetrominoLines) > 4 {
			return nil, fmt.Errorf("ERROR")
		}
	}

	if len(tetrominoLines) > 0 {
		Tetromino = append(Tetromino, tetrominoLines)
	}

	return Tetromino, nil
}
