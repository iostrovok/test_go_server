package Common

import (
	"sort"
)

/*
	Here.
	1-1) Sets type message "UserRequest" which web server sends to application server.
	1-2) Sets function for using "UserRequest".

	2-1) Sets type message "ResultMessage" which application server sends to web server with result data.
	2-2) Sets function for using "ResultMessage".
*/

type StorageType string

const (
	LenfromWebClient             = 50
	WordType         StorageType = "word"
	LetterType       StorageType = "letter"
)

var fromWebClient chan UserRequest = nil

func GetFromWebClient() chan UserRequest {
	if fromWebClient == nil {
		fromWebClient = make(chan UserRequest, LenfromWebClient)
	}
	return fromWebClient
}

type OneWord struct {
	W string
	N int
}

type ResultMessage struct {
	MyType        StorageType
	Total         int
	Words         []OneWord
	CountWord     int
	CountUniqWord int
}

type UserRequest struct {
	N  int
	Ch chan ResultMessage
}

func SendUserRequest(n int) chan ResultMessage {
	ch := make(chan ResultMessage, 10)
	mess := UserRequest{
		N:  n,
		Ch: ch,
	}
	GetFromWebClient() <- mess
	return ch
}

func SortAndCut(n int, list []OneWord) []string {

	sort.Sort(ListOneWord(list))
	out := []string{}
	for i := 0; i < len(list) && i < n; i++ {
		if list[i].W == "" {
			break
		}
		out = append(out, list[i].W)
	}

	return out
}

/* SORT __REVERSE__ WORDS  ------->

USING:

import "sort"

out := []ListOneWord{}
...
sort.Sort(ListOneWord.List(out))

*/
type ListOneWord []OneWord

func (d ListOneWord) Len() int {
	return len(d)
}

func (d ListOneWord) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d ListOneWord) Less(i, j int) bool {
	return d[i].N > d[j].N
}

/* <------- SORT DATA */
