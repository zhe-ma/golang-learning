package main

import (
	"fmt"
	"net"
)

//-----------------------------------
// TCP Server

func startTcpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:9003")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err)
					break
				}

				msg := string(buf[:n])
				fmt.Println(conn.RemoteAddr(), ":", msg)

				if msg == "quit" {
					conn.Write([]byte("Bye!"))
					conn.Close()
					break
				}
			}
		}()
	}

}

func main() {
	startTcpServer()
}
