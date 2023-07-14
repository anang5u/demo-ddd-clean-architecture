package common

import (
	"database/sql"
	"demo-ddd-clean-architecture/app/exception"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbMock struct {
	db     *gorm.DB
	mockDb *sql.DB
	mock   sqlmock.Sqlmock
}

func NewDbMock() *dbMock {
	mockDb, mock, err := sqlmock.New()
	exception.PanicIfNeeded(err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	exception.PanicIfNeeded(err)

	return &dbMock{
		db:     db,
		mockDb: mockDb,
		mock:   mock,
	}
}

// GetDb
func (m *dbMock) GetDb() *gorm.DB {
	return m.db
}

// GetMock
func (m *dbMock) GetMock() sqlmock.Sqlmock {
	return m.mock
}

// Close
func (m *dbMock) Close() {
	dbInstance, _ := m.db.DB()
	_ = dbInstance.Close()

	m.mockDb.Close()
}
