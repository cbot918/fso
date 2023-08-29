package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var lg = fmt.Println
var lf = fmt.Printf

func main() {
	conn, err := net.Dial("tcp", ":5555")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	lg("connected to server")
	var count int
	for {
		count += 1
		message := "Hello from client" + fmt.Sprintf(" %d", count)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending data:", err)
			break
		}
		lg("hello sended")
		// buffer := make([]byte, 1024)
		// n, err := conn.Read(buffer)
		// if err != nil {
		// 	fmt.Println("Error reading data:", err)
		// 	continue
		// }

		// response := string(buffer[:n])
		// fmt.Println("Received from server:", response)

		time.Sleep(1 * time.Second) // Wait before sending the next message

	}

	// m := []byte("hi from client")
	// _, err = conn.Write(m)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// buf := make([]byte, 1024)
	// _, err = conn.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(buf))
}
