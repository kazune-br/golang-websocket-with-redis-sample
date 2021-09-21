package main

import (
	"github.com/kazune-br/golang-websocket-with-redis-sample/internal/infrastructures/router"
	"github.com/kazune-br/golang-websocket-with-redis-sample/pkg/logging"
)

func main() {
	logging.InitLogger()
	//var ctx = context.Background()
	//for {
	//	// Publish a generated user to the new_users channel
	//	err := conn.Client.Publish(ctx,
	//		"new_users",
	//		struct {
	//			value string
	//		}{
	//			value: "Hello",
	//		},
	//	).Err()
	//	if err != nil {
	//		panic(err)
	//	}
	//	// Sleep random time
	//	rand.Seed(time.Now().UnixNano())
	//	n := rand.Intn(4)
	//	time.Sleep(time.Duration(n) * time.Second)
	//}

	//err := conn.Client.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := conn.Client.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//defer conn.Client.Close()
	router.Run()
}
