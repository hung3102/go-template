package main

import (
	"gcim/example/config"
	"gcim/example/pkg/getdownloadurlexample"
	uploadexample "gcim/example/pkg/uploadExample"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Load .env file
	config.LoadEnv()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/upload-sample/:eventId/:orgCspDocId", uploadexample.UploadExample)
	e.GET("/get-download-url", getdownloadurlexample.GetDownloadURLExample)
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":1313"))
}
