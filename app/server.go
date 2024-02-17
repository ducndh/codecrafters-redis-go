package main

import (
	"fmt"
	"net"
	"os"
)

const PING = "*1\r\n$4\r\nping\r\n"

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}

	fmt.Printf("received %d bytes", n)
	fmt.Printf("received the following data: %s", string(buf[:n]))

	switch received := string(buf[:n]); received {
	case PING:
		returnPing(conn)
	default:
		fmt.Printf("%s.\n", received)
	}
}

func returnPing(conn net.Conn) {
	message := []byte("+PONG\r\n")
	n, err := conn.Write(message)
	if err != nil {
		fmt.Println("Error pong back to ping command: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("sent %d bytes", n)
	fmt.Printf("sent the following data: %s", string(message))
}
