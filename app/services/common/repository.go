package common

import (
	"database/sql"

	"gorm.io/gorm"
)

type (
	// Config
	ConfigRepository interface {
		Get(key string) string
	}

	// Logger
	LoggerRepository interface {
		Debug(s interface{})
		Info(s interface{})
		Warn(s interface{})
		Error(s interface{})
	}

	// Database
	DatabaseRepository interface {
		GetDbConn() *gorm.DB
		WithSqlDb(sqlDb *sql.DB) *database
		AutoMigrate()
	}
)
