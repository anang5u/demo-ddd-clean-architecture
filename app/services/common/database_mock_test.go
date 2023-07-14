package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommon_DbMock(t *testing.T) {
	dbMock := NewDbMock()
	defer dbMock.Close()

	db := dbMock.GetDb()

	assert.NotNil(t, db)
}
