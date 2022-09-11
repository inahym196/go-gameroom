
function initCanvas(canvas: HTMLCanvasElement) {
    const context = canvas.getContext('2d')!;
    const board_image = new Image()
    board_image.src = '/public/images/board.png'
    board_image.onload = () => {
        canvas.width = board_image.width
        canvas.height = board_image.height
        context.drawImage(board_image, 0, 0)
    }
}
function isMyTurn(): boolean {
    if (turn === undefined || order === undefined) {
        return false
    } else if (turn % 2 === 0 && order === 'first') {
        return true
    } else if (turn % 2 === 1 && order === 'draw') {
        return true
    }
    return false
}

function putPieceEvent(e: MouseEvent, ws: WebSocket) {
    const rect = canvas.getBoundingClientRect();
    const clientRatio = canvas.width / canvas.clientWidth
    const clickPoint = {
        x: (e.clientX - rect.left) * clientRatio,
        y: (e.clientY - rect.top) * clientRatio,
    }
    const pieceArea = { x: 70, y: 10, length: 500, grid: 9 }
    if (!(pieceArea.x <= clickPoint.x && clickPoint.x <= pieceArea.x + pieceArea.length) ||
        !(pieceArea.y <= clickPoint.y && clickPoint.y <= pieceArea.y + pieceArea.length)) {
        return
    }
    const clickPiece = {
        x: Math.floor((clickPoint.x - pieceArea.x) / (pieceArea.length / (pieceArea.grid + 1))),
        y: Math.floor((clickPoint.y - pieceArea.y) / (pieceArea.length / (pieceArea.grid + 1))),
    }
    console.log(clickPiece)
    const putInfo = { 'type': 'put', 'piece': clickPiece };
    if (isMyTurn() === true) {
        console.log('send');
        ws.send(JSON.stringify(putInfo));
    }
    else {
        console.log('not send');
    }
}

const canvas = <HTMLCanvasElement>document.getElementById('screen')!
initCanvas(canvas)
const url = "ws://" + window.location.host + window.location.pathname + "/ws";
const ws = new WebSocket(url);
var wsOpened = false;
let turn: number;
let order: 'first' | 'draw' | 'audience';
ws.onopen = function (event: Event) {
    wsOpened = true;
    console.log("ws connected");
    const joinInfo = { 'type': 'join' }
    ws.send(JSON.stringify(joinInfo));
};
ws.onmessage = function (msg) {
    if (msg.data) {
        const data = JSON.parse(msg.data);
        console.log(data);
        switch (data.Type) {
            case "join":
                turn = data.Board.turn;
                order = data.Order;
                break;
            case "board":
                console.log(data);
                break;
        }
    }
};
canvas.addEventListener('click', (e) => putPieceEvent(e, ws));