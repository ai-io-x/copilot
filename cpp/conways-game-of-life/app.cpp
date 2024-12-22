#include <iostream>
#include <vector>
#include <cstdlib>
#include <unistd.h>

#define SIZE 20
#define ALIVE 1
#define DEAD 0

using namespace std;

void initializeBoard(vector<vector<int>>& board);
void printBoard(const vector<vector<int>>& board);
void nextGeneration(const vector<vector<int>>& board, vector<vector<int>>& newBoard);
int countAliveNeighbors(const vector<vector<int>>& board, int x, int y);

int main() {
    vector<vector<int>> board(SIZE, vector<int>(SIZE));
    vector<vector<int>> newBoard(SIZE, vector<int>(SIZE));

    initializeBoard(board);

    while (true) {
        printBoard(board);
        nextGeneration(board, newBoard);

        // Copy newBoard to board
        board = newBoard;

        usleep(500000); // Pause for 500 milliseconds
    }

    return 0;
}

void initializeBoard(vector<vector<int>>& board) {
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            board[i][j] = rand() % 2;
        }
    }
}

void printBoard(const vector<vector<int>>& board) {
    system("clear"); // Clear the terminal
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            cout << (board[i][j] == ALIVE ? "O" : " ");
        }
        cout << endl;
    }
    cout << endl;
}

void nextGeneration(const vector<vector<int>>& board, vector<vector<int>>& newBoard) {
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

int countAliveNeighbors(const vector<vector<int>>& board, int x, int y) {
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