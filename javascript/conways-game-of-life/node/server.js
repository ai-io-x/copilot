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