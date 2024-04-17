## Websocket Server and Client.

### Basic implementation of a websocket server and client.
- This project uses the [Gorilla WebSocket](https://github.com/gorilla/websocket) library.
- Broadcast the message to all connected clients.
- The websocket client can connect to the server, send and receive messages.

### Running this code.
- This code requires a working Go development environment.
- Then clone this repo with:
```bash
git clone https://github.com/lucaslopesx/websocket-server.git
```
- Change to the project directory:
```bash
cd websocket-server
```
- go to the directory and run:
```bash
go run server.go
```
- Open a web browser and visit:
```
http://localhost:8080/client
```

