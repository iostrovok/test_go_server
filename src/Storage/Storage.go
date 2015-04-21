package Storage

import (
	"Common"
	"Server/Pipe"
	"regexp"
	"strings"
	"sync"
)

const (
	LEN_PRIFIX int = 3
)

// construct a regexp to extract values:
var (
	re = regexp.MustCompile("[^A-Za-z]+")
)

type Storage struct {
	listWords map[string]Part
	letters   *Part
	debug     bool
	Mu        *sync.Mutex
}

func New(idDebug ...bool) (*Storage, error) {
	debug := false
	if len(idDebug) > 0 {
		debug = idDebug[0]
	}
	return &Storage{
		listWords: map[string]Part{},
		letters:   newPart(Common.LetterType, ""),
		debug:     debug,
		Mu:        &sync.Mutex{},
	}, nil
}

func (storage *Storage) StartOne(WG *sync.WaitGroup, fromUser *Pipe.Pipe, n int) {

	WG.Add(1)

	go func(WG *sync.WaitGroup, ChanOut <-chan string, n int) {
		defer WG.Done()

		for {
			select {
			case m, isGood := <-ChanOut:

				if !isGood {
					return
				}

				str := re.ReplaceAllString(strings.ToLower(m), " ")
				wordsList := strings.Fields(str)

				for _, w := range wordsList {
					if w == "" {
						continue
					}

					key := ""
					if len(w) < LEN_PRIFIX {
						key = w + "      "
						key = key[:LEN_PRIFIX]
					} else {
						key = w[:LEN_PRIFIX]
					}

					lets := strings.Split(w, "")
					for _, l := range lets {
						storage.letters.Add(l)
					}

					p, find := storage.listWords[key]
					if !find {
						storage.Mu.Lock()

						p, find = storage.listWords[key]
						if !find {
							p = *newPart(Common.WordType, key, storage.debug)
							storage.listWords[key] = p
						}

						storage.Mu.Unlock()
					}
					p.chanAdd <- w
				}
			}
		}
	}(WG, fromUser.ChanOut, n)
}

func (storage *Storage) StartReadWeb(WG *sync.WaitGroup, fromWeb <-chan Common.UserRequest) {

	WG.Add(1)

	go func(WG *sync.WaitGroup, fromWeb <-chan Common.UserRequest) {
		defer WG.Done()

		for {
			select {
			case m, isGood := <-fromWeb:
				if !isGood {
					return
				}
				total := len(storage.listWords) + 1
				storage.letters.GetTop(total, m)

				for _, p := range storage.listWords {
					p.GetTop(total, m)
				}

			}
		}
	}(WG, fromWeb)
}

func (storage *Storage) Start(WG *sync.WaitGroup,
	fromUser *Pipe.Pipe, fromWeb <-chan Common.UserRequest) {

	storage.StartReadWeb(WG, fromWeb)

	for i := 0; i < 100; i++ {
		storage.StartOne(WG, fromUser, i)
	}
}
