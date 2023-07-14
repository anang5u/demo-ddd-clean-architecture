package loans

import (
	"demo-ddd-clean-architecture/app/domain/identityrules"
	"demo-ddd-clean-architecture/app/domain/installment"
	"demo-ddd-clean-architecture/app/domain/loanapplication"
	"demo-ddd-clean-architecture/app/domain/pricing"
	"demo-ddd-clean-architecture/app/services/common"
	"sync"

	"gorm.io/gorm"
)

var loanService *LoanService

type LoanService struct {
	Common       *common.CommonService
	Loans        loanapplication.LoanApplicationRepository
	Installment  installment.InstallmentRepository
	IdentityRule identityrules.IdentityRuleRepository
	Pricing      pricing.PricingRepository

	sync.Mutex
}

// LoanServiceConfiguration is an alias for a function that will take in a pointer to an LoanService and modify it
type LoanServiceConfiguration func(conf *LoanService) error

func newLoanService(cfgs ...LoanServiceConfiguration) *LoanService {
	// Create the NewLoanService
	conf := &LoanService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the config into the configuration function
		err := cfg(conf)
		if err != nil {
			return nil
		}
	}
	return conf
}

// withService applies a given repositories to the LoanService
func withService() LoanServiceConfiguration {
	// return a function that matches the LoanServiceConfiguration signature
	return func(s *LoanService) error {
		s.Common = common.Apply() // with common service
		return nil
	}
}

// withRepository applies a given repositories to the LoanService
func withRepository(db *gorm.DB) LoanServiceConfiguration {
	// return a function that matches the LoanServiceConfiguration signature
	return func(s *LoanService) error {
		s.Installment = installment.NewInstallment().WithDbConn(db)
		s.Loans = loanapplication.NewLoanApplication().WithDbConn(db)
		s.IdentityRule = identityrules.NewIdentityRules()
		s.Pricing = pricing.NewPricing(db)
		return nil
	}
}

// Apply Common Service
func Apply(db *gorm.DB) *LoanService {
	if loanService == nil {
		s := newLoanService(
			withService(),
			withRepository(db),
		)
		loanService = s
	}
	return loanService
}
