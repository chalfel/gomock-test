package http_test

import (
	"errors"
	"testing"

	mock_ledger "github.com/chalfel/gomock-test/pkg/ledger/mock"
	"github.com/chalfel/gomock-test/pkg/presentation/http"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTst(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockTestService := mock_ledger.NewMockTestService(ctrl)
	mockTestService.EXPECT().Create(gomock.Any()).Return(nil)

	testHandler := http.NewTestHandler(mockTestService)
	to := testHandler.Create("test")

	assert.Equal(t, nil, to)
}

func TestErrorCreateTst(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockTestService := mock_ledger.NewMockTestService(ctrl)
	mockTestService.EXPECT().Create(gomock.Any()).Return(errors.New("error"))

	testHandler := http.NewTestHandler(mockTestService)
	to := testHandler.Create("test")

	assert.Equal(t, "error", to.Error())
}
