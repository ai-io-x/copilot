A simple implementation of Conway's Game of Life in C that runs in the terminal:

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define SIZE 20
#define ALIVE 1
#define DEAD 0

void initializeBoard(int board[SIZE][SIZE]);
void printBoard(int board[SIZE][SIZE]);
void nextGeneration(int board[SIZE][SIZE], int newBoard[SIZE][SIZE]);
int countAliveNeighbors(int board[SIZE][SIZE], int x, int y);

int main() {
    int board[SIZE][SIZE];
    int newBoard[SIZE][SIZE];

    initializeBoard(board);

    while (1) {
        printBoard(board);
        nextGeneration(board, newBoard);

        // Copy newBoard to board
        for (int i = 0; i < SIZE; i++) {
            for (int j = 0; j < SIZE; j++) {
                board[i][j] = newBoard[i][j];
            }
        }

        usleep(500000); // Pause for 500 milliseconds
    }

    return 0;
}

void initializeBoard(int board[SIZE][SIZE]) {
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            board[i][j] = rand() % 2;
        }
    }
}

void printBoard(int board[SIZE][SIZE]) {
    system("clear"); // Clear the terminal
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            printf(board[i][j] == ALIVE ? "O" : " ");
        }
        printf("\n");
    }
    printf("\n");
}

void nextGeneration(int board[SIZE][SIZE], int newBoard[SIZE][SIZE]) {
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            int aliveNeighbors = countAliveNeighbors(board, i, j);

            if (board[i][j] == ALIVE) {
                newBoard[i][j] = (aliveNeighbors < 2 || aliveNeighbors > 3) ? DEAD : ALIVE;
            } else {
                newBoard[i][j] = (aliveNeighbors == 3) ? ALIVE : DEAD;
            }
        }
    }
}

int countAliveNeighbors(int board[SIZE][SIZE], int x, int y) {
    int aliveCount = 0;
    for (int i = -1; i <= 1; i++) {
        for (int j = -1; j <= 1; j++) {
            if (i == 0 && j == 0) continue;
            int newX = x + i;
            int newY = y + j;
            if (newX >= 0 && newX < SIZE && newY >= 0 && newY < SIZE) {
                aliveCount += board[newX][newY];
            }
        }
    }
    return aliveCount;
}
```

To compile and run this code:

1. Save it to a file named `game_of_life.c`.
2. Open a terminal and navigate to the directory containing `game_of_life.c`.
3. Compile the code with `gcc -o game_of_life game_of_life.c`.
4. Run the compiled code with `./game_of_life`.

This will display the Game of Life in the terminal, updating every 500 milliseconds.

Similar code found with 2 license types