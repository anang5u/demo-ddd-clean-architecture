package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const configurationFile = "../../../.env.test"

func TestCommon_newCommonService(t *testing.T) {
	service := newCommonService(
		withConfigRepository(configurationFile),
		withRepository(),
	)

	assert.NotNil(t, service)
}
