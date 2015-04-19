package main

/*
Using:
go run ./src/test_script/add_word_line.go -i=10000000
*/

import (
	"flag"
	"fmt"
	"net"
)

var (

	//Max int = 10000000
	Max int = 100

	line10 string = `We would like you to build a simple Go application.`
	line20 string = `We would like you to build a simple Go application. When started, it
will listen on port but this may.`
	line30 string = `We would like you to build a simple Go application. When started, it
will listen on port but this may be configurable through a
command line flag. Clients will be `
	line40 string = `We would like you to build a simple Go application. 
When started, it will listen on port but this may
 be configurable through a command  line flag. Clients will be
  able to connect to this port and arbitrary natural to `
)

func main() {

	readFlags()

	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Println(err)
		return
	}

	countWords := 0

	for countWords < Max {

		fmt.Fprintf(conn, line10)
		fmt.Fprintf(conn, line20)
		fmt.Fprintf(conn, line30)
		fmt.Fprintf(conn, line40)

		countWords += 10 + 20 + 30 + 40

		if countWords%1000000 == 0 {
			fmt.Printf(" ... countWords: %d\n", countWords)
		}
	}

	fmt.Printf("countWords: %d\n", countWords)
	fmt.Printf("FINISH\n")
}

func readFlags() {
	i := flag.Int("i", 0, "-i=1000")

	flag.Parse()

	if *i != 0 {
		Max = *i
	}
}
