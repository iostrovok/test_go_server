package Storage

import (
	"Common"
	"fmt"
	"sync"
)

type MessWeb struct {
	total int
	mes   Common.UserRequest
}

type Part struct {
	totalWord int
	chanAdd   chan string
	chanWeb   chan MessWeb
	mu        *sync.Mutex
	words     *IndexW
	debug     bool
}

func newPart(myType Common.StorageType, prefix string, idDebug ...bool) *Part {

	debug := false
	if len(idDebug) > 0 {
		debug = idDebug[0]
	}

	chanAdd := make(chan string, 20)
	chanWeb := make(chan MessWeb, 20)
	totalWord := 0
	words := newIndexW()

	if debug && myType == Common.WordType {
		fmt.Printf("newPart. Type is %s. Prefix: %s.\n", myType, prefix)
	}

	go func() {
		for {
			select {
			case m, isGood := <-chanAdd:
				if !isGood {
					return
				}

				totalWord++
				words.Add(m)

				if debug && myType == Common.WordType {
					fmt.Printf("Type is %s. Prefix: %s. Total words: %d. Uniq words: %d. WORD: %s\n",
						myType, prefix, words.TotalWords, words.Len(), m)
				}

			case m, isGood := <-chanWeb:
				if !isGood {
					return
				}

				mes := Common.ResultMessage{
					MyType:        myType,
					Total:         m.total,
					Words:         words.GetТор(m.mes.N),
					CountWord:     words.TotalWords,
					CountUniqWord: words.Len(),
				}

				m.mes.Ch <- mes
			}
		}
	}()

	return &Part{
		totalWord: totalWord,
		chanAdd:   chanAdd,
		chanWeb:   chanWeb,
		words:     words,
		debug:     debug,
	}
}

func (p *Part) Add(str string) {
	p.chanAdd <- str
}

func (p *Part) GetTop(total int, m Common.UserRequest) {
	p.chanWeb <- MessWeb{
		total: total,
		mes:   m,
	}
}
