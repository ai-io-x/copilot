use rand::Rng;
use std::{thread, time};

const SIZE: usize = 25;
const ALIVE: char = 'O';
const DEAD: char = '-';

fn main() {
    let mut board = initialize_board();
    loop {
        print_board(&board);
        board = next_generation(&board);
        thread::sleep(time::Duration::from_millis(500));
    }
}

fn initialize_board() -> Vec<Vec<char>> {
    let mut board = vec![vec![DEAD; SIZE]; SIZE];
    let mut rng = rand::thread_rng();
    for i in 0..SIZE {
        for j in 0..SIZE {
            board[i][j] = if rng.gen_bool(0.5) { ALIVE } else { DEAD };
        }
    }
    board
}

fn print_board(board: &Vec<Vec<char>>) {
    print!("\x1B[2J\x1B[1;1H"); // Clear the terminal
    for row in board {
        for &cell in row {
            print!("{}", cell);
        }
        println!();
    }
    println!();
}

fn next_generation(board: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let mut new_board = vec![vec![DEAD; SIZE]; SIZE];
    for i in 0..SIZE {
        for j in 0..SIZE {
            let alive_neighbors = count_alive_neighbors(board, i, j);
            if board[i][j] == ALIVE {
                new_board[i][j] = if alive_neighbors < 2 || alive_neighbors > 3 { DEAD } else { ALIVE };
            } else {
                new_board[i][j] = if alive_neighbors == 3 { ALIVE } else { DEAD };
            }
        }
    }
    new_board
}

fn count_alive_neighbors(board: &Vec<Vec<char>>, x: usize, y: usize) -> usize {
    let mut alive_count = 0;
    for i in -1..=1 {
        for j in -1..=1 {
            if i == 0 && j == 0 {
                continue;
            }
            let new_x = x as isize + i;
            let new_y = y as isize + j;
            if new_x >= 0 && new_x < SIZE as isize && new_y >= 0 && new_y < SIZE as isize {
                if board[new_x as usize][new_y as usize] == ALIVE {
                    alive_count += 1;
                }
            }
        }
    }
    alive_count
}