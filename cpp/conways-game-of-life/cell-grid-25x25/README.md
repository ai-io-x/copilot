Here's a simple implementation of Conway's Game of Life in C++ that runs in the terminal with a fixed 25x25 cell grid. Each live cell is represented by 'O' and each dead cell by '-'.

```cpp
#include <iostream>
#include <vector>
#include <cstdlib>
#include <unistd.h>

#define SIZE 25
#define ALIVE 'O'
#define DEAD '-'

using namespace std;

void initializeBoard(vector<vector<char>>& board);
void printBoard(const vector<vector<char>>& board);
void nextGeneration(const vector<vector<char>>& board, vector<vector<char>>& newBoard);
int countAliveNeighbors(const vector<vector<char>>& board, int x, int y);

int main() {
    vector<vector<char>> board(SIZE, vector<char>(SIZE));
    vector<vector<char>> newBoard(SIZE, vector<char>(SIZE));

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

void initializeBoard(vector<vector<char>>& board) {
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            board[i][j] = rand() % 2 == 0 ? DEAD : ALIVE;
        }
    }
}

void printBoard(const vector<vector<char>>& board) {
    system("clear"); // Clear the terminal
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            cout << board[i][j];
        }
        cout << endl;
    }
    cout << endl;
}

void nextGeneration(const vector<vector<char>>& board, vector<vector<char>>& newBoard) {
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

int countAliveNeighbors(const vector<vector<char>>& board, int x, int y) {
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
```

To compile and run this code:

1. Save it to a file named `game_of_life.cpp`.
2. Open a terminal and navigate to the directory containing `game_of_life.cpp`.
3. Compile the code with `g++ -o game_of_life game_of_life.cpp`.
4. Run the compiled code with `./game_of_life`.

This will display the Game of Life in the terminal, updating every 500 milliseconds.

Similar code found with 2 license types