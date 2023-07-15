package identityrules

import (
	"demo-ddd-clean-architecture/app/helper"
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type identityRules struct {
	db *gorm.DB
}

// NewIdentityRules
func NewIdentityRules() *identityRules {
	return &identityRules{}
}

// WithDbConn
func (m *identityRules) WithDbConn(db *gorm.DB) *identityRules {
	m.db = db
	return m
}

// GetContractNumberCharacterLength
func (m *identityRules) GetContractNumberCharacterLength() int {
	return contractNumberCharacterLength
}

// GenerateContractNumber
func (m *identityRules) GenerateContractNumber(length ...int) (string, error) {
	var characterLength int

	if len(length) > 0 {
		characterLength = length[0]
	} else {
		characterLength = contractNumberCharacterLength
	}

	// Format YYMM+randomNumber
	return fmt.Sprintf("%s-%s", time.Now().Format("0601"), helper.RndNumeric((characterLength - 4))), nil
}

// CreateMask
func (m *identityRules) CreateMask(str string) string {
	result := ""

	arr := strings.Split(str, " ")
	for _, s := range arr {
		firstChar := s[0:1]
		lastChar := string(s[len(s)-1:])
		maskedChar := string(s[1:(len(s) - 1)])

		result = fmt.Sprintf("%s%s", result, firstChar)
		runes := []rune(maskedChar)
		for i := 0; i < len(runes); i++ {
			result = fmt.Sprintf("%s%s", result, "*")
		}
		result = fmt.Sprintf("%s%s ", result, lastChar)
	}

	return result
}

// GetCurrency
func (m *identityRules) GetCurrency() string {
	return currency
}

// GenerateToken
func (m *identityRules) GenerateToken(length ...int) string {
	var characterLength int

	if len(length) > 0 {
		characterLength = length[0]
	} else {
		characterLength = tokenCharacterLength
	}

	// Format YYMM+randomNumber
	return helper.RndNumeric(characterLength)
}

// MakeExpiredAtHours
func (m *identityRules) MakeExpiredAtHours(hours ...int) time.Time {
	// time.Now().Add(time.Hour * hours + time.Minute * mins + time.Second * sec)
	now := helper.GetTimeNow()
	h := expiredAtHours
	if len(hours) > 0 {
		h = hours[0]
	}
	hourDuration := time.Duration(h)

	return now.Add(time.Hour * hourDuration)
}

// MakeHash
func (m *identityRules) MakeHash(str string) string {
	bStr := []byte(str)
	hash, err := bcrypt.GenerateFromPassword(bStr, bcrypt.MinCost)
	if err != nil {
		return ""
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// CompareHash
func (m *identityRules) CompareHash(strHashed string, strPlain string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(strHashed)
	bytePlain := []byte(strPlain)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
