package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// create tcp connection
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// by using goroutine here, we can have multiple connection, each can run and not being blocked
		go func() {
			for {
				_, err := io.WriteString(conn, time.Now().Format("15:05:05\n"))
				if err != nil {
					return
				}
				time.Sleep(time.Second)
			}
		}()
	}

}
