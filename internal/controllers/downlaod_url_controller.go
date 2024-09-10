package controllers

import (
	"gcim/example/internal/api"
	"gcim/example/internal/presenters"
	"gcim/example/internal/usecases"
	"gcim/example/internal/usecases/dto/input"

	"github.com/labstack/echo/v4"
)

type DownloadUrlController struct {
	getDownloadUrlUsecase usecases.IGetDownloadUrlUsecase
	downloadUrlPresenter  presenters.IDownloadUrlPresenter
	errorPresenter        presenters.IErrorPresenter
}

func NewDownloadUrlController(
	getDownloadUrlUsecase usecases.IGetDownloadUrlUsecase,
	downloadUrlPresenter presenters.IDownloadUrlPresenter,
	errorPresenter presenters.IErrorPresenter,

) *DownloadUrlController {
	return &DownloadUrlController{
		getDownloadUrlUsecase: getDownloadUrlUsecase,
		downloadUrlPresenter:  downloadUrlPresenter,
		errorPresenter:        errorPresenter,
	}
}

func (c *DownloadUrlController) GetDownloadUrl(e echo.Context, params api.GetDownloadUrlParams) error {
	input := &input.GetDownloadUrlInput{
		Path: params.Path,
	}
	ctx := e.Request().Context()

	output, err := c.getDownloadUrlUsecase.Execute(ctx, input)
	if err != nil {
		return c.errorPresenter.PresentInternalServerError(e, err)
	}

	return c.downloadUrlPresenter.PresentGetDownloadUrl(e, output)
}
