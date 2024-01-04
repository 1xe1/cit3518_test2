package main

import (
	"fmt"
	"net"
	"os"
)

const (
	serverAddress = "localhost:5000"
)

func main() {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	conn.Write([]byte(username))
	conn.Write([]byte(password))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	response := string(buffer[:n])
	fmt.Println("Server response:", response)
}
