package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action      string              `json: "action"`
	Message     string              `json: "message"`
	UserName    string              `json: "username"`
	MessageType string              `json: "message_type"`
	Conn        WebSocketConnection `json: "-"`
}

type WsJsonResponse struct {
	Action  string `json: "action"`
	Message string `json: "message"`
	UserID  int    `json: "user_id"`
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var clients = make(map[WebSocketConnection]string)
var WsChan = make(chan WsPayload)

func (app *application) WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
}
