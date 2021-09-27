package main

import (
	"github.com/kazune-br/golang-websocket-with-redis-sample/internal/infrastructures/router"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
)

func main() {
	logging.InitLogger()
	router.Run()
}
