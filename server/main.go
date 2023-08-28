package main

import "log"

const tcpPort = ":5555"

func main() {

	tcp := NewTCP(tcpPort)

	if err := tcp.Run(); err != nil {
		log.Fatal(err)
	}
}

// GET / HTTP/1.1\r\nHost: localhost:5555\r\nUser-Agent: curl/7.68.0\r\nAccept: */*\r\n\r\n
