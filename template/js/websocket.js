
const ws_url = "ws://127.0.0.1:8080/ws";
const socket = new WebSocket(ws_url);

const default_name_chat_box = "chat_box";
const default_name_chat_message = "chat_message";
const default_name_username = "username"
var chat_box = document.getElementById(default_name_chat_box);
var chat_message = document.getElementById(default_name_chat_message);
var username = document.getElementById(default_name_username);

socket.onmessage = (event) => {
    const message = JSON.parse(event.data);
    var item = document.createElement("div");
    item.innerText = `${message.username}: ${message.body}\n`
    chat_box.appendChild(item);
}

function connectToChat() {
    if (!username.value) {
      alert("Please enter a username");
      return;
    }
    socket.send(JSON.stringify({type: "join",
                                username: username.value,
                                body: "[join]"}));
}

function sendMessage() {
    socket.send(JSON.stringify({type: "message",
                                username: username.value,
                                body: chat_message.value}));
    chat_message.value = "";
}