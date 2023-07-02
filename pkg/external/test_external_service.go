package external

import "github.com/chalfel/gomock-test/pkg/ledger"

type TestExternalService struct {
}

func NewTestExternalService() ledger.TestService {
	return &TestExternalService{}
}

func (t *TestExternalService) Create(id string) error {
	return nil
}
