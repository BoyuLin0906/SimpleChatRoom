# SimpleChatRoom

## Setup

- Golang :
  - Purpose : Build a simple server using the GIN framework and Gorilla websocket
  - Environment : Window 10 Professional
  - install from VSCode (go 1.17)

## Usage

1. Run the server under the root project path (Window)
    ```
    go run .\main.go
    ```

2. Open the website `http://127.0.0.1:8080/`
    - input the username
    - input the chat message

## Todo

1. Optimize client mechanism
2. More fool-proof mechanism
3. Hub and different chat room
4. User service (DB + API)
5. Chat room history
6. Use message queue to send message

## Ref
- Gorilla websocket chat example :  https://github.com/gorilla/websocket/tree/master/examples/chat
