package presenters

import (
	"net/http"
	"template/example/internal/api"
	"template/example/internal/usecases/dto/output"

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
