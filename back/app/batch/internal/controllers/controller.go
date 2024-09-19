package controllers

import (
	"github.com/topgate/gcim-temporary/back/app/batch/internal/interfaces/openapi"
	"github.com/topgate/gcim-temporary/back/app/batch/internal/interfaces/props"
)

// Controller - server
type Controller struct {
	props *props.ControllerProps
}

// NewController - constructor
func NewController(cp *props.ControllerProps) openapi.StrictServerInterface {
	return &Controller{
		props: cp,
	}
}
