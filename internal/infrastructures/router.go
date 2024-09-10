package infrastructure

import (
	"gcim/example/internal/api"
	"gcim/example/internal/controllers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userController *controllers.TaskController
}

func NewServer(
	userController *controllers.TaskController,
) *Server {
	return &Server{
		userController: userController,
	}
}

func (s *Server) CreateTask(ctx echo.Context) error {
	return s.userController.CreateTask(ctx)
}

func (s *Server) GetTask(ctx echo.Context, taskId string) error {
	return s.userController.GetTask(ctx, taskId)
}

func (s *Server) UploadExample(ctx echo.Context, eventId string, orgCspDocId string) error {
	return nil
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

	e.Logger.Fatal(e.Start(":1313"))
}
