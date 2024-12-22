A simple implementation of Conway's Game of Life in Go that runs in the terminal:

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
	SIZE  = 20
	ALIVE = 1
	DEAD  = 0
)

func main() {
	board := initializeBoard()
	for {
		printBoard(board)
		board = nextGeneration(board)
		time.Sleep(500 * time.Millisecond)
	}
}

func initializeBoard() [][]int {
	board := make([][]int, SIZE)
	for i := range board {
		board[i] = make([]int, SIZE)
		for j := range board[i] {
			board[i][j] = rand.Intn(2)
		}
	}
	return board
}

func printBoard(board [][]int) {
	clearScreen()
	for i := range board {
		for j := range board[i] {
			if board[i][j] == ALIVE {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func nextGeneration(board [][]int) [][]int {
	newBoard := make([][]int, SIZE)
	for i := range newBoard {
		newBoard[i] = make([]int, SIZE)
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

func countAliveNeighbors(board [][]int, x, y int) int {
	aliveCount := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newX := x + i
			newY := y + j
			if newX >= 0 && newX < SIZE && newY >= 0 && newY < SIZE {
				aliveCount += board[newX][newY]
			}
		}
	}
	return aliveCount
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
```

To run this code:

1. Save it to a file named `game_of_life.go`.
2. Open a terminal and navigate to the directory containing `game_of_life.go`.
3. Run the code with `go run game_of_life.go`.

This will display the Game of Life in the terminal, updating every 500 milliseconds.