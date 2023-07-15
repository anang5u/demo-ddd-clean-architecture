package common

import (
	"demo-ddd-clean-architecture/app/config"
	"demo-ddd-clean-architecture/app/exception"
	"os"

	"github.com/joho/godotenv"
)

type modConfig struct {
}

func NewConfig(filenames ...string) *modConfig {
	if len(filenames) > 0 {
		err := godotenv.Load(filenames...)
		exception.PanicIfNeeded(err)
	} else {
		err := godotenv.Load(".env")
		exception.PanicIfNeeded(err)
	}

	return &modConfig{}
}

// Get
func (c *modConfig) Get(key string) string {
	cfgValue := ""

	// default config diambil dari config/environment.go
	if cfgConfigValue, ok := config.Environment[key]; ok {
		cfgValue = cfgConfigValue.(string)
	}

	// override form env config
	cfgEnvValue := os.Getenv(key)
	if len(cfgEnvValue) > 0 {
		cfgValue = cfgEnvValue
	}

	return cfgValue
}
