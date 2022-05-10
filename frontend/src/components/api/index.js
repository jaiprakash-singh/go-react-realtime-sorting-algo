// api/index.js
var socket = new WebSocket("ws://localhost:9001/ws");

let connect = cb => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
        console.log("OnMessage:",msg);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };