package Stat

import (
	"Common"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	DefaultCountWords int    = 5
	InternalJsonError string = `{"error":"Internal Erroror. Internal code 1.","data":null,"result":1}`
)

/* Handlers ====> */
func H(res http.ResponseWriter, req *http.Request) {

	n, debug := GetParams(req)

	wordTop := []Common.OneWord{}
	lettersTop := []Common.OneWord{}
	countWord := 0
	countUniqWord := 0

	// Main request to get the top of words and letters
	resCh := Common.SendUserRequest(n)

	count := -1
	for count != 0 {
		select {
		case m, isGood := <-resCh:
			if !isGood {
				break
			}

			if count == -1 {
				count = m.Total - 1
			} else {
				count--
			}

			if m.MyType == Common.WordType {
				wordTop = append(wordTop, m.Words...)
				countWord += m.CountWord
				countUniqWord += m.CountUniqWord
			} else if m.MyType == Common.LetterType {
				lettersTop = append(lettersTop, m.Words...)
			}
		}
	}

	nStr := strconv.Itoa(n)
	out := map[string]interface{}{
		"top_" + nStr + "_words":   Common.SortAndCut(n, wordTop),
		"top_" + nStr + "_letters": Common.SortAndCut(n, lettersTop),
		"count":                    countWord,
	}
	if debug {
		out["debug_words"] = wordTop
		out["debug_letters"] = lettersTop
		out["count_uniq"] = countUniqWord
	}

	SendJsonSuccess(res, out)
}

/* <===== Handlers */
func GetParams(req *http.Request) (int, bool) {

	debug := false
	var err error

	n, err := strconv.Atoi(string(req.FormValue("n")))
	if n == 0 || err != nil {
		n = DefaultCountWords
	}

	isDebug := string(req.FormValue("d"))
	if isDebug == "true" {
		debug = true
	}

	return n, debug
}

func SendJsonSuccess(res http.ResponseWriter, d ...interface{}) {
	var data interface{}
	if len(d) > 0 {
		data = d[0]
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		b = []byte(InternalJsonError)
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(b)
}
