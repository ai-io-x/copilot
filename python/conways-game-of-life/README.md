A simple implementation of Conway's Game of Life in Python that runs in the terminal:

```python
import random
import os
import time

SIZE = 20
ALIVE = 1
DEAD = 0

def initialize_board():
    return [[random.randint(0, 1) for _ in range(SIZE)] for _ in range(SIZE)]

def print_board(board):
    os.system('clear')  # Use 'cls' for Windows
    for row in board:
        print(''.join('O' if cell == ALIVE else ' ' for cell in row))
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
                alive_count += board[new_x][new_y]
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