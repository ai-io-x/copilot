A simple implementation of Conway's Game of Life in Java that runs in the terminal:

```java
import java.util.Random;

public class GameOfLife {
    private static final int SIZE = 20;
    private static final int ALIVE = 1;
    private static final int DEAD = 0;

    public static void main(String[] args) throws InterruptedException {
        int[][] board = new int[SIZE][SIZE];
        initializeBoard(board);
        while (true) {
            printBoard(board);
            board = nextGeneration(board);
            Thread.sleep(500); // Pause for 500 milliseconds
        }
    }

    private static void initializeBoard(int[][] board) {
        Random random = new Random();
        for (int i = 0; i < SIZE; i++) {
            for (int j = 0; j < SIZE; j++) {
                board[i][j] = random.nextInt(2);
            }
        }
    }

    private static void printBoard(int[][] board) {
        for (int i = 0; i < SIZE; i++) {
            for (int j = 0; j < SIZE; j++) {
                System.out.print(board[i][j] == ALIVE ? "O" : " ");
            }
            System.out.println();
        }
        System.out.println();
    }

    private static int[][] nextGeneration(int[][] board) {
        int[][] newBoard = new int[SIZE][SIZE];

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

    private static int countAliveNeighbors(int[][] board, int x, int y) {
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
}
```

To run this code:

1. Save it to a file named `GameOfLife.java`.
2. Open a terminal and navigate to the directory containing `GameOfLife.java`.
3. Compile the code with `javac GameOfLife.java`.
4. Run the compiled code with `java GameOfLife`.

This will display the Game of Life in the terminal, updating every 500 milliseconds.

Similar code found with 2 license types