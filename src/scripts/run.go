package main

import (
	"Common"
	"Server/Pipe"
	"Server/Server"
	"Storage"
	"Web/WebServer"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

var (
	ServerPort int  = 5555
	WebPort    int  = 8080
	Debug      bool = false
)

func main() {

	readFlags()
	runtime.GOMAXPROCS(maxParallelism())

	/*
		Common vars
	*/
	var WG sync.WaitGroup = sync.WaitGroup{}
	newWord := Pipe.New()

	// Start storage for our word and letters
	stor, errStor := Storage.New(Debug)
	if errStor != nil {
		log.Panicf("%s", errStor)

	}
	stor.Start(&WG, newWord, Common.GetFromWebClient())

	// Start web server
	http, errWeb := WebServer.New(WebPort, Debug)
	if errWeb != nil {
		log.Panicf("%s", errWeb)

	}
	http.Listen(&WG)

	// Start application server
	server, errSer := Server.New("", ServerPort, Debug)
	if errSer != nil {
		log.Panicf("%s", errSer)
	}
	server.Listen(&WG, newWord)

	// Running...
	WG.Wait()
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return 2 * maxProcs
	}
	return 2 * numCPU
}

func readFlags() {
	help := flag.Bool("h", false, "-h")
	helpLong := flag.Bool("help", false, "-help")
	webport := flag.Int("webport", 8080, "-webport=8080")
	port := flag.Int("port", 5555, "-port=5555")
	debug := flag.Bool("d", false, "-d")

	flag.Parse()

	if *help || *helpLong {
		viewHelp()
		os.Exit(0)
	}

	Debug = *debug
	ServerPort = *port
	WebPort = *webport

	if ServerPort < 1 || 65535 < ServerPort {
		log.Panicf("Bad port nubmer %d. You have to set from 1 to 65535\n", ServerPort)
	}
	if WebPort < 1 || 65535 < WebPort {
		log.Panicf("Bad port nubmer %d. You have to set from 1 to 65535\n", WebPort)
	}
	if WebPort == ServerPort {
		log.Panicf("You have to set different ports for web server and application. Now they are %d.\n", WebPort)
	}
}

func viewHelp() {

	fmt.Println(`

Simple application for processing text from tcp connect and providing statistical data through http.

Using

> go run ./src/scripts/run.go
> go run ./src/scripts/run.go [-d] [-webport=NUMBER] [-port=NUMBER]

Options:
 	-webport=WEB_SERVER_PORT Port for web server. Default it's 8080
 	-port=APP_SERVER_PORT Port for application server. Default it's 5555
 	-d Debug mode
 	-h, -help This help message

Example:

go run ./src/scripts/run.go -d -webport=8081 -port=6666

Geting statistical data through http:

http://localhost:8080/stats?n=100
http://localhost:8080/stats?n=100&d=true

Params:
 	n=number. Count top 
 	d=true.  Debug mode

Please report bugs to <ostrovok\@gmail.com>.
`)

}
