

const url = "ws://" + window.location.host + window.location.pathname + "/ws";
var wsOpened = false;
initWebsocket();
function initWebsocket() {
    ws = new WebSocket(url);
    ws.onopen = function (event) {
        wsOpened = true;
        console.log("ws connected");
        ws.send('hello server');
    };
    ws.onmessage = function (msg) {
        if (msg.data) {
            console.log("recieved : " + msg.data);
        }
    };
}