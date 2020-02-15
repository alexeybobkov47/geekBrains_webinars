package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/sacOO7/gowebsocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New("ws://echo.websocket.org/")

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Print("Соединение установлено")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Printf("Ошибка при соединении: %v", err)
	}

	socket.OnPingReceived = func(_ string, socket gowebsocket.Socket) {
		log.Print("Получен Ping")
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Printf("Сообщение %v", message)
	}

	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Print("Соединение закрыто")
			socket.Close()
			return
		case <-time.After(time.Second * 10):
			socket.SendText("Hello, World!")
		}
	}
}
