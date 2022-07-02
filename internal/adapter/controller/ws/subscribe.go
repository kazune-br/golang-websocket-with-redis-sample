package ws

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
	"net/http"
)

type SubscribeController struct {
	rdb *redis.Client
}

func NewSubscribeController(rdb *redis.Client) *SubscribeController {
	return &SubscribeController{
		rdb,
	}
}

func (sc *SubscribeController) WS(c *gin.Context) {
	var upgrader websocket.Upgrader
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.Error(err, "err")
		return
	}

	go sc.subscribe(conn)
}

func (sc *SubscribeController) subscribe(conn *websocket.Conn) {
	ctx := context.Background()
	subscriber := sc.rdb.Subscribe(ctx, "sample")
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
