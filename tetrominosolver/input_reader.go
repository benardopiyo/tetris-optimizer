package tetrominosolver

import (
	"bufio"
	"fmt"
	"os"
)

// ReadInputFile reads the Tetromino shapes from a file
func ReadInputFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tetrominos [][]string
	var currentTetromino []string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 4 {
			currentTetromino = append(currentTetromino, line)
		} else if line != "" {
			return nil, fmt.Errorf("ERROR")
		}

		if line == "" {
			if len(currentTetromino) > 0 {
				tetrominos = append(tetrominos, currentTetromino)
				currentTetromino = []string{}
			}
		}
	}
	if len(currentTetromino) > 0 {
		tetrominos = append(tetrominos, currentTetromino)
	}

	return tetrominos, nil
}
