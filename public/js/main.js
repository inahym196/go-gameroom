"use strict";
const url = "ws://" + window.location.host + window.location.pathname + "/ws";
var wsOpened = false;
function loadCanvas(canvas) {
    const context = canvas.getContext('2d');
    const board_image = new Image();
    board_image.src = '/public/images/board.png';
    board_image.onload = () => {
        canvas.width = board_image.width;
        canvas.height = board_image.height;
        context.drawImage(board_image, 0, 0);
    };
}
const canvas = document.getElementById('screen');
const userAgent = navigator.userAgent;
loadCanvas(canvas);
canvas.addEventListener('click', (e) => {
    const rect = canvas.getBoundingClientRect();
    const clientRatio = canvas.width / canvas.clientWidth;
    const clickPoint = {
        x: (e.clientX - rect.left) * clientRatio,
        y: (e.clientY - rect.top) * clientRatio,
    };
    const pieceArea = {
        x: 70,
        y: 10,
        length: 500,
        grid: 9,
    };
    if (!(pieceArea.x <= clickPoint.x && clickPoint.x <= pieceArea.x + pieceArea.length) ||
        !(pieceArea.y <= clickPoint.y && clickPoint.y <= pieceArea.y + pieceArea.length)) {
        return;
    }
    const clickPiece = {
        x: Math.floor((clickPoint.x - pieceArea.x) / (pieceArea.length / (pieceArea.grid + 1))),
        y: Math.floor((clickPoint.y - pieceArea.y) / (pieceArea.length / (pieceArea.grid + 1))),
    };
    console.log(clickPiece);
});
