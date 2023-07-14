package common

import "log"

type cmnLogger struct {
	conf ConfigRepository
}

func NewLogger(config ConfigRepository) *cmnLogger {
	return &cmnLogger{
		conf: config,
	}
}

// Debug
func (c *cmnLogger) Debug(s interface{}) {
	log.Printf("[DEBUG] %s", s)
}

// Info
func (c *cmnLogger) Info(s interface{}) {
	log.Printf("[INFO] %s", s)
}

// Warn
func (c *cmnLogger) Warn(s interface{}) {
	log.Printf("[WARNING] %s", s)
}

// Error
func (c *cmnLogger) Error(s interface{}) {
	log.Printf("[ERROR] %s", s)
}
