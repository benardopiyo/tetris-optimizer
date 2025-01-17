package main_test

import (
	"testing"

	"tetris-optimizer/tetrominosolver"
)

func TestIsValid(t *testing.T) {
	t.Run("Should return false when tetromino is not valid", func(t *testing.T) {
		tetromino := []string{"....", "....", "....", "...."}
		tetrominos := [][]string{tetromino}
		result := tetrominosolver.IsValid(tetrominos)
		if result != false {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("Should return true when tetromino is valid", func(t *testing.T) {
		tetromino := []string{"#...", "##..", "#...", "...."}
		tetrominos := [][]string{tetromino}
		result := tetrominosolver.IsValid(tetrominos)
		if result != true {
			t.Errorf("Expected true, got %v", result)
		}
	})
}

func TestTrimUnusedLines(t *testing.T) {
	t.Run("Should return tetromino without unused lines", func(t *testing.T) {
		tetromino := []string{"#...", "##..", "#...", "...."}
		tetrominos := [][]string{tetromino}
		result := tetrominosolver.TrimUnusedLines(tetrominos)
		if len(result[0]) != 3 {
			t.Errorf("Expected 3, got %v", len(result[0]))
		}
	})
}
