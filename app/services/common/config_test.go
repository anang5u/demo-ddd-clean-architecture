package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cfg = NewConfig(configurationFile)

func TestConfig_GetAppName(t *testing.T) {
	appName := cfg.Get("APP_NAME")
	assert.NotEmpty(t, appName)
}

func TestConfig_GetAppPort(t *testing.T) {
	appPort := cfg.Get("APP_PORT")
	assert.NotEmpty(t, appPort)
}
