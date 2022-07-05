package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
	"net/http"
	"os/exec"
	"strings"
)

type ShellController struct{}

type Command struct {
	Command string `json:"command"`
}

func NewShellController() *ShellController {
	return &ShellController{}
}

func (pc *ShellController) WS(c *gin.Context) {
	var upgrader websocket.Upgrader
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.Error(err, "err")
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			pc.ErrorAndClose(conn, messageType)
			break
		}

		var cmd Command
		if err := json.Unmarshal(p, &cmd); err != nil {
			pc.ErrorAndClose(conn, messageType)
			break
		}

		logging.Infof("Succeeded in getting commands %#v", cmd)
		split := strings.Split(cmd.Command, " ")
		var output []byte
		if len(split) == 1 {
			logging.Info("exec with no args")
			output, err = exec.Command(cmd.Command).CombinedOutput()
			// even though error happened, just return error message and keep connection for user to continue sending commands
			if err != nil {
				logging.Error(err, "err")
				pc.Error(conn, cmd.Command, messageType, err)
			}
		} else {
			logging.Info("exec with args")
			fmt.Println(split)
			output, err = exec.Command(split[0], split[1:]...).CombinedOutput()
			fmt.Println(string(output))
			// even though error happened, just return error message and keep connection for user to continue sending commands
			if err != nil {
				logging.Error(err, "err")
				pc.Error(conn, split[0], messageType, err)
			}
		}

		if len(output) == 0 {
			continue
		}

		if err := conn.WriteMessage(messageType, output); err != nil {
			pc.ErrorAndClose(conn, messageType)
			break
		}
	}
}

func (pc *ShellController) ErrorAndClose(conn *websocket.Conn, messageType int) {
	if err := conn.WriteMessage(messageType, []byte("error")); err != nil {
		logging.Error(err, "err")
		conn.Close()
	}
}

func (pc *ShellController) Error(conn *websocket.Conn, cmd string, messageType int, e error) {
	if strings.Contains(e.Error(), "executable file not found") {
		msg := "sh: command not found: " + cmd + "\n"
		if err := conn.WriteMessage(messageType, []byte(msg)); err != nil {
			logging.Error(err, "err")
		}
		return
	}
	if err := conn.WriteMessage(messageType, []byte(e.Error()+"\n")); err != nil {
		logging.Error(err, "err")
	}
}
