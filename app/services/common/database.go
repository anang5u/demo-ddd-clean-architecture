package common

import (
	"database/sql"
	"demo-ddd-clean-architecture/app/exception"
	"demo-ddd-clean-architecture/app/migration"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type database struct {
	conf   ConfigRepository
	logger LoggerRepository

	sqlDb *sql.DB
	db    *gorm.DB
}

func NewDatabase(config ConfigRepository, logger LoggerRepository) *database {
	return &database{
		conf:   config,
		logger: logger,
	}
}

// connect
func (m *database) connect() (err error) {
	if m.sqlDb == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			m.conf.Get("DB_USER"),
			m.conf.Get("DB_PASSWORD"),
			m.conf.Get("DB_HOST"),
			m.conf.Get("DB_PORT"),
			m.conf.Get("DB_NAME"),
		)
		m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		m.db, err = gorm.Open(mysql.New(mysql.Config{
			Conn: m.sqlDb,
		}), &gorm.Config{})
	}

	return
}

// Close Connection
func (m *database) Close(db *gorm.DB) {
	dbInstance, _ := db.DB()
	_ = dbInstance.Close()
}

// WithSqlDb
func (m *database) WithSqlDb(sqlDb *sql.DB) *database {
	m.sqlDb = sqlDb
	return m
}

// GetDbConn
func (m *database) GetDbConn() *gorm.DB {
	if m.db == nil {
		if err := m.connect(); err != nil {
			m.logger.Error(err)
		}
	}
	return m.db
}

// AutoMigrate
func (m *database) AutoMigrate() {
	db := m.GetDbConn()
	if db == nil {
		m.logger.Error("Error while establishing database connection!")
		return
	}
	defer m.Close(db)

	if len(migration.ModelMigrations) > 0 {
		start := time.Now()
		err := db.AutoMigrate(migration.ModelMigrations...)
		exception.PanicIfNeeded(err)

		m.logger.Info(fmt.Sprintf("MIGRATE FINISH IN : %s", time.Since(start)))

		seeds := migration.DataSeeds()
		if len(seeds) > 0 {
			for i := range seeds {
				tx := db.Begin()

				defer func() {
					if r := recover(); r != nil {
						tx.Rollback()
					}
				}()

				if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seeds[i]).Error; nil != err {
					m.logger.Error(fmt.Sprintf("Seeds Error: %s", err.Error()))
					tx.Rollback()
				}

				if err := tx.Commit().Error; nil != err {
					m.logger.Error(fmt.Sprintf("Seeds Error: %s", err.Error()))
					tx.Rollback()
				}
			}
		}
	}
}
