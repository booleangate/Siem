package main

import (
	"flag"
	svr "github.com/stinkyfingers/Siem/server"
	"log"
	"net"
)

var (
	port       = flag.String("port", "8080", "--port=<port>")
	servertype = flag.String("s", "tcp", "-s=udp")
)

func main() {
	flag.Parse()
	switch *servertype {
	case "tcp":
		log.Print("Starting TCP")
		tcp()
	case "udp":
		log.Print("Starting UDP")
		udp()
	default:
		log.Fatal("Invalid servertype -s")
	}

}
func tcp() {
	//TCP
	server, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go svr.TcpConnectionHandler(conn)
	}
}
func udp() {
	// UDP
	addr, err := net.ResolveUDPAddr("udp", ":"+*port)
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		svr.UdpConnectionHandler(conn)
	}
}
