package api

import (
	"fmt"
	"os"
	"os/signal"
	"pierflow/internal/business"
	"pierflow/internal/business/projects"
	"pierflow/internal/logger"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartApiServer(config *ServerConfig) error {

	logLevel, err := logFromString(config.Log)
	if err != nil {
		return err
	}

	// Initialize logger with the specified log level
	err = logger.InitLogLevel(logLevel)
	if err != nil {
		return err
	}

	basePath := checkBasePathAndCreatePathIfNecessary(config.BasePath)

	logger.Infof("Try to start Pierflow API server on http://%s:%d with base path '%s', with log level '%s' and database with '%s/%s'",
		config.Host, config.Port, basePath, config.Log, basePath, config.DbPath,
	)

	// DbProject manager
	pm, err := projects.NewProjectManager(basePath, config.DbPath)
	if err != nil {
		return fmt.Errorf("failed to initialize project manager: %w", err)
	}

	// System Manager
	sm := business.NewSystemManager(basePath)

	server := echo.New()

	server.HideBanner = true
	server.HidePort = true

	// Middlewares
	server.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogHost:       true,
			LogURI:        true,
			LogMethod:     true,
			LogStatus:     true,
			LogLatency:    true,
			LogError:      true,
			LogValuesFunc: loggingRequestFunc,
		}),
		middleware.Recover(),
	)

	// Static file serving
	if config.Web != nil {
		server.StaticFS("/", echo.MustSubFS(config.Web, "web"))
	}
	// Routing
	group := server.Group("/api")
	err = registerEndpoints(pm, sm, group)
	if err != nil {
		return fmt.Errorf("failed to registerEndpoints API endpoints: %w", err)
	}
	
	err = pm.ListenForComposeEvents()
	if err != nil {
		return fmt.Errorf("failed to start listening for compose events: %w", err)
	}

	// Listening for exit signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go listenForExit(quit)

	logger.Infof("Start succesfull Pierflow API server on http://%s:%d with base path '%s', with log level '%s' and database with '%s/%s'",
		config.Host, config.Port, config.BasePath, config.Log, basePath, config.DbPath,
	)

	return server.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}
