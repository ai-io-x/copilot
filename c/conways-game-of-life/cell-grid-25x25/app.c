#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define SIZE 25
#define ALIVE 'O'
#define DEAD '-'

void initializeBoard(char board[SIZE][SIZE]);
void printBoard(char board[SIZE][SIZE]);
void nextGeneration(char board[SIZE][SIZE], char newBoard[SIZE][SIZE]);
int countAliveNeighbors(char board[SIZE][SIZE], int x, int y);

int main() {
    char board[SIZE][SIZE];
    char newBoard[SIZE][SIZE];

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

void initializeBoard(char board[SIZE][SIZE]) {
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            board[i][j] = rand() % 2 == 0 ? DEAD : ALIVE;
        }
    }
}

void printBoard(char board[SIZE][SIZE]) {
    system("clear"); // Clear the terminal
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            printf("%c", board[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}

void nextGeneration(char board[SIZE][SIZE], char newBoard[SIZE][SIZE]) {
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

int countAliveNeighbors(char board[SIZE][SIZE], int x, int y) {
    int aliveCount = 0;
    for (int i = -1; i <= 1; i++) {
        for (int j = -1; j <= 1; j++) {
            if (i == 0 && j == 0) continue;
            int newX = x + i;
            int newY = y + j;
            if (newX >= 0 && newX < SIZE && newY >= 0 && newY < SIZE) {
                if (board[newX][newY] == ALIVE) {
                    aliveCount++;
                }
            }
        }
    }
    return aliveCount;
}