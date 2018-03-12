package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Start is mine
func (s *Server) Start() (chan bool, error) {
	if s.initialized != true {
		if err := s.Init(); err != nil {
			return nil, err
		}
	}

	s.worker.start(s.Ctx)

	s.router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	s.router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {

		if s.worker.counter > 0 {
			s.logger.Infof("Result counter increment: %d", s.worker.counter)
			resp := []byte(strconv.Itoa(s.worker.counter))
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		} else {
			s.logger.Error("Increment counter error")
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("409 - Something bad happened!"))
		}

	})

	addr := fmt.Sprintf("%s:%s", s.Viper.GetString("ip"), s.Viper.GetString("port"))
	http.ListenAndServe(addr, s.router)

	return s.exitChan, nil
}

//Init the server
func (s *Server) Init() error {
	s.logger.Infof("Initializing the server")
	s.exitChan = make(chan bool)
	s.initialized = true

	s.router = mux.NewRouter()

	s.worker = &Worker{
		counterChan: make(chan bool),
	}
	s.worker.Period = 1000000000

	s.worker.SetLogger(s.logger)
	return nil
}

//Shutdown the server
func (s *Server) Shutdown(rootContext context.Context) {
	s.logger.Infof("Shutting down the server")
	ctx, cancel := context.WithTimeout(rootContext, 5*time.Second)
	defer cancel()
	s.Server.Shutdown(ctx)
}
