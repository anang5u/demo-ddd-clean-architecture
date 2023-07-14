package loans

import (
	"demo-ddd-clean-architecture/app/services/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

const configurationFile = "../../../.env.test"

var dbMock = common.Apply().DbMock

func TestLoans_loanService(t *testing.T) {
	service := newLoanService(
		withService(),
		withRepository(dbMock.GetDb()),
	)

	assert.NotNil(t, service)
}
