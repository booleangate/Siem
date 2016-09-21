package server

import (
	"bytes"
	"io"
	"log"
	"net"
)

func TcpConnectionHandler(connection net.Conn) error {
	var err error
	defer connection.Close()

	var buf bytes.Buffer
	io.Copy(&buf, connection)

	log.Print(string(buf.Bytes()), "\n")
	log.Print("++++++++++++++ END LOG +++++++++++++++++++\n")
	return err
}
