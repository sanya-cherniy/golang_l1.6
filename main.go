package main

import (
	"fmt"
	"time"
)

func worker(stop <-chan bool) {
	for {
		select {
		default:
			fmt.Println("The worker is working")
		case <-stop: // при записи сообщения в канал выходим из функции
			fmt.Println("The worker is stopped")
			return
		}
	}
}

func main() {
	stop := make(chan bool)
	go worker(stop)

	time.Sleep(2 * time.Second)
	stop <- true // Останавливаем горутину при помощи отправки сообщения в канал
}
