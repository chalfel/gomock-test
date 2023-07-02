package http

import "github.com/chalfel/gomock-test/pkg/ledger"

type TestHandler struct {
	testService ledger.TestService
}

func NewTestHandler(testService ledger.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

func (t *TestHandler) Create(id string) error {
	return t.testService.Create(id)
}
