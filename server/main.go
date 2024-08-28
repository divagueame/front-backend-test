package main

import (
	"bufio"
	"fmt"
	"net"
)

func confirmHandshake(writer *bufio.Writer) {
	_, err := writer.WriteString("hello\n")
	fmt.Println("Handshake Confirmation ii")
	if err != nil {
		fmt.Println("error....", err)
		return
	}
	writer.Flush()
}

func handleConn(conn net.Conn) {

	defer conn.Close()

	writer := bufio.NewWriter(conn)
	confirmHandshake(writer)

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if line == "quit\r\n" {
			break
		}

		if err != nil {
			break
		}
		fmt.Println("Client sent:", line)

	}

	_, err2 := writer.WriteString("(15,15)")
	if err2 != nil {
		fmt.Println("error.2...", err2)
		return
	}
	writer.Flush()

}

func main() {

	conn, err := net.Listen("tcp", ":8124")
	if err != nil {
		fmt.Println("Error...")
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Starting server at 8124")

	for {
		conn, err := conn.Accept()
		if err != nil {
			fmt.Println("Error connecting...")
			continue
		}

		go handleConn(conn)
	}
}
