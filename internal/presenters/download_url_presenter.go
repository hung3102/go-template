package presenters

import (
	"gcim/example/internal/api"
	"gcim/example/internal/usecases/dto/output"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IDownloadUrlPresenter interface {
	PresentGetDownloadUrl(c echo.Context, output *output.GetDownloadUrlOutput) error
}

type DownloadUrlPresenter struct{}

func NewDownloadUrlPresenter() IDownloadUrlPresenter {
	return &DownloadUrlPresenter{}
}

func (p *DownloadUrlPresenter) PresentGetDownloadUrl(c echo.Context, output *output.GetDownloadUrlOutput) error {
	response := api.GetDownloadUrlResponse{
		Url: &output.Url,
	}

	return c.JSON(http.StatusOK, response)
}
