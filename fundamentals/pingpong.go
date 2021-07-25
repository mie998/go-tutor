package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "ping")
	go pong(pings, pongs)
	fmt.Println(<-pongs)
}
