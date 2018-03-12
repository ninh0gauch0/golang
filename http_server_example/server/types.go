package server

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

//Server is mine
type Server struct {
	Server      *http.Server
	exitChan    chan bool
	initialized bool
	worker      *Worker
	Viper       *viper.Viper
	Ctx         *context.Context
	LoggerTrait
	router *mux.Router
}

//Worker is mine
type Worker struct {
	exitChan    chan bool
	counterChan chan bool
	LoggerTrait
	ticker  *time.Ticker
	status  string
	counter int
	Period  time.Duration
}

//LoggerTrait a logger trait that let's you configure a log
type LoggerTrait struct {
	logger *log.Logger
}

//SetLogger let's you configure a logger
func (lt *LoggerTrait) SetLogger(l *log.Logger) {
	if l != nil {
		lt.logger = l
	}
}
