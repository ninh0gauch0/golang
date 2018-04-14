package main

import (
	log "github.com/sirupsen/logrus"
)

const (
	version = "1.0.0"
)

func init() {
	logger := log.StandardLogger()
	logger.Formatter = &log.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	}
}

func main() {

}
