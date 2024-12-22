Sure, here's an implementation of Conway's Game of Life using HTML, JavaScript, and WebSocket. This code will run entirely in the client browser, with the WebSocket used to update the grid without re-creating it.

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
        for (let i = 0; i < SIZE; i++) {
            for (let j = 0; j < SIZE; j++) {
                const cell = document.createElement('div');
                cell.className = 'cell';
                cell.id = `cell-${i}-${j}`;
                cell.textContent = DEAD;
                grid.appendChild(cell);
            }
        }

        // Connect to WebSocket server
        const socket = new WebSocket('ws://localhost:8080');

        socket.onmessage = function(event) {
            const board = JSON.parse(event.data);
            updateGrid(board);
        };

        function updateGrid(board) {
            for (let i = 0; i < SIZE; i++) {
                for (let j = 0; j < SIZE; j++) {
                    const cell = document.getElementById(`cell-${i}-${j}`);
                    cell.textContent = board[i][j] ? ALIVE : DEAD;
                }
            }
        }
    </script>
</body>
</html>
```

### JavaScript (server.js)
This is a simple WebSocket server that you can run with Node.js to simulate the game logic and send updates to the client.

```javascript
const WebSocket = require('ws');

const SIZE = 25;
const ALIVE = 1;
const DEAD = 0;

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', ws => {
    let board = initializeBoard();

    setInterval(() => {
        board = nextGeneration(board);
        ws.send(JSON.stringify(board));
    }, 500);
});

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
                aliveCount += board[newX][newY];
            }
        }
    }
    return aliveCount;
}
```

### Instructions to Run

1. Save the HTML code to a file named `index.html`.
2. Save the JavaScript server code to a file named `server.js`.
3. Open a terminal and navigate to the directory containing `server.js`.
4. Run `npm install ws` to install the WebSocket library.
5. Start the WebSocket server with `node server.js`.
6. Open `index.html` in a web browser.

This setup will display the Game of Life in the browser, updating every 500 milliseconds based on the WebSocket server's updates.

Similar code found with 1 license type