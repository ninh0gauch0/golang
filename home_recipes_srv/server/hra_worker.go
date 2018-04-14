package server

import (
	"context"

	log "github.com/sirupsen/logrus"
)

// Init - Starts the worker
func (w *Worker) Init(ctx context.Context, logger *log.Entry) {
	w.SetLogger(logger)
	w.Ctx = ctx
}

// GetRecipeByID - Given an id, returns a recipe
func (w *Worker) GetRecipeByID(id string) Recipe {
	w.logger.Debugf("Worker - getRecipebyId [IN]")
	rsp := Recipe{Name: "dummy"}
	w.logger.Debugf("Worker - getRecipebyId [OUT]")
	return rsp
}
