package infrastructure

import (
	"gcim/example/internal/api"
	"gcim/example/internal/controllers"
	uploadexample "gcim/example/pkg/uploadExample"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userController *controllers.TaskController
	getDownloadUrl *controllers.DownloadUrlController
	uploadExample  *uploadexample.UploadExample
}

func NewServer(
	userController *controllers.TaskController,
	getDownloadUrl *controllers.DownloadUrlController,
	uploadExample *uploadexample.UploadExample,
) *Server {
	return &Server{
		userController: userController,
		getDownloadUrl: getDownloadUrl,
		uploadExample:  uploadExample,
	}
}

func (s *Server) CreateTask(ctx echo.Context) error {
	return s.userController.CreateTask(ctx)
}

func (s *Server) GetTask(ctx echo.Context, taskId string) error {
	return s.userController.GetTask(ctx, taskId)
}

func (s *Server) GetDownloadUrl(ctx echo.Context, params api.GetDownloadUrlParams) error {
	return s.getDownloadUrl.GetDownloadUrl(ctx, params)
}

func (s *Server) UploadExample(ctx echo.Context, eventId string, orgCspDocId string) error {
	return s.uploadExample.Run(ctx, eventId, orgCspDocId)
}

func InitRouter() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	container := BuildContainer()

	var server *Server
	if err := container.Invoke(func(s *Server) {
		server = s
	}); err != nil {
		log.Fatalf("Error resolving dependencies: %v", err)
	}

	api.RegisterHandlers(e, server)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1313"))
}
