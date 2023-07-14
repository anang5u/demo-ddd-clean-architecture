package common

import "demo-ddd-clean-architecture/app/domain/identityrules"

var cmnService *CommonService

type CommonService struct {
	Config       ConfigRepository
	Log          LoggerRepository
	Db           DatabaseRepository
	DbMock       DatabaseMockRepository
	IdentityRule identityrules.IdentityRuleRepository
}

// CommonServiceConfiguration is an alias for a function that will take in a pointer to an CommonService and modify it
type CommonServiceConfiguration func(conf *CommonService) error

func newCommonService(cfgs ...CommonServiceConfiguration) *CommonService {
	// Create the NewCommonService
	conf := &CommonService{}
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

// withConfigRepository applies a given repositories to the CommonService
func withConfigRepository(filenames ...string) CommonServiceConfiguration {
	// return a function that matches the CommonServiceConfiguration signature
	return func(s *CommonService) error {
		s.Config = NewConfig(filenames...)
		return nil
	}
}

// withRepository applies a given repositories to the CommonService
func withRepository() CommonServiceConfiguration {
	// return a function that matches the CommonServiceConfiguration signature
	return func(s *CommonService) error {
		s.Log = NewLogger(s.Config)
		s.Db = NewDatabase(s.Config, s.Log)
		s.DbMock = NewDbMock()
		s.IdentityRule = identityrules.NewIdentityRules()
		return nil
	}
}

// Apply Common Service
func Apply(configurationFilenames ...string) *CommonService {
	if cmnService == nil {
		s := newCommonService(
			withConfigRepository(configurationFilenames...),
			withRepository(),
		)
		cmnService = s
	}
	return cmnService
}
