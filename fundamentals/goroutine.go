package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("goroutine")
	go f("direct")

	time.Sleep(time.Second)
	fmt.Println("done")
}
