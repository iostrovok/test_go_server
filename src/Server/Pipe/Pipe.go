package Pipe

type Pipe struct {
	ChanIn  chan string
	ChanOut chan string
	List    []string
}

func New() *Pipe {
	p := &Pipe{
		ChanIn:  make(chan string, 100),
		ChanOut: make(chan string, 100),
		List:    []string{},
	}
	go p.Run()
	return p
}

func (pipe *Pipe) Push(s string) {
	pipe.ChanIn <- s
}

func (pipe *Pipe) Shift() string {
	s := <-pipe.ChanOut
	return s
}

func (pipe *Pipe) Run() {
	var nextStr string
	isRead := false

	maxLen := 0

	for {
		if isRead {
			select {
			case str, isGood := <-pipe.ChanIn:
				if !isGood {
					return
				}
				pipe.List = append(pipe.List, str)
				if maxLen < len(pipe.List) {
					maxLen = len(pipe.List)
				}
			case pipe.ChanOut <- nextStr:
				isRead = false
				if len(pipe.List) > 0 {
					nextStr = pipe.List[0]
					pipe.List = pipe.List[1:]
					isRead = true
				}
			}
		} else {
			select {
			case str, isGood := <-pipe.ChanIn:
				if !isGood {
					return
				}
				if isRead {
					pipe.List = append(pipe.List, str)
				} else {
					isRead = true
					nextStr = str
				}

			}
		}
	}
	close(pipe.ChanOut)
}
