package server

import (
	"log"
	"net"
)

func UdpConnectionHandler(conn *net.UDPConn) error {
	conn.SetReadBuffer(65507)
	buf := make([]byte, 65507)
	n, add, err := conn.ReadFromUDP(buf)
	log.Println("Received ", string(buf[0:n]), " from ", add, "\n")
	log.Print("++++++++++++++ END LOG +++++++++++++++++++\n")

	return err
}
