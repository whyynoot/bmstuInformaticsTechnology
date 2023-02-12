package app

import (
	"bmstuInformaticsTechnologies/internal/configuration"
	"bmstuInformaticsTechnologies/internal/handlers/admin"
	"bmstuInformaticsTechnologies/internal/handlers/products"
	"bmstuInformaticsTechnologies/internal/product_service"
	"bmstuInformaticsTechnologies/pkg/client/postrgresql"
	"bmstuInformaticsTechnologies/pkg/handlers/notfound"
	"bmstuInformaticsTechnologies/pkg/handlers/ping"
	"bmstuInformaticsTechnologies/pkg/handlers/static"
	"bmstuInformaticsTechnologies/pkg/logging"
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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
		Addr:         ":" + strconv.Itoa(cfg.Server.Port),
		WriteTimeout: time.Duration(cfg.Server.Timeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.Timeout) * time.Second,
	}

	// Initializing services
	dbClient, err := postrgresql.NewClient(context.Background(), cfg.DataBaseURL)
	if err != nil {
		app.logger.Fatal("unable to connect to database", zap.String("error", err.Error()))
	}

	productService := product_service.NewProductService(app.logger, dbClient)

	// Starting handlers registration
	pingHandler := ping.Handler{}
	pingHandler.Register(app.router)

	staticHandler := static.Handler{}
	staticHandler.Register(app.router)

	productsHandler := products.NewProductHandler(logger, productService)
	productsHandler.Register(app.router)

	adminHandler := admin.NewAdminHandler(logger, productService)
	adminHandler.Register(app.router)

	app.router.NotFoundHandler = notfound.NotFoundHandler()

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
	app.logger.Fatal(app.server.ListenAndServe().Error())
}
