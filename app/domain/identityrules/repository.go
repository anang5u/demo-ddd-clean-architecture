package identityrules

import (
	"time"

	"gorm.io/gorm"
)

type IdentityRuleRepository interface {
	WithDbConn(db *gorm.DB) *identityRules
	GetContractNumberCharacterLength() int
	GenerateContractNumber(length ...int) (string, error)
	CreateMask(str string) string
	GetCurrency() string
	GenerateToken(length ...int) string
	MakeHash(str string) string
	CompareHash(strHashed string, strPlain string) bool
	MakeExpiredAtHours(hours ...int) time.Time
}
