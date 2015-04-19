package WebServer

import (
	"Web/Handlers/NotFound"
	"Web/Handlers/Stat"
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

type Server struct {
	_pat  *pat.PatternServeMux
	port  int
	debug bool
}

func (server *Server) Listen(WG *sync.WaitGroup) {

	WG.Add(1)
	go func() {
		log.Printf("%s", http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil))
		WG.Done()
	}()

	if server.debug {
		log.Printf("Web server listens on :%d\n", server.port)
	}
}

func New(port int, idDebug ...bool) (*Server, error) {

	debug := false
	if len(idDebug) > 0 {
		debug = idDebug[0]
	}

	_pat := pat.New()

	server := &Server{
		_pat:  _pat,
		debug: debug,
		port:  port,
	}

	http.Handle("/", _pat)

	GETS := map[string]interface{}{
		"/stats": Stat.H,
		"/404":   NotFound.H,
	}
	server.init_handlers(_pat, GETS)

	return server, nil
}

func (server *Server) init_handlers(m *pat.PatternServeMux, list map[string]interface{}) {
	for path, fun := range list {
		if server.debug {
			log.Printf("Web server reads uri: %s\n", path)
		}
		m.Get(path, http.HandlerFunc(fun.(func(http.ResponseWriter, *http.Request))))
		m.Get(path+"/", http.HandlerFunc(fun.(func(http.ResponseWriter, *http.Request))))
	}
}
