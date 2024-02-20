package main

import (
	"fmt"
	"net"
	"os"
)

const PING = "*1\r\n$4\r\nping\r\n"

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Listening on port 6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// close connection when finished
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			os.Exit(1)
		}
		received := string(buf[:n])
		switch command := string(received[1]); command {
		// case "1":
		// 	returnPing(conn)
		// case "2":
		// 	returnEcho(conn, received)
		default:
			fmt.Printf("%s.\n", received)
		}
	}
}

func returnPing(conn net.Conn) {
	message := []byte("+PONG\r\n")
	_, err := conn.Write(message)
	if err != nil {
		fmt.Println("Error pong back to ping command: ", err.Error())
		os.Exit(1)
	}
}

func returnEcho(conn net.Conn, arg string) {
	message := []byte(arg)
	_, err := conn.Write(message[10:])
	if err != nil {
		fmt.Println("Error pong back to ping command: ", err.Error())
		os.Exit(1)
	}
}

// func redisParser(received string) (command string, arg string, err error) {

// }
