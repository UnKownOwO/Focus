package main

import (
	"focus/game"
	"focus/http"
	"log"
	"time"
)

func main() {
	log.Println("启动HTTP服务器中...")
	go http.NewHttpServer("localhost", 443).Start()
	log.Println("启动完毕")

	log.Println("启动游戏服务器中...")
	go game.NewGameServer("0.0.0.0", 22102).Start()
	log.Println("启动完毕")

	for {
		time.Sleep(1000)
	}
}
