package main

/*
Using:
cat ./test_data.txt |  go run ./src/test_script/add_word_std.go
*/

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("START\n")

	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}

	for hasMoreInLine {
		fmt.Fprintf(conn, string(line)+"\n")
		line, hasMoreInLine, err = bio.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("FINISH\n")

}
