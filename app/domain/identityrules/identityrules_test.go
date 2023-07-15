package identityrules

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var identityRule = NewIdentityRules()

func TestIdentityRules_GenerateContractNumber(t *testing.T) {
	length := identityRule.GetContractNumberCharacterLength()
	contractNumber, err := identityRule.GenerateContractNumber(length)

	assert.Nil(t, err)
	assert.Equal(t, len(contractNumber) > 0, true)
}

func TestIdentityRules_CreateMask(t *testing.T) {
	res := identityRule.CreateMask("Budi")

	log.Println(res)
}
