package common

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

type (
	// Config
	ConfigRepository interface {
		Get(key string) string
	}

	// Logger
	LoggerRepository interface {
		Debug(s ...interface{})
		Info(s ...interface{})
		Warn(s ...interface{})
		Error(s ...interface{})
	}

	// Database
	DatabaseRepository interface {
		GetDbConn() *gorm.DB
		WithSqlDb(sqlDb *sql.DB) *database

		// Transaction
		TxBegin() *gorm.DB
		TxRecover(tx *gorm.DB) func()
		TxRollback(tx *gorm.DB)
		TxCommit(tx *gorm.DB) error

		AutoMigrate()
	}

	// Database Mock (mysql)
	DatabaseMockRepository interface {
		GetDb() *gorm.DB
		Close()
		GetMock() sqlmock.Sqlmock
	}
)
