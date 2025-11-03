package main

import "github.com/gorilla/websocket"

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
