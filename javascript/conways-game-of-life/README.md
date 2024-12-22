Sure, here's an implementation of Conway's Game of Life using only HTML and JavaScript that runs entirely in the client browser. This code will create a fixed 25x25 cell grid, with each live cell represented by 'O' and each dead cell by '-'. The grid will be updated directly in the browser without any server-side code.

### HTML (index.html)
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Conway's Game of Life</title>
    <style>
        body {
            font-family: monospace;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background-color: #000;
            color: #fff;
        }
        #grid {
            display: grid;
            grid-template-columns: repeat(25, 20px);
            grid-template-rows: repeat(25, 20px);
        }
        .cell {
            width: 20px;
            height: 20px;
            display: flex;
            justify-content: center;
            align-items: center;
        }
    </style>
</head>
<body>
    <div id="grid"></div>
    <script>
        const SIZE = 25;
        const ALIVE = 'O';
        const DEAD = '-';
        const grid = document.getElementById('grid');

        // Initialize the grid
        let board = initializeBoard();
        createGrid(board);

        function initializeBoard() {
            const board = [];
            for (let i = 0; i < SIZE; i++) {
                const row = [];
                for (let j = 0; j < SIZE; j++) {
                    row.push(Math.random() < 0.5 ? ALIVE : DEAD);
                }
                board.push(row);
            }
            return board;
        }

        function createGrid(board) {
            grid.innerHTML = '';
            for (let i = 0; i < SIZE; i++) {
                for (let j = 0; j < SIZE; j++) {
                    const cell = document.createElement('div');
                    cell.className = 'cell';
                    cell.id = `cell-${i}-${j}`;
                    cell.textContent = board[i][j];
                    grid.appendChild(cell);
                }
            }
        }

        function updateGrid(board) {
            for (let i = 0; i < SIZE; i++) {
                for (let j = 0; j < SIZE; j++) {
                    const cell = document.getElementById(`cell-${i}-${j}`);
                    cell.textContent = board[i][j];
                }
            }
        }

        function nextGeneration(board) {
            const newBoard = [];
            for (let i = 0; i < SIZE; i++) {
                const row = [];
                for (let j = 0; j < SIZE; j++) {
                    const aliveNeighbors = countAliveNeighbors(board, i, j);
                    if (board[i][j] === ALIVE) {
                        row.push(aliveNeighbors < 2 || aliveNeighbors > 3 ? DEAD : ALIVE);
                    } else {
                        row.push(aliveNeighbors === 3 ? ALIVE : DEAD);
                    }
                }
                newBoard.push(row);
            }
            return newBoard;
        }

        function countAliveNeighbors(board, x, y) {
            let aliveCount = 0;
            for (let i = -1; i <= 1; i++) {
                for (let j = -1; j <= 1; j++) {
                    if (i === 0 && j === 0) continue;
                    const newX = x + i;
                    const newY = y + j;
                    if (newX >= 0 && newX < SIZE && newY >= 0 && newY < SIZE) {
                        if (board[newX][newY] === ALIVE) {
                            aliveCount++;
                        }
                    }
                }
            }
            return aliveCount;
        }

        function gameLoop() {
            board = nextGeneration(board);
            updateGrid(board);
            setTimeout(gameLoop, 500);
        }

        gameLoop();
    </script>
</body>
</html>
```

### Instructions to Run

1. Save the HTML code to a file named `index.html`.
2. Open `index.html` in a web browser.

This will display the Game of Life in the browser, updating every 500 milliseconds. The grid will be updated directly in the browser without any server-side code.

Similar code found with 2 license types