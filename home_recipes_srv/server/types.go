package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server struct
type Server struct {
	LoggerTrait
	Server      *http.Server
	Addr        string
	router      *mux.Router
	Ctx         context.Context
	worker      *Worker
	initialized bool
}

// Worker struct
type Worker struct {
	LoggerTrait
	Ctx context.Context
}

/** Logger **/

// LoggerTrait - a logger trait that let's you configure a log
type LoggerTrait struct {
	logger *log.Entry
}

// SetLogger - let's you configure a logger
func (lt *LoggerTrait) SetLogger(l *log.Entry) {
	if l != nil {
		lt.logger = l
	}
}

// GetLogger - returns the logger
func (lt *LoggerTrait) GetLogger() *log.Entry {
	return lt.logger
}

/** DTOs **/

// MongoConf DTO
type MongoConf struct {
	host string
	port string
}

// GetHost - returns the host
func (mc *MongoConf) GetHost() string {
	return mc.host
}

// GetPort - returns the port
func (mc *MongoConf) GetPort() string {
	return mc.port
}

// HRAResponse DTO
type HRAResponse struct {
	status Status `json:"status"`
	recipe Recipe `json:"recipe"`
}

/* Errors */

// err       error
type NorgateMathError struct {
	lat, long string
	err       error
}

/*  // EXAMPLE

type DeployRequest struct {
	ID          string `json:"_id"`
	Ns          string `json:"ns"`
	Description string `json:"description"`
	Deploy      struct {
		DocTypes []struct {
			DocTypeID string `json:"docTypeId"`
		} `json:"docTypes"`
		Document []struct {
			Name     string `json:"name"`
			DocTypes []struct {
				DocTypeID string `json:"docTypeId"`
			} `json:"docTypes"`
		} `json:"document"`
	} `json:"deploy"`
}
*/
