package installment

import (
	"demo-ddd-clean-architecture/app/helper"
	"demo-ddd-clean-architecture/app/model"
	"demo-ddd-clean-architecture/app/services/common"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var dbMock = common.Apply().DbMock
var mock = dbMock.GetMock()

var modIns = NewInstallment()

var customerId = "71e11445-8f94-493d-bd19-d7e43a1e576c"
var customer = model.Customer{
	Base: model.Base{
		Id: helper.UuidMustParse(customerId),
	},
}

func TestInstallment_IsNewInstallmentExists(t *testing.T) {
	db := dbMock.GetDb()
	defer dbMock.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	isExists, err := modIns.WithDbConn(db).IsNewInstallmentExists(&customer)

	assert.Nil(t, err)
	assert.Equal(t, isExists == true, true)
}
