package main

/*
Using:
cat ./test_data.txt |  go run ./src/test_script/add_word_std.go
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strings"
)

var (
	fileRead    string = ""
	regSplitEnd        = regexp.MustCompile("\n")
	re                 = regexp.MustCompile("[^A-Za-z]+")
)

func main() {

	readFlags()

	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Println(err)
		return
	}

	data, errReadFile := ioutil.ReadFile(fileRead)
	if errReadFile != nil {
		fmt.Println(errReadFile)
		return
	}

	d := string(data)
	//d = string(data)

	list := regSplitEnd.Split(string(d), -1)

	fmt.Printf("Length text: %d\n", len(string(d)))
	fmt.Printf("Count lines: %d\n", len(list))

	maxLine := 0
	countWords := 0
	sendWords := 10000000

	words := []string{}

	for _, str := range list {
		if maxLine < len(str) {
			maxLine = len(str)
		}

		s := re.ReplaceAllString(strings.ToLower(string(str)), " ")
		w := strings.Fields(s)
		words = append(words, w...)
	}

	fmt.Printf("len(words): %d\n", len(words))
	fmt.Println("Start sending.")

	i := 0
	for i < len(words) && i < sendWords {

		line := ""
		for j := 0; j < 10 && i < len(words); j++ {
			line += words[i] + " "
			countWords++
			i++
		}

		//fmt.Printf("%s", line+"\n")

		fmt.Fprintf(conn, line+"\n")

		if countWords%1000000 == 0 {
			fmt.Printf(" ... countWords: %d\n", countWords)
		}
	}

	fmt.Printf("maxLine: %d\n", maxLine)
	fmt.Printf("countWords: %d\n", countWords)
	fmt.Printf("FINISH\n")
}

func readFlags() {
	file := flag.String("f", "", "-f")

	flag.Parse()

	fileRead = *file

	if fileRead == "" {
		fmt.Println("Please set file name: -f=./my_file.txt")
		os.Exit(0)
	}
}
