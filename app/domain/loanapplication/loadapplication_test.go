package loanapplication

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

var modLoans = NewLoanApplication()

var customerId = "71e11445-8f94-493d-bd19-d7e43a1e576c"
var customer = model.Customer{
	Base: model.Base{
		Id: helper.UuidMustParse(customerId),
	},
}

func TestInstallment_IsApprovedExists(t *testing.T) {
	db := dbMock.GetDb()
	defer dbMock.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	isExists, err := modLoans.WithDbConn(db).IsApprovedExists(customer.Id)

	assert.Nil(t, err)
	assert.Equal(t, isExists == true, true)
}

func TestInstallment_CalculateTenor(t *testing.T) {
	limit, err := modLoans.CalculateTenor(&model.LoanApplication{
		Limit: 500000,
	})
	assert.Nil(t, err)
	assert.Equal(t, limit == 3, true)

	limit, err = modLoans.CalculateTenor(&model.LoanApplication{
		Limit: 2000000,
	})
	assert.Nil(t, err)
	assert.Equal(t, limit == 4, true)
}
