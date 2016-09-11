package server

import (
	"bytes"
	"io"
	"log"
	"net"
)

func ConnectionHandler(connection net.Conn) error {
	var err error
	defer connection.Close()

	var buf bytes.Buffer
	io.Copy(&buf, connection)

	log.Print(string(buf.Bytes()))

	return err
}
