package transaction

import (
	"demo-ddd-clean-architecture/app/services/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbMock = common.Apply("../../../.env.test").DbMock

func Test_trxService(t *testing.T) {
	service := newTrxService(
		withService(),
		withRepository(),
	).WithDb(dbMock.GetDb())

	assert.NotNil(t, service)
}
