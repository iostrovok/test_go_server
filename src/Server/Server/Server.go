package Server

import (
	"Server/Pipe"
	"io"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	CONN_TYPE = "tcp"
)

// construct a regexp to extract values:
var (
	re = regexp.MustCompile("[^A-Za-z]+")
)

//type checkLetter func(rune) bool

type Server struct {
	Listener net.Listener
	debug    bool
}

func New(host string, port int, idDebug ...bool) (*Server, error) {
	// Listen for incoming connections.

	debug := false
	if len(idDebug) > 0 {
		debug = idDebug[0]
	}

	hostPort := strconv.Itoa(port)
	if port < 1 {
		log.Fatal("Bad port for connection: " + hostPort)
	}

	hostPort = ":" + hostPort
	if host != "" {
		hostPort = host + hostPort
	}

	if debug {
		log.Println("Application server tries listening: " + CONN_TYPE + ", " + hostPort)
	}

	l, err := net.Listen(CONN_TYPE, hostPort)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return nil, err
	}

	return &Server{
		Listener: l,
		debug:    debug,
	}, nil
}

func (server *Server) Listen(WG *sync.WaitGroup, outChan *Pipe.Pipe) {

	WG.Add(1)
	go func() {
		defer WG.Done()

		// Close the listener when the application closes.
		defer server.Listener.Close()
		for {
			// Listen for an incoming connection.
			conn, err := server.Listener.Accept()
			if err != nil {
				log.Panicf("Error accepting: %s\n", err)
			}
			// Handle connections in a new goroutine.
			go server.handler(conn, outChan)
		}

	}()

	if server.debug {
		log.Println("Application server - 'Listen'")
	}
}

func (server *Server) handler(conn net.Conn, outChan *Pipe.Pipe) {

	var buf = make([]byte, 1024)
	str := ""
	noLast := true

	for noLast {

		n, err := conn.Read(buf[0:])
		if err == io.EOF {
			noLast = false
		} else if err != nil {
			return
		}

		s := re.ReplaceAllString(string(buf[:n]), " ")

		i := strings.LastIndex(s, " ")
		if i == len(s)-1 {
			outChan.Push(str + s)
			str = ""
		} else if i > -1 {
			outChan.Push(str + s[:i])
			str = s[i+1:]
		} else {
			str += s
		}
	}

	if str != "" {
		outChan.Push(str)
	}

}
