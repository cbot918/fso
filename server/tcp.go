package main

import (
	"fmt"
	"net"

	// "github.com/cbot918/fss/utils"
	u "github.com/cbot918/fss/utils"
)

var lf = fmt.Printf

type TCP struct {
	Port string
	HA   *u.HTTPAnalyzer
}

func NewTCP(port string) *TCP {
	return &TCP{
		Port: port,
		HA:   u.NewHTTPAnalyzer(),
	}
}

func (tcp *TCP) Run() error {
	u.Lg("tcp listening: 5555")
	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		go tcp.HandleConn(conn)
	}

}

func (tcp *TCP) HandleConn(conn net.Conn) {
	defer conn.Close()
	// buffer with 128 byte
	buf := make([]byte, 128)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			// if client disconnect
			if err.Error() == "EOF" {
				u.Lg("\nuser disconnect\n")
				break
			}
			u.Lg(err)
			continue
		}
		// determine a http request
		if tcp.HA.IsHTTP(buf) {
			u.Lg("get a http request")

		}
		lf("%#v\n", string(buf))
	}

	// m := []byte("hi from server")
	// _, err = conn.Write(m)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
