package common

import (
	"demo-ddd-clean-architecture/app/exception"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
}

func NewConfig(filenames ...string) *config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &config{}
}

// Get
func (c *config) Get(key string) string {
	return os.Getenv(key)
}
