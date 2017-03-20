package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:2003")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	io := bufio.NewReader(conn)
	var err error
	for firstLine, i :=  "", 0; err == nil && i < 100; firstLine, err = io.ReadString('\n') {
		i++
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("The line I read: " + firstLine)
	}
	conn.Close()
}
