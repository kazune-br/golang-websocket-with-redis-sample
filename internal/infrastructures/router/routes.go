package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kazune-br/golang-websocket-with-redis-sample/internal/adapter/controller"
	"github.com/kazune-br/golang-websocket-with-redis-sample/internal/adapter/controller/ws"
	"github.com/kazune-br/golang-websocket-with-redis-sample/internal/adapter/middleware"
	"github.com/kazune-br/golang-websocket-with-redis-sample/internal/infrastructures/redis"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
)

func Run() {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Use(middleware.SetLogger())

	r.GET("/healthcheck", controller.HealthCheck)
	r.GET("/ws/ping", ws.NewPingController().WS)
	r.GET("/ws/subscribe", ws.NewSubscribeController(redis.NewRedisConnector().Client).WS)
	r.GET("/ws/shell", ws.NewShellController().WS)

	logging.Info("starting server")
	if err := r.Run(fmt.Sprintf(":%d", 8000)); err != nil {
		logging.Fatal(err, "failed to initialize server")
		panic(err)
	}
}
