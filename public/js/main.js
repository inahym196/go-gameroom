

const url = "ws://" + window.location.host + window.location.pathname + "/ws";
var wsOpened = false;
initWebsocket();
function initWebsocket() {
    ws = new WebSocket(url);
    ws.onopen = function (event) {
        wsOpened = true;
        console.log("ws connected");
        ws.send('request board');
    };
    ws.onmessage = function (msg) {
        if (msg.data) {
            const board = JSON.parse(msg.data);
            console.log(board);
        }
    };
}