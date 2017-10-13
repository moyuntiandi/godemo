package main

import (
	"log"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://xxx.xxx.xxx"
	url := "ws://xxx.xxx.xxx:80/socket"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := ws.Write([]byte("all")); err != nil {
		log.Fatal(err)
	}
	var msg string
	for {
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(msg)
		//...........do anything you want
	}
}
