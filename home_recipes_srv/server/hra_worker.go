package server

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Init - Starts the worker
func (w *Worker) Init(ctx context.Context, logger *log.Entry) {
	w.SetLogger(logger)
	w.Ctx = ctx
}

// GetRecipeByID - Given an id, returns a recipe
func (w *Worker) GetRecipeByID(id string) HRAResponse {
	w.logger.Debugf("Worker - getRecipebyId [IN]")
	rsp := HRAResponse{}

	dummyRecipe := Recipe{}

	if dummyRecipe.Name == "" {
		funcError := FunctionalError{}
		funcError.SetError(fmt.Errorf("Bad initialization"))

		//  everything is ok if we try to assign a value of type *FunctionalError to HRSError
		rsp.Error = &funcError
	}

	rsp.Status = Status{
		Code:        "200",
		Description: "Query completed",
	}

	rsp.Status.getObjectInfo()
	w.logger.Debugf("Worker - getRecipebyId [OUT]")
	return rsp
}
