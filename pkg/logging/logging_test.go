package logging_test

import (
	"defi/pkg/logging"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	configuration := strings.NewReader(`
logging:
  level: debug
  outputPaths:
    - ./test/logs/
    - ./test/logs/all.log
  encoding: json
  encoderConfig:
    messageKey: msg
    timeKey: ts
    timeEncoder: ISO8601
    callerEncoder: short
    callerKey: caller
    levelKey: lvl
    levelEncoder: capital`)

	now := time.Now()
	filePath := "./test/logs/" +
		fmt.Sprintf("%d_%02d_%02d_%d-%d-%d", now.Year(),
			now.Month(), now.Day(), now.Hour(), now.Minute(),
			now.Minute()) + ".log"

	err := logging.Init(configuration)
	assert.Nil(t, err)
	assert.NotNil(t, logging.GetLogger())

	assert.FileExists(t, filePath)
	defer os.RemoveAll(filePath)

	filePath = "./test/logs/all.log"

	assert.FileExists(t, filePath)
	defer os.RemoveAll(filePath)
}
