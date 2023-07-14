package common

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var gormDB = NewDatabase(cfg, logger)

func TestGorm_GetDbConn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %s", err)
	}
	defer db.Close()

	// a SELECT VERSION() query will be run when gorm opens the database
	// so we need to expect that here
	columns := []string{"version"}
	mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(
		mock.NewRows(columns).FromCSVString("1"),
	)
	gdb := NewDatabase(cfg, logger).WithSqlDb(db).GetDbConn()

	assert.NotNil(t, gdb)
}
