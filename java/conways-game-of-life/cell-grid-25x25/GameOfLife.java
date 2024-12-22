import java.util.Random;

public class GameOfLife {
    private static final int SIZE = 25;
    private static final char ALIVE = 'O';
    private static final char DEAD = '-';

    public static void main(String[] args) throws InterruptedException {
        char[][] board = new char[SIZE][SIZE];
        initializeBoard(board);
        while (true) {
            printBoard(board);
            board = nextGeneration(board);
            Thread.sleep(500); // Pause for 500 milliseconds
        }
    }

    private static void initializeBoard(char[][] board) {
        Random random = new Random();
        for (int i = 0; i < SIZE; i++) {
            for (int j = 0; j < SIZE; j++) {
                board[i][j] = random.nextBoolean() ? ALIVE : DEAD;
            }
        }
    }

    private static void printBoard(char[][] board) {
        System.out.print("\033[H\033[2J"); // Clear the terminal
        System.out.flush();
        for (int i = 0; i < SIZE; i++) {
            for (int j = 0; j < SIZE; j++) {
                System.out.print(board[i][j]);
            }
            System.out.println();
        }
        System.out.println();
    }

    private static char[][] nextGeneration(char[][] board) {
        char[][] newBoard = new char[SIZE][SIZE];

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

        return newBoard;
    }

    private static int countAliveNeighbors(char[][] board, int x, int y) {
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
}