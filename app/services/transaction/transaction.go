package transaction

import (
	"demo-ddd-clean-architecture/app/domain/identityrules"
	"demo-ddd-clean-architecture/app/domain/installment"
	"demo-ddd-clean-architecture/app/domain/pricing"
	"demo-ddd-clean-architecture/app/domain/transaction"
	"demo-ddd-clean-architecture/app/services/common"
	"demo-ddd-clean-architecture/app/services/loans"
	"sync"

	"gorm.io/gorm"
)

var trxService *TrxService

type TrxService struct {
	Common       *common.CommonService
	Loans        *loans.LoanService
	Trx          transaction.TrxRepository
	Installment  installment.InstallmentRepository
	IdentityRule identityrules.IdentityRuleRepository
	Pricing      pricing.PricingRepository
	db           *gorm.DB
	sync.Mutex
}

// TrxServiceConfiguration is an alias for a function that will take in a pointer to an TrxService and modify it
type TrxServiceConfiguration func(conf *TrxService) error

func newTrxService(cfgs ...TrxServiceConfiguration) *TrxService {
	// Create the NewTrxService
	conf := &TrxService{}
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

// withService applies a given repositories to the TrxService
func withService() TrxServiceConfiguration {
	// return a function that matches the TrxServiceConfiguration signature
	return func(s *TrxService) error {
		s.Common = common.Apply() // with common service
		if s.db == nil {
			s.db = s.Common.Db.GetDbConn()
		}

		s.Loans = loans.Apply(s.db)
		return nil
	}
}

// withRepository applies a given repositories to the TrxService
func withRepository() TrxServiceConfiguration {
	// return a function that matches the TrxServiceConfiguration signature
	return func(s *TrxService) error {
		s.Trx = transaction.NewTransaction(s.db)
		s.Installment = installment.NewInstallment().WithDbConn(s.db)
		s.IdentityRule = identityrules.NewIdentityRules()
		s.Pricing = pricing.NewPricing(s.db)
		return nil
	}
}

// Apply Transaction Service
func Apply() *TrxService {
	if trxService == nil {
		s := newTrxService(
			withService(),
			withRepository(),
		)
		trxService = s
	}
	return trxService
}

// WithDb
func (s *TrxService) WithDb(db *gorm.DB) *TrxService {
	s.db = db
	return s
}
