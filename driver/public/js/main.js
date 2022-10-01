"use strict";
function initCanvas(canvas) {
    const context = canvas.getContext('2d');
    const board_image = new Image();
    board_image.src = '/public/images/board.png';
    board_image.onload = () => {
        canvas.width = board_image.width;
        canvas.height = board_image.height;
        context.drawImage(board_image, 0, 0);
    };
}
function isMyTurn() {
    if (turn === undefined || order === undefined) {
        return false;
    }
    else if (turn % 2 === 0 && order === 'first') {
        return true;
    }
    else if (turn % 2 === 1 && order === 'draw') {
        return true;
    }
    return false;
}
function putPieceEvent(e, ws) {
    const rect = canvas.getBoundingClientRect();
    const clientRatio = canvas.width / canvas.clientWidth;
    const clickPoint = {
        x: (e.clientX - rect.left) * clientRatio,
        y: (e.clientY - rect.top) * clientRatio,
    };
    const pieceArea = { x: 70, y: 10, length: 500, grid: 9 };
    if (!(pieceArea.x <= clickPoint.x && clickPoint.x <= pieceArea.x + pieceArea.length) ||
        !(pieceArea.y <= clickPoint.y && clickPoint.y <= pieceArea.y + pieceArea.length)) {
        return;
    }
    const clickPiece = {
        x: Math.floor((clickPoint.x - pieceArea.x) / (pieceArea.length / (pieceArea.grid + 1))),
        y: Math.floor((clickPoint.y - pieceArea.y) / (pieceArea.length / (pieceArea.grid + 1))),
    };
    console.log(clickPiece);
    const putInfo = { 'type': 'put', "data": { "putpoint": clickPiece } };
    if (isMyTurn() === true) {
        console.log('send');
        ws.send(JSON.stringify(putInfo));
    }
    else {
        console.log('not send');
    }
}
const canvas = document.getElementById('screen');
initCanvas(canvas);
const url = "ws://" + window.location.host + window.location.pathname + "/ws";
const ws = new WebSocket(url);
var wsOpened = false;
let turn;
let order;
ws.onopen = function (event) {
    wsOpened = true;
    console.log("ws connected");
    const joinInfo = { 'type': 'join', "data": { "gameType": "XOGame" } };
    ws.send(JSON.stringify(joinInfo));
};
ws.onmessage = function (msg) {
    if (msg.data) {
        const data = JSON.parse(msg.data);
        console.log(data);
        if (data.turn !== undefined) {
            turn = data.turn;
        }
        switch (data.Type) {
            case "join":
                turn = data.Board.turn;
                order = data.Order;
                break;
            case "board":
                console.log(data);
                break;
            case "putresult":
                console.log(data.data.pieces);
                const pieces = data.data.pieces;
                updatePieces(canvas, pieces);
                break;
        }
    }
};
canvas.addEventListener('click', (e) => putPieceEvent(e, ws));
function updatePieces(canvas, pieces) {
    const context = canvas.getContext('2d');
    const pieceArea = { x: 70, y: 10, length: 500, grid: 9 };
    context.save();
    context.translate(pieceArea.x, pieceArea.y);
    drawCrossAndCircle(context, pieceArea, pieces, "green");
    drawCrossAndCircle(context, pieceArea, pieces, "pink");
    context.restore();
}
;
function drawCrossLine(ctx, x, y, width) {
    ctx.moveTo(x - width, y - width);
    ctx.lineTo(x + width, y + width);
    ctx.moveTo(x - width, y + width);
    ctx.lineTo(x + width, y - width);
}
function drawCrossAndCircle(ctx, pieceArea, pieces, choice_color) {
    const COLOR = {
        green: {
            normal: "#90FF3B",
            dark: "#559723"
        },
        pink: {
            normal: "#FF22FF",
            dark: "#A723FF"
        }
    };
    for (let n = 0; n < 2; n++) {
        ctx.beginPath();
        ctx.lineWidth = 10;
        let shadowX;
        let shadowY;
        if (n === 0) {
            ctx.strokeStyle = COLOR[choice_color].dark;
            shadowX = 2;
            shadowY = 3;
        }
        else {
            ctx.strokeStyle = COLOR[choice_color].normal;
            shadowX = 0;
            shadowY = 0;
        }
        for (let xi = 0; xi < pieceArea.grid + 1; xi++) {
            for (let yj = 0; yj < pieceArea.grid + 1; yj++) {
                const piece_shape = pieces[yj][xi];
                if (piece_shape === '')
                    continue;
                const scaleWidth = pieceArea.length / (pieceArea.grid + 1);
                const posX = (scaleWidth / 2) + (scaleWidth * xi) + shadowX;
                const posY = (scaleWidth / 2) + (scaleWidth * yj) + shadowY;
                if (piece_shape === "X") {
                    drawCrossLine(ctx, posX, posY, 15);
                }
                else if (piece_shape === "O") {
                    ctx.moveTo(posX + 15, posY);
                    ctx.arc(posX, posY, 15, 0, Math.PI * 2, true);
                }
            }
        }
        ctx.stroke();
    }
    ;
}
