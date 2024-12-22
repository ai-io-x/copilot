Here's a simple implementation of Conway's Game of Life in Go that runs in the terminal with a fixed 25x25 cell grid. Each live cell is represented by 'O' and each dead cell by '-'.

```go
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	SIZE  = 25
	ALIVE = 'O'
	DEAD  = '-'
)

func main() {
	board := initializeBoard()
	for {
		printBoard(board)
		board = nextGeneration(board)
		time.Sleep(500 * time.Millisecond)
	}
}

func initializeBoard() [][]rune {
	board := make([][]rune, SIZE)
	for i := range board {
		board[i] = make([]rune, SIZE)
		for j := range board[i] {
			if rand.Intn(2) == 0 {
				board[i][j] = DEAD
			} else {
				board[i][j] = ALIVE
			}
		}
	}
	return board
}

func printBoard(board [][]rune) {
	clearScreen()
	for i := range board {
		for j := range board[i] {
			fmt.Print(string(board[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func nextGeneration(board [][]rune) [][]rune {
	newBoard := make([][]rune, SIZE)
	for i := range newBoard {
		newBoard[i] = make([]rune, SIZE)
		for j := range newBoard[i] {
			aliveNeighbors := countAliveNeighbors(board, i, j)
			if board[i][j] == ALIVE {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					newBoard[i][j] = DEAD
				} else {
					newBoard[i][j] = ALIVE
				}
			} else {
				if aliveNeighbors == 3 {
					newBoard[i][j] = ALIVE
				} else {
					newBoard[i][j] = DEAD
				}
			}
		}
	}
	return newBoard
}

func countAliveNeighbors(board [][]rune, x, y int) int {
	aliveCount := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newX := x + i
			newY := y + j
			if newX >= 0 && newX < SIZE && newY >= 0 && newY < SIZE {
				if board[newX][newY] == ALIVE {
					aliveCount++
				}
			}
		}
	}
	return aliveCount
}

func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
```

To run this code:

1. Save it to a file named `game_of_life.go`.
2. Open a terminal and navigate to the directory containing `game_of_life.go`.
3. Run the code with `go run game_of_life.go`.

This will display the Game of Life in the terminal, updating every 500 milliseconds.