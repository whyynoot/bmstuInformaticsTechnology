package main

import (
	application "bmstuInformaticsTechnologies/internal/app"
	"bmstuInformaticsTechnologies/internal/configuration"
	"bmstuInformaticsTechnologies/pkg/logging"
	"go.uber.org/zap"
	"os"
)

const (
	LoggerConfigurationPath      = "./config/logger.yaml"
	ApplicationConfigurationPath = "./config/config.yaml"
)

func main() {
	// Getting logger file
	loggerConfigurationFile, err := os.Open(LoggerConfigurationPath)
	if err != nil {
		panic(err)
	}

	// Initializing logger configuration
	err = logging.Init(loggerConfigurationFile)
	if err != nil {
		panic(err)
	}
	// Getting logger
	logger := logging.GetLogger()
	logger.Info("Getting app configuration")

	// Initialization of configuration
	applicationConfigurationFile, err := os.Open(ApplicationConfigurationPath)
	if err != nil {
		logger.Fatal("Error opening configuration file:", zap.Error(err))
	}

	config, err := configuration.NewProgramConfig(applicationConfigurationFile)
	if err != nil {
		logger.Fatal("Program configuration initialization fatal error:", zap.Error(err))
	}

	// Initialization of app
	app, err := application.NewApplication(config, logger)
	if err != nil {
		logger.Fatal("Application initialization fatal error:", zap.Error(err))
	}

	app.Run()
}
