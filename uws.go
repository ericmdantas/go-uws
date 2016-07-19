package uws

import (
	"golang.org/x/net/websocket"
)

type clientCb func(client Client)

var qOn []map[string]clientCb

type WS struct {
}

type Client struct {
	Conn *websocket.Conn
}

func On(token string, cb func(client Client)) {
}

func Emit(token string, info interface{}) {
}
