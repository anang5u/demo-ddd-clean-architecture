package identityrules

import "gorm.io/gorm"

type IdentityRuleRepository interface {
	WithDbConn(db *gorm.DB) *identityRules
	GetContractNumberCharacterLength() int
	GenerateContractNumber(length ...int) (string, error)
}
