package app

import (
	"bmstuInformaticsTechnologies/internal/category_service"
	"bmstuInformaticsTechnologies/internal/configuration"
	"bmstuInformaticsTechnologies/internal/handlers/admin"
	"bmstuInformaticsTechnologies/internal/handlers/auth"
	"bmstuInformaticsTechnologies/internal/handlers/categories"
	"bmstuInformaticsTechnologies/internal/handlers/products"
	"bmstuInformaticsTechnologies/internal/product_service"
	"bmstuInformaticsTechnologies/internal/user_service"
	"bmstuInformaticsTechnologies/pkg/client/postrgresql"
	"bmstuInformaticsTechnologies/pkg/handlers/api_docs"
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

// @title Product Service
// @version 1.0.0
// @BasePath /

const (
	BaseApiPath = "/api"
)

// Application struct is main struct containing services and needs
type Application struct {
	logger logging.LoggerInterface
	router *mux.Router
	server *http.Server
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
	categoryService := category_service.NewCategoryService(app.logger, dbClient)
	userService := user_service.NewUserService(app.logger, dbClient)

	// Starting handlers registration
	pingHandler := ping.Handler{}
	pingHandler.Register(app.router)

	staticHandler := static.Handler{}
	staticHandler.Register(app.router)

	productsHandler := products.NewProductHandler(logger, productService)
	productsHandler.Register(app.router)

	adminHandler := admin.NewAdminHandler(logger, productService)
	adminHandler.Register(app.router)

	categoriesHandler := categories.NewCategoriesHandler(logger, categoryService)
	categoriesHandler.Register(app.router)

	documentationHandler := api_docs.Handler{}
	documentationHandler.Register(app.router)

	userHandler := auth.NewAuthHandler(logger, userService)
	userHandler.Register(app.router)

	app.router.NotFoundHandler = notfound.NotFoundHandler()

	return &app, nil
}

// Run is an entry point to start our application
func (app *Application) Run() {
	// Restarting application on panic

	app.logger.Info("Successfully initialized, starting server")
	app.logger.Fatal(app.server.ListenAndServe().Error())
}
