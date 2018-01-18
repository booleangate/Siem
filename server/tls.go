package server

import (
	"log"
	"net"
)

func HandleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		log.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
