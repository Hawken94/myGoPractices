package main

import (
	"fmt"
	"net"
	"strings"
)

func connectHandler(c net.Conn) {
	fmt.Printf("come from %s\n", c.RemoteAddr())
	if c == nil {
		return
	}
	buf := make([]byte, 4096)

	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inString := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(inString, " ")
		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			fmt.Printf("Unsupport command :%s\n", inputs[0])
		}
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}

func main() {
	server, err := net.Listen("tcp", ":1082")
	if err != nil {
		fmt.Printf("Fail to start server,%s \n", err)
	}
	fmt.Println("Server started...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect,%s \n", err)
			break
		}
		go connectHandler(conn)
	}
}
