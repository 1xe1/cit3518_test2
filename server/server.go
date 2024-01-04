package main

import (
	"fmt"
	"net"
	"os"
)

const (
	port = ":5000"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := buffer[:n]
	username := string(data)

	n, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data = buffer[:n]
	password := string(data)

	if username == "std1" && password == "p@ssw0rd" {
		conn.Write([]byte("Hello"))
	} else {
		conn.Write([]byte("รหัสผิดพลาด"))
	}
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
