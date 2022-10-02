
type WSSendType = 'join' | 'put' | "select-order" | 'request-info'

class WSSendData {
    type: WSSendType
    data: { [key: string]: any }
    constructor(type: WSSendType, data: { [key: string]: any }) {
        this.type = type
        this.data = data
    }
    ToString(): string {
        const data: { [key: string]: any } = { "type": this.type, "data": this.data }
        return JSON.stringify(data)
    }
}

type Status = "Selecting..." | "Waiting..." | "Game End"

function clickFirstOrder() {
    const sendData = new WSSendData('select-order', { 'order': 'first' })
    ws.send(sendData.ToString())
    console.log(sendData.ToString())
    document.getElementById("select-order")!.style.display = "none"
}

function clickDrawOrder() {
    const sendData = new WSSendData('select-order', { 'order': 'draw' })
    ws.send(sendData.ToString())
    console.log(sendData.ToString())
    document.getElementById("select-order")!.style.display = "none"
}

function initCanvas(canvas: HTMLCanvasElement) {
    const context = canvas.getContext('2d')!
    const board_image = new Image()
    board_image.src = '/public/images/board.png'
    board_image.onload = () => {
        canvas.width = board_image.width
        canvas.height = board_image.height
        context.drawImage(board_image, 0, 0)
    }
}
function isMyTurn(): boolean {
    console.log({ turn, order })
    if (turn === undefined || order === undefined) {
        return false
    } else if (turn % 2 === 0 && order === 'first') {
        return true
    } else if (turn % 2 === 1 && order === 'draw') {
        return true
    }
    return false
}

interface PieceArea {
    x: number,
    y: number,
    length: number,
    grid: number
}

function putPieceEvent(e: MouseEvent, ws: WebSocket) {
    const rect = canvas.getBoundingClientRect()
    const clientRatio = canvas.width / canvas.clientWidth
    const clickPoint = {
        x: (e.clientX - rect.left) * clientRatio,
        y: (e.clientY - rect.top) * clientRatio,
    }
    const pieceArea: PieceArea = { x: 70, y: 10, length: 500, grid: 9 }
    if (!(pieceArea.x <= clickPoint.x && clickPoint.x <= pieceArea.x + pieceArea.length) ||
        !(pieceArea.y <= clickPoint.y && clickPoint.y <= pieceArea.y + pieceArea.length)) {
        return
    }
    const clickPiece = {
        x: Math.floor((clickPoint.x - pieceArea.x) / (pieceArea.length / (pieceArea.grid + 1))),
        y: Math.floor((clickPoint.y - pieceArea.y) / (pieceArea.length / (pieceArea.grid + 1))),
    }
    console.log(clickPiece)
    const putInfo = { 'type': 'put', "data": { "putpoint": clickPiece } }
    if (isMyTurn() === true) {
        console.log('send')
        ws.send(JSON.stringify(putInfo))
    }
    else {
        console.log('not send')
    }
}

const canvas = <HTMLCanvasElement>document.getElementById('screen')!
initCanvas(canvas)
const url = "ws://" + window.location.host + window.location.pathname + "/ws"
const ws = new WebSocket(url)
var wsOpened = false
let turn: number
type orderType = 'first' | 'draw' | 'audience'
let order: orderType
let boardStatus: "Init" | "Setting" | "Waiting" | "Starting" | "End" | undefined
let hasOwn: boolean = false
ws.onopen = function (event: Event) {
    wsOpened = true
    console.log("ws connected")
    const sendData = new WSSendData("join", {})
    ws.send(JSON.stringify(sendData))
}
ws.onmessage = function (msg) {
    if (msg.data === undefined) return
    const data = JSON.parse(msg.data)
    console.log(data)

    switch (data.Type) {
        case "putresult":
            const pieces = data.data.pieces
            order = data.data.order || order
            turn = data.data.turn || turn
            console.log(data.data.turn, turn)
            updatePieces(canvas, pieces)
            break
        case "board-info":
            console.log("board-info")
            boardStatus = data.data.status
            order = data.data.order
            turn = data.data.turn
            console.log(data.data.turn, turn)
            switch (boardStatus) {
                case "Setting":
                    hasOwn = data.data.owner
                    if (hasOwn === true) {
                        document.getElementById("select-order")!.style.display = "block"
                        document.getElementById("status")!.innerHTML = "Select order"
                    } else {
                        document.getElementById("select-order")!.style.display = "none"
                        document.getElementById("status")!.innerHTML = "Opponent selecting order"
                    }
                    break
            }
    }
}
canvas.addEventListener('click', (e) => putPieceEvent(e, ws))

function updatePieces(canvas: HTMLCanvasElement, pieces: string[][]) {
    const context = canvas.getContext('2d')!
    const pieceArea: PieceArea = { x: 70, y: 10, length: 500, grid: 9 }
    context.save()
    context.translate(pieceArea.x, pieceArea.y)
    drawCrossAndCircle(context, pieceArea, pieces, "green")
    drawCrossAndCircle(context, pieceArea, pieces, "pink")
    context.restore()
}

function drawCrossLine(ctx: CanvasRenderingContext2D, x: number, y: number, width: number) {
    ctx.moveTo(x - width, y - width)
    ctx.lineTo(x + width, y + width)
    ctx.moveTo(x - width, y + width)
    ctx.lineTo(x + width, y - width)
}

function drawCrossAndCircle(ctx: CanvasRenderingContext2D, pieceArea: PieceArea, pieces: string[][], choice_color: string) {
    const COLOR: { [index: string]: { normal: string, dark: string } } = {
        green: {
            normal: "#90FF3B",
            dark: "#559723"
        },
        pink: {
            normal: "#FF22FF",
            dark: "#A723FF"
        }
    }

    for (let n = 0; n < 2; n++) {
        ctx.beginPath()
        ctx.lineWidth = 10
        let shadowX
        let shadowY
        if (n === 0) {
            ctx.strokeStyle = COLOR[choice_color].dark
            shadowX = 2
            shadowY = 3
        } else {
            ctx.strokeStyle = COLOR[choice_color].normal
            shadowX = 0
            shadowY = 0
        }
        for (let xi = 0; xi < pieceArea.grid + 1; xi++) {
            for (let yj = 0; yj < pieceArea.grid + 1; yj++) {
                const piece = pieces[yj][xi]
                if (piece === 'None' ||
                    choice_color === 'green' && piece.match(/P/) ||
                    choice_color === "pink" && piece.match(/G/)) continue
                const scaleWidth = pieceArea.length / (pieceArea.grid + 1)
                const posX = (scaleWidth / 2) + (scaleWidth * xi) + shadowX
                const posY = (scaleWidth / 2) + (scaleWidth * yj) + shadowY
                if (piece.match(/X/)) {
                    drawCrossLine(ctx, posX, posY, 15)
                } else if (piece.match(/O/)) {
                    ctx.moveTo(posX + 15, posY)
                    ctx.arc(posX, posY, 15, 0, Math.PI * 2, true)
                }
            }
        }
        ctx.stroke()
    }
}