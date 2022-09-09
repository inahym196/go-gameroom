
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
    const putInfo = { 'name': 'first', 'request': 'put', 'piece': clickPiece }
    ws.send(JSON.stringify(putInfo))
}

const canvas = <HTMLCanvasElement>document.getElementById('screen')!
initCanvas(canvas)
const url = "ws://" + window.location.host + window.location.pathname + "/ws";
const ws = new WebSocket(url);
var wsOpened = false;
ws.onopen = function (event: Event) {
    wsOpened = true;
    console.log("ws connected");
    const joinInfo = { 'name': 'first', 'request': 'board' }
    ws.send(JSON.stringify(joinInfo));
};
ws.onmessage = function (msg) {
    if (msg.data) {
        const board = JSON.parse(msg.data);
        if (board) {
            console.log(board)
        }
    }
};
canvas.addEventListener('click', (e) => putPieceEvent(e, ws));