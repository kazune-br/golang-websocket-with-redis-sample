package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
	"net/http"
)

var upgrader websocket.Upgrader

type WSController struct {
	rdb *redis.Client
}

func NewWSController(rdb *redis.Client) *WSController {
	return &WSController{
		rdb,
	}
}

func (w *WSController) WS(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.Error(err, "err")
		return
	}

	go w.subscribe(conn)

	for {
		messageType, _, err := conn.ReadMessage()
		if err != nil {
			logging.Error(err, "err")
			return
		}
		if err := conn.WriteMessage(messageType, []byte("pong")); err != nil {
			logging.Error(err, "err")
			return
		}
	}
}

func (w *WSController) subscribe(conn *websocket.Conn) {
	ctx := context.Background()
	subscriber := w.rdb.Subscribe(ctx, "sample")
	for {
		// through redis pub/sub, reading a message published by background process
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			logging.Error(err, "error")
			return
		}

		logging.Infof("#%v", msg)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			logging.Error(err, "err")
			return
		}
	}
}
