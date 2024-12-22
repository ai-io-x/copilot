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