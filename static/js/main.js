const othello = new Othello("othello-board", 670, 670);
const othelloInfo = new GameInfo("othello-info", 670, 200);

function updateBoardTest() {
    let boardArr = [
        [0,0,0,1,0,0,0,0],
        [0,0,0,0,0,0,0,0],
        [0,0,0,0,0,0,0,0],
        [0,3,1,2,1,0,0,0],
        [0,3,0,1,2,0,0,0],
        [0,3,0,0,0,0,0,0],
        [0,3,0,0,0,0,0,0],
        [0,0,0,0,0,0,0,0]
    ];
    othello.update(boardArr);
}

const socket = new WebSocket("ws://" + window.location.host + "/ws");
console.log("Attemptin Websocket Connection");

socket.onopen = () => {
    console.log("on open");
    socket.send("Hi From The Client!");
}

function wsSend(axisData) {
    let data = {"x":axisData["x"], "y":axisData["y"]};
    let jsonData = JSON.stringify(data);
    socket.send(jsonData);
}

socket.onclose = (event) => {
    socket.send("on close");
}

socket.onmessage = (message) => {
    console.log("on message");
    boardArr = [
        [0,0,0,0,0,0,0,0],
        [0,0,0,0,0,0,0,0],
        [0,0,0,0,0,0,0,0],
        [0,0,0,2,1,0,0,0],
        [0,0,0,1,2,0,0,0],
        [0,0,0,0,0,0,0,0],
        [0,0,0,0,0,0,0,0],
        [0,0,0,0,0,0,0,0]
    ];
    othello.update(boardArr);

    try {
        let boardArr = JSON.parse(message);
        othello.update(boardArr);
    } catch (err) {
        console.log("invalid json data.");
    }
}

socket.onerror = (error) => {
    console.log("on error");
}