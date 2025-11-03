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

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
