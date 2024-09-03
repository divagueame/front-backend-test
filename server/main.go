package main

import (
	"divagueame/canvas-server/state"
	// "divagueame/canvas-server/canvas"
	"bufio"
	"fmt"
	"strings"
	// "log"
	"net"
)

func confirmHandshake(writer *bufio.Writer) {
	_, err := writer.WriteString("hello\n")
	if err != nil {
		fmt.Println("error....", err)
		return
	}
	writer.Flush()
}

func parseCommand(line string) (string, error) {
	// fmt.Println("parsing command", line)
	if line == "coord\n\n" {
		return "coord", nil
	} else if line == "render\n\n" {
		return "render", nil
	}

	return line, nil
}

func handleConn(conn net.Conn) {

	defer conn.Close()

	resStr := ""
	// resStr := "(15,15)"

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

		command, err := parseCommand(line)
		if err != nil {
			fmt.Println("Error!!", err)
			break
		}
		// fmt.Println("Parsed command:", command)

		if strings.TrimSpace(command) == "render" {
			resStr = state.GetCanvas()
		}

		if strings.TrimSpace(command) == "coord" {
			fmt.Println("Adding:", state.GetPositionStr())
			resStr += state.GetPositionStr()
			resStr += "\n"
			// resStr = "(15,15)"
		}

		// fmt.Println("", resStr)

		_, err2 := writer.WriteString(resStr)
		if err2 != nil {
			fmt.Println("error.2...", err2)
			return
		}
		writer.Flush()
	}

}

func main() {
	state.Initialize()
	// pos := state.GetPosition()
	// fmt.Print("meowj", state.GetPositionStr())
	// state.PrintState()
	// state.PrintCoord()
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
