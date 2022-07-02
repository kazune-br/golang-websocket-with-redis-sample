package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
	"net/http"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (pc *PingController) WS(c *gin.Context) {
	var upgrader websocket.Upgrader
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.Error(err, "err")
		return
	}

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
