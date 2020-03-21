package main

import (
	"log"
)

func send(chanel chan<- string) {
	const sendText string = "カンツォーーネ！！"

	log.Print("送信：" + sendText)
	chanel <- sendText
}

func main() {
	chanel := make(chan string)
	go send(chanel)

	sendText := <-chanel
	log.Print("受信：" + sendText)
}
