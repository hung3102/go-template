package controllers

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/general/internal/interfaces/openapi"
)

// GetHealth - サーバー生存確認
// (GET /health)
func (c *Controller) GetHealth(context.Context, openapi.GetHealthRequestObject) (openapi.GetHealthResponseObject, error) {
	res := openapi.GetHealth200JSONResponse{
		Status: "ok",
	}

	return res, nil
}
