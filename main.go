package main

import (
	"flag"
	svr "github.com/stinkyfingers/Siem/server"
	"log"
	"net"
)

var (
	port = flag.String("port", "8080", "--port=<port>")
)

func main() {
	flag.Parse()

	server, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer server.Close()

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//handle the connection, on it's own thread, per connection
		go svr.ConnectionHandler(connection)

	}
}
