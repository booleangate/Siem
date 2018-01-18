package main

import (
	"crypto/rand"
	"crypto/tls"
	"flag"
	svr "github.com/stinkyfingers/Siem/server"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	port       = flag.String("port", "8080", "--port=<port>")
	servertype = flag.String("s", "tcp", "-s=udp")
	public     = flag.String("public", "", "-public")
	private    = flag.String("private", "", "-private")
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
	case "tls":
		log.Print("Starting TLS")
		t()
	case "http":
		log.Print("Starting HTTP")
		h()
	default:
		log.Fatal("Invalid servertype -s")
	}

}

func h() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		log.Print(string(b))
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(":"+*port, mux))
}

func t() {
	//TLS
	// https://jannewmarch.gitbooks.io/network-programming-with-go-golang-/content/security/tls.html
	cert, err := tls.LoadX509KeyPair(*public, *private)
	if err != nil {
		log.Fatal(err)
		return
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	service := "0.0.0.0:1200"

	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Accepted")
		go svr.HandleClient(conn)
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
