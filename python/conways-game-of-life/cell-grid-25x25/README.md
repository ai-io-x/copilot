Here's a simple implementation of Conway's Game of Life in Python that runs in the terminal with a fixed 25x25 cell grid. Each live cell is represented by 'O' and each dead cell by '-'.

```python
import random
import os
import time

SIZE = 25
ALIVE = 'O'
DEAD = '-'

def initialize_board():
    return [[ALIVE if random.randint(0, 1) == 1 else DEAD for _ in range(SIZE)] for _ in range(SIZE)]

def print_board(board):
    os.system('clear')  # Use 'cls' for Windows
    for row in board:
        print(''.join(row))
    print()

def next_generation(board):
    new_board = [[DEAD] * SIZE for _ in range(SIZE)]
    for i in range(SIZE):
        for j in range(SIZE):
            alive_neighbors = count_alive_neighbors(board, i, j)
            if board[i][j] == ALIVE:
                new_board[i][j] = ALIVE if 2 <= alive_neighbors <= 3 else DEAD
            else:
                new_board[i][j] = ALIVE if alive_neighbors == 3 else DEAD
    return new_board

def count_alive_neighbors(board, x, y):
    alive_count = 0
    for i in range(-1, 2):
        for j in range(-1, 2):
            if i == 0 and j == 0:
                continue
            new_x, new_y = x + i, y + j
            if 0 <= new_x < SIZE and 0 <= new_y < SIZE:
                if board[new_x][new_y] == ALIVE:
                    alive_count += 1
    return alive_count

def main():
    board = initialize_board()
    while True:
        print_board(board)
        board = next_generation(board)
        time.sleep(0.5)

if __name__ == "__main__":
    main()
```

To run this code:

1. Save it to a file named `game_of_life.py`.
2. Open a terminal and navigate to the directory containing `game_of_life.py`.
3. Run the code with `python3 game_of_life.py`.

This will display the Game of Life in the terminal, updating every 500 milliseconds.