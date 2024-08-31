package main

import (
	"divagueame/canvas-server/state"
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

func parseCommand(line string) (string, error) {
	if line == "coord\n\n" {
		return "co", nil
	}

	return line, nil
}

func handleConn(conn net.Conn) {

	defer conn.Close()

	writer := bufio.NewWriter(conn)
	confirmHandshake(writer)

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if line == "quit\r\n" {
			break
		}

		// parseCommand(line)
		// command, err := parseCommand(line)
		// if err == nil {
		// 	fmt.Println("Parsed command:", command)
		// }

	}

	_, err2 := writer.WriteString("(15,15)")
	if err2 != nil {
		fmt.Println("error.2...", err2)
		return
	}
	writer.Flush()

}

func main() {
	state.Initialize()
	// state.PrintState()
	// state.PrintCoord()
	state.PrintCanvas()
	// state.ChangeDirection("right", 2)
	// state.PrintState()

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
