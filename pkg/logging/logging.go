package logging

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
	"time"
)

// Logger using a global logger interface
// Choosing the global logger depends on using a lot of concurrency structures
var logger *zap.Logger

type Logger struct {
	zap.Config `yaml:"logging"`
}

// GetLogger is function returning current logger
// If no logger was initialized panics
func GetLogger() LoggerInterface {
	if logger != nil {
		return logger
	}
	panic("nil logger access, be aware of nil pointers difference")
}

// Init getting Logger configuration
func Init(file io.Reader) error {
	cfg := Logger{}
	yamlDecoder := yaml.NewDecoder(file)

	if err := yamlDecoder.Decode(&cfg); err != nil {
		return fmt.Errorf("unkown yaml configuration recieved: %v", err)
	}

	// Making run-time logs on path specified
	// If there is  output paths, that are not stdout or stderr, those will be created with date name
	// You can either select between .log file or ./path/file

	for i, outputPath := range cfg.OutputPaths {
		if strings.HasPrefix(outputPath, "./") && !strings.HasSuffix(outputPath, ".log") {
			if _, err := os.Stat(outputPath); !os.IsExist(err) {
				err = os.MkdirAll(outputPath, 0750)
				if err != nil {
					return err
				}
			}
			now := time.Now()
			filePath := outputPath +
				fmt.Sprintf("%d_%02d_%02d_%d-%d-%d", now.Year(),
					now.Month(), now.Day(), now.Hour(), now.Minute(),
					now.Minute()) + ".log"
			f, err := os.Create(filePath)
			if err != nil {
				return err
			}
			fmt.Println("Created log file: ", filePath)
			cfg.OutputPaths[i] = filePath
			f.Close()
		} else if strings.HasSuffix(outputPath, ".log") {
			// Making sure that the file was created, if not create it
			if _, err := os.Stat(outputPath); !os.IsExist(err) {
				// Making sure that /path/to/log.log was already created, creating if not
				splitedOutPutPath := strings.Split(outputPath, "/")
				if _, err := os.Stat(outputPath); len(splitedOutPutPath) >= 2 && !os.IsExist(err) {
					err = os.MkdirAll(strings.Join(splitedOutPutPath[:len(splitedOutPutPath)-1], "/"), 0750)
				}
				f, err := os.Create(outputPath)
				if err != nil {
					return err
				}
				f.Close()
			}
		}
	}

	// Building logger with configuration information
	var err error
	logger, err = cfg.Build()
	if err != nil {
		return fmt.Errorf("unable to build logger: %s", err)
	}
	defer logger.Sync()
	logger.Info("Successfully initialized logger")

	return nil
}
