package handler

import (
	"net/http/httptest"
	"testing"

	"bank_account/internal/models"
	"bank_account/internal/service"
	"bank_account/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAllAccountsMock(t *testing.T) {
	serviceMock := &mocks.AccountService{}

	service := &service.Service{
		AccountService: serviceMock,
	}

	handler := NewHandler(service)

	cases := []struct {
		name         string
		expectedCode int
		wantResponse string
		mockOn       bool
		mockService  interface{}
	}{
		{
			name:         "valid",
			expectedCode: 200,
			wantResponse: `"account_number": "1234567890",`,
			mockOn:       true,
			mockService: func() ([]*models.Account, error) {
				return []*models.Account{
					{
						Owner:         "test",
						AccountNumber: "1234567890",
						Balance:       100,
					},
				}, nil
			},
		},
		{
			name:         "invalid",
			expectedCode: 404,
			wantResponse: `"error": "account not found"`,
			mockOn:       true,
			mockService: func() ([]*models.Account, error) {
				return nil, models.ErrNotFound
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockOn {
				serviceMock.On("GetAllAccounts").
					Return(tc.mockService).Once()
			}

			res := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/accounts", nil)

			handler.GetAllAccounts(res, req)

			assert.Equal(t, tc.expectedCode, res.Code)
			assert.Contains(t, res.Body.String(), tc.wantResponse)
		})
	}
}
