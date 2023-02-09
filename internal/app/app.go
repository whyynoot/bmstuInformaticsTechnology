package app

import (
	"bmstuInformaticsTechnologies/internal/configuration"
	"bmstuInformaticsTechnologies/pkg/handlers/ping"
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// Application struct is main struct containing services and needs
type Application struct {
	logger logging.LoggerInterface
	router *mux.Router
	server *http.Server

	// api, http.Server + mux.Router ?
	// database, dbManager, may move to pkg
	//
}

// NewApplication is a constructor for our application based on configuration.Config
func NewApplication(cfg *configuration.Config, logger logging.LoggerInterface) (*Application, error) {
	app := Application{}

	app.logger = logger

	app.router = mux.NewRouter()

	app.server = &http.Server{
		Handler:      app.router,
		Addr:         ":8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	pingHandler := ping.Handler{}
	pingHandler.Register(app.router)

	return &app, nil
}

// Run is an entry point to start our application
func (app *Application) Run() {
	// Restarting application on panic
	defer func() {
		if err := recover(); err != nil {
			app.logger.Error("Fatal error, recovered from panic", zap.Any("error", err))
		}
	}()

	app.logger.Info("Successfully initialized, starting server")
	app.logger.Error(app.server.ListenAndServe().Error())
}
