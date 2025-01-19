package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("The worker is stopped")
			return
		default:
			fmt.Println("The worker is working")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
