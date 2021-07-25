package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	messages := make(chan string, 3)
	go func() {
		messages <- "ping"
		messages <- "pong"
		messages <- "done"
	}()

	for msg := range messages {
		fmt.Println(msg)
	}

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
