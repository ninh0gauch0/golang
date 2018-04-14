package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init the configuration needed to start the server
func (s *Server) Init() bool {
	s.router = mux.NewRouter()

	//Reading configuration file
	dat, err := ioutil.ReadFile("jsons/mongoconfig.json")
	check(err)
	var result MongoConf
	check(json.Unmarshal(dat, &result))

	s.initialized = true
	return true
}

// Start the server
func (s *Server) Start(config map[string]string) chan bool {

	// Recovering config server
	addr, ok := config["addr"]
	if !ok {
		addr = ":8080"
	}
	s.Addr = addr

	if s.initialized != true {
		err := s.Init()
		if err {
			return nil
		}
	}

	s.logger.Infof("Starting server....")
	s.worker = &Worker{}
	s.worker.Init(s.Ctx, s.GetLogger())

	s.addRoutes()

	exitChan := make(chan bool)

	// Go routines and channel to orchestate
	go func() {
		<-exitChan
		s.logger.Infoln("Stopping server")
		// Server shutdown
		err := s.Server.Shutdown(s.Ctx)

		if err != nil {
			s.logger.Errorln("Error shutdowning server - error: ", err.Error())
		}
	}()
	go func() {
		log.Printf("Listening on.... %s", s.Addr)
		log.Fatal(http.ListenAndServe(addr, s.router))
	}()

	return exitChan
}

// addRoutes - Define API routes
func (s *Server) addRoutes() {
	s.router.HandleFunc("/hms/recipes/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		recipe := s.worker.GetRecipeByID(id)
		s.logger.Infoln("Receipe returned: %s", recipe.Name)
		// returns
		fmt.Fprintf(w, "You've requested the recipe: %s\n", id)
	}).Methods("GET").Schemes("http")

	s.router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		/*
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
		*/
	})
}

//TODO: Change error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}
