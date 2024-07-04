# Tetris Optimizer

## Overview

* Tetris Optimizer program is designed to solve a classic Tetris-style puzzle. 
* Given a text file containing a list of tetrominoes, the program assembles these pieces into the smallest possible square grid, identifying each tetromino with unique uppercase Latin letters (A, B, C, etc.). 
* This program handles invalid formats gracefully and adheres to good Go coding practices, including unit testing.

## Objectives

- **Compile Successfully**: Ensure the Go program compiles without errors.
- **Assemble Tetrominoes**: Fit the tetrominoes into the smallest possible square.
- **Print Identifiers**: Use uppercase Latin letters to identify each tetromino in the solution.
- **File Handling**: The program reads from a text file and expects at least one tetromino.
- **Error Handling**: Print "ERROR" for bad file formats or invalid tetromino formats.
- **Good Practices**: Follow Go coding conventions and practices.
- **Testing**: Provide unit tests to ensure the correctness of the implementation.

## Requirements

- **Go Version**: 1.18 or higher
- **Dependencies**: Standard Go packages only

## Usage

To run the Tetris Optimizer program, use the following command:

```
go run . <path-to-text-file>
```

* Replace `<path-to-text-file>` with the path to your text file containing the list of tetrominoes.

**Example 1:** Valid tetromino having correct file format 

Suppose you have a file named validsample.txt with the following content:

```bash
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....
```

* Run the program with:

```bash
go run . validsample.txt
```

* If the program successfully solves the puzzle, it will print the smallest square grid containing the tetrominoes, using letters to represent each piece. For the above example, the output will be:

```bash
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

**Example2 :** Invalid Tetromino

* If the file format is incorrect or the tetrominoes cannot be placed into a square, the program will print:



```bash
ERROR
```

## File Format

The text file should contain tetrominoes in the following format:

1. Each tetromino is represented by a series of lines.
2. Each line contains only # (part of the tetromino) or . (empty space).
3. Tetrominoes are separated by blank lines.
4. There must be at least one tetromino in the file.

## Rules of the Game
1. Contains `#` and `.` only.
2. Each `#` character must be connected to at least one other `#` character (horizontally or vertically; not diagonally).
3. Tetromino must have exactly 4 `#` characters.
4. Total number of connections between `#` characters must be at least 6.

## Contributing

* Contributions are welcome! 
* Please ensure that your changes adhere to Go coding standards and include relevant tests. 
* Submit a pull request with a detailed description of your changes.

