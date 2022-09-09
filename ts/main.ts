

const url = "ws://" + window.location.host + window.location.pathname + "/ws";
var wsOpened = false;

function createDice(context: CanvasRenderingContext2D, x: number, y: number, grid: number, length: number) {
    if ([0, grid].includes(x) || [0, grid].includes(y)) {
        createDice6(context, length);
    } else if ([1, grid - 1].includes(x) || [1, grid - 1].includes(y)) {
        createDice4(context, length);
    } else if ([2, grid - 2].includes(x) || [2, grid - 2].includes(y)) {
        if ([2, grid - 2].includes(x) && [2, grid - 2].includes(y)) {
            createDice1(context, length, 12);
        } else {
            createDice2(context, length);
        }
    } else {
        createDice1(context, length, 8);
    }
}

function createDice1(context: CanvasRenderingContext2D, length: number, radius: number) {
    drawCircle(context, length / 2, length / 2, radius);
}

function createDice2(context: CanvasRenderingContext2D, length: number) {
    const [short, long] = [(length / 7) * 2, (length / 7) * 5];
    const radius = 6;
    drawCircle(context, short, length / 2, radius);
    drawCircle(context, long, length / 2, radius);
}

function createDice4(context: CanvasRenderingContext2D, length: number) {
    const [short, long] = [(length / 7) * 2, (length / 7) * 5];
    const radius = 6;
    drawCircle(context, short, short, radius);
    drawCircle(context, short, long, radius);
    drawCircle(context, long, long, radius);
    drawCircle(context, long, short, radius);
}

function createDice6(context: CanvasRenderingContext2D, length: number) {
    const radius = 6;
    context.save();
    for (let i = 0; i < 2; i++) {
        context.translate(length / 3, 0);
        for (let i = 0; i < 3; i++) {
            drawCircle(context, 0, 13 * (i + 1), radius);
        }
    }
    context.restore();
}
function drawCircle(context: CanvasRenderingContext2D, posX: number, posY: number, radius: number) {
    context.moveTo(radius + posX, posY);
    context.arc(posX, posY, radius, 0, Math.PI * 2, true);
}

interface Board {
    width: number,
    height: number,
    center: { x: number, y: number }
}

interface PieceArea {
    length: number,
    grid: number,
    x: number,
    y: number,
    center: number,
    pieceLength: number
}


class View {
    context: CanvasRenderingContext2D
    constructor(canvas: HTMLCanvasElement, userAgent: string) {
        this.context = canvas.getContext('2d')!;
        const board: Board = { width: 640, height: 520, center: { x: 640 / 2, y: 520 / 2 } }
        const pieceArea: PieceArea = { length: 500, grid: 9, x: 70, y: 10, center: 500 / 2, pieceLength: 500 / 10 }
        canvas.width = board.width;
        canvas.height = board.height;
        this.drawBoard(board, pieceArea);
    }

    createRoundRect(ctx: CanvasRenderingContext2D, x: number, y: number, width: number, height: number, radius: number) {
        ctx.moveTo(x + radius, y);
        ctx.lineTo(x + width - radius, y);
        ctx.arcTo(x + width, y, x + width, y + radius, radius);
        ctx.lineTo(x + width, y + height - radius);
        ctx.arcTo(x + width, y + height, x + width - radius, y + height, radius);
        ctx.lineTo(x + radius, y + height);
        ctx.arcTo(x, y + height, x, y + height - radius, radius);
        ctx.lineTo(x, y + radius);
        ctx.arcTo(x, y, x + radius, y, radius);
        ctx.closePath();
    }

    drawBoardFrame(context: CanvasRenderingContext2D, board: Board) {
        context.save();
        context.beginPath();
        const gradient = (context: CanvasRenderingContext2D, board: Board) => {
            const _gradient = context.createRadialGradient(
                board.center.x, board.center.x, 0,
                board.center.y, board.center.y, board.width
            );
            _gradient.addColorStop(0, '#444444');
            _gradient.addColorStop(1, '#222222');
            return _gradient;
        };
        context.fillStyle = gradient(context, board);
        context.strokeStyle = '#444444';
        context.lineWidth = 2;
        this.createRoundRect(context, 0, 0, board.width, board.height, 50);
        context.fill();
        context.stroke();
        context.restore();
    }

    drawBoardPieceArea(context: CanvasRenderingContext2D, pieceArea: PieceArea) {
        context.save();
        context.translate(pieceArea.x, pieceArea.y);
        context.beginPath();
        context.strokeStyle = 'black';
        context.lineWidth = 1;
        context.shadowBlur = 2;
        context.shadowOffsetX = -1;
        context.shadowOffsetY = -1;
        context.shadowColor = 'silver';
        for (let y = 0; y < pieceArea.grid + 1; y++) {
            for (let x = 0; x < pieceArea.grid + 1; x++) {
                context.save();
                context.translate(x * pieceArea.pieceLength, y * pieceArea.pieceLength);
                this.createRoundRect(context, 1, 1, pieceArea.pieceLength - 2, pieceArea.pieceLength - 2, pieceArea.pieceLength / 10);
                context.restore();
            }
        }
        context.stroke();
        context.restore();
    }

    drawDices(context: CanvasRenderingContext2D, pieceArea: PieceArea) {
        context.save();
        context.translate(pieceArea.x, pieceArea.y);
        context.beginPath();
        context.strokeStyle = '#464646';
        context.lineWidth = 1;
        context.shadowBlur = 1;
        context.shadowOffsetX = -1;
        context.shadowOffsetY = -1;
        context.shadowColor = 'rgba(30,30,30,1)';
        for (let y = 0; y < pieceArea.grid + 1; y++) {
            for (let x = 0; x < pieceArea.grid + 1; x++) {
                context.save();
                context.translate(x * pieceArea.pieceLength, y * pieceArea.pieceLength);
                createDice(context, x, y, pieceArea.grid, pieceArea.pieceLength);
                context.restore();
            }
        }
        context.stroke();
        context.restore();
    }

    drawLineupPieceArea(context: CanvasRenderingContext2D, pieceArea: PieceArea) {
        context.save();
        context.translate(pieceArea.x + pieceArea.length, pieceArea.y);
        context.beginPath();
        context.shadowColor = '#555555';
        context.fillStyle = 'rgba(46,46,46,1)';
        context.shadowBlur = 1;
        context.shadowOffsetX = 2;
        context.shadowOffsetY = 1;
        for (let i = 1; i < pieceArea.grid + 1; i++) {
            if (i === 5) continue;
            context.moveTo(pieceArea.pieceLength, i * pieceArea.pieceLength);
            context.arc(30, i * pieceArea.pieceLength, 23, 0, Math.PI * 2);
        }
        context.fill();
        context.restore();
    }

    drawPointArea(context: CanvasRenderingContext2D, pieceArea: PieceArea) {
        context.save();
        context.translate(pieceArea.x, pieceArea.y);
        context.shadowColor = '#555555';
        context.fillStyle = 'rgba(46,46,46,1)';
        context.shadowBlur = 1;
        context.shadowOffsetX = 2;
        context.shadowOffsetY = 1;
        context.translate(-20, 0);
        for (let i = 1; i < 40; i++) {
            context.moveTo(5, i * 12.5);
            context.arc(5, i * 12.5, 5, 0, Math.PI * 2);
        }
        context.fill();
        context.restore();
    }

    drawWinSet(context: CanvasRenderingContext2D, pieceArea: PieceArea) {
        context.save();
        context.translate(pieceArea.x, pieceArea.y);
        context.shadowColor = '#555555';
        context.fillStyle = 'rgba(46,46,46,1)';
        context.shadowBlur = 1;
        context.shadowOffsetX = 2;
        context.shadowOffsetY = 1;
        context.translate(-70, 150);
        drawCircle(context, 25, 0, 18);
        drawCircle(context, 25, 25, 18);
        drawCircle(context, 25, 50, 18);
        drawCircle(context, 25, 100, 18);
        drawCircle(context, 25, 150, 18);
        drawCircle(context, 25, 175, 18);
        drawCircle(context, 25, 200, 18);
        context.fill();
        context.restore();
    }

    drawBoard(board: Board, pieceArea: PieceArea) {
        const context = this.context;
        this.drawBoardFrame(context, board);
        this.drawBoardPieceArea(context, pieceArea);
        this.drawDices(context, pieceArea);
        this.drawLineupPieceArea(context, pieceArea);
        this.drawPointArea(context, pieceArea);
        this.drawWinSet(context, pieceArea);
        context.restore();
    }
}
const canvas = <HTMLCanvasElement>document.getElementById('screen')
if (!(canvas instanceof HTMLCanvasElement)) {
    throw new Error('#screen is cannot ref')
}
const userAgent = navigator.userAgent;
const view = new View(canvas, userAgent);