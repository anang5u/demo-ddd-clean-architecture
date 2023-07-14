package common

import (
	"fmt"
	"log"
)

type cmnLogger struct {
	conf ConfigRepository
}

func NewLogger(config ConfigRepository) *cmnLogger {
	return &cmnLogger{
		conf: config,
	}
}

// Debug
func (c *cmnLogger) Debug(s ...interface{}) {
	log.Printf("[DEBUG] %s", c.toString(s))
}

// Info
func (c *cmnLogger) Info(s ...interface{}) {
	log.Printf("[INFO] %s", c.toString(s))
}

// Warn
func (c *cmnLogger) Warn(s ...interface{}) {
	log.Printf("[WARNING] %s", c.toString(s))
}

// Error
func (c *cmnLogger) Error(s ...interface{}) {
	log.Printf("[ERROR] %s", c.toString(s))
}

func (c *cmnLogger) toString(s ...interface{}) string {
	result := ""
	if len(s) == 0 {
		return result
	}
	for _, str := range s {
		result = fmt.Sprintf("%s%s", result, str)
	}

	return result
}
