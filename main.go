package main

import (
	"log"
	"simple_chatroom/router"
)

const port string = "8080"

func main() {
	router_main := router.InitServer()
	err := router_main.Run(":" + port)
	if err != nil {
		log.Fatalf("Start server: %+v", err)
	}
}
