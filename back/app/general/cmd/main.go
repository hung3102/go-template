package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/54m/echo-routing/output"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	"github.com/topgate/gcim-temporary/back/app/general/internal/config"
	"github.com/topgate/gcim-temporary/back/app/general/internal/controllers"
	"github.com/topgate/gcim-temporary/back/app/general/internal/initialize"
	"github.com/topgate/gcim-temporary/back/app/general/internal/interfaces/openapi"
	authenticationMiddleware "github.com/topgate/gcim-temporary/back/app/internal/middleware/authentication"
	"github.com/topgate/gcim-temporary/back/app/internal/middleware/logger"
	"github.com/topgate/gcim-temporary/back/app/internal/middleware/validator"
	"github.com/topgate/gcim-temporary/back/pkg/tests"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	internalPort = func() string {
		p := os.Getenv("PORT")
		if p == "" {
			p = "1234"
		}
		return p
	}()
	externalPort = "27001"
)

func main() {
	tests.IsTest = false

	ctx := context.Background()
	e := echo.New()

	useLoggerMiddleware(e)

	cfg, err := config.ReadConfig(ctx)
	if err != nil {
		log.Fatalf("failed to read config: %+v", err)
	}

	useCORSMiddleware(e, cfg.CORSAllowOrigins)

	externalDeps, err := initialize.NewExternalDependencies(ctx, *cfg)
	if err != nil {
		log.Fatalf("failed to initialize external dependencies: %+v", err)
	}
	useRequestValidator(e, *externalDeps)
	e.HTTPErrorHandler = initialize.ErrorHandler()

	useCaseDependencies := initialize.NewUseCaseDependencies(*cfg, *externalDeps)
	useCases := initialize.NewUseCases(*cfg, *useCaseDependencies)
	cp := initialize.NewControllerProps(cfg, *useCases)

	useAuthenticationMiddleware(e, *cfg, *useCaseDependencies)

	strictServer := controllers.NewController(cp)
	si := openapi.NewStrictHandler(strictServer, []openapi.StrictMiddlewareFunc{})
	openapi.RegisterHandlersWithBaseURL(e, si, "/api")

	serverAddr := ":" + internalPort

	go func() {
		// Wait for interrupt signal to gracefully shut down the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	output.Do(e.Routes())

	fmt.Println("Accessible URL↓")
	fmt.Printf("http://localhost:%s/api/health\n", externalPort)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(serverAddr))
}

func useLoggerMiddleware(e *echo.Echo) {
	e.Use(echoMiddleware.Recover())

	stdoutCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		zapcore.DebugLevel,
	)
	stderrCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stderr),
		zapcore.ErrorLevel,
	)
	zapLogger := zap.New(zapcore.NewTee(stdoutCore, stderrCore))

	e.Use(echozap.ZapLogger(zapLogger))
	e.Use(logger.InjectionMiddleware(zapLogger))
}

func useCORSMiddleware(e *echo.Echo, corsAllowOrigins string) {
	e.Use(
		echoMiddleware.CORSWithConfig(
			echoMiddleware.CORSConfig{
				AllowOrigins: lo.Uniq(strings.Split(corsAllowOrigins, ",")),
				AllowHeaders: []string{
					echo.HeaderOrigin,
					echo.HeaderXRequestedWith,
					echo.HeaderContentType,
					echo.HeaderAccept,
					echo.HeaderCookie,
					echo.HeaderAuthorization,
				},
				AllowMethods: []string{
					http.MethodGet,
					http.MethodPut,
					http.MethodPost,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodOptions,
				},
				AllowCredentials: true,
			},
		),
	)
}

func useRequestValidator(e *echo.Echo, dep initialize.ExternalDependencies) {
	ovm, err := validator.NewMiddleware(dep.OpenAPISpec())
	if err != nil {
		log.Fatalf("failed to create validator middleware: %+v", err)
	}

	e.Use(ovm.Middleware)
}

func useAuthenticationMiddleware(e *echo.Echo, cfg config.Config, usecaseDependencies initialize.UseCaseDependencies) {
	param := authenticationMiddleware.NewAuthenticatorParam{
		JWTSecret:         []byte(cfg.JWTSecret),
		CookieSessionName: "gcim_api_session",
		SkipPaths: []string{
			"/api/authentication/sign-in",
			"/api/health",
		},
		ContextKey:            "user",
		AuthenticationService: usecaseDependencies.AuthenticationService,
		UserSessionRepository: usecaseDependencies.SessionRepository,
	}
	e.Use(authenticationMiddleware.NewAuthenticator(param))
}
