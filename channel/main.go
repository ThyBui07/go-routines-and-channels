package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		// infinity loop to prevent deadlocks
		for {
			time.Sleep(5 * time.Second)
			c1 <- time.Now().String()
		}
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- time.Now().String()
		time.Sleep(time.Second)
		fmt.Println("routine 2 closed")
	}()
	for {
		//select to listen from which channel has info to listen
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
