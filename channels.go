package main

import (
	"log"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- "foo"
			time.Sleep(time.Second * 1)
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:
			log.Println("c1", msg1)
		}
	}
}
