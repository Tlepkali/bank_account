package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bank_account/internal/models"
	mockModel "bank_account/internal/models/mock"
	"bank_account/internal/repository"
	"bank_account/internal/service"
	"bank_account/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteAccountMock(t *testing.T) {
	serviceMock := &mocks.AccountService{}

	service := &service.Service{
		AccountService: serviceMock,
	}

	handler := NewHandler(service)

	cases := []struct {
		name          string
		accountNumber string
		expectedCode  int
		wantResponse  string
		mockOn        bool
		mockService   interface{}
	}{
		{
			name:          "valid",
			accountNumber: "1234567890",
			expectedCode:  200,
			wantResponse:  `"message": "Account deleted successfully"`,
			mockOn:        true,
			mockService: func(accountNumber string) error {
				return nil
			},
		},
		{
			name:          "invalid",
			accountNumber: "12345678901",
			expectedCode:  404,
			wantResponse:  `"error": "account not found"`,
			mockOn:        true,
			mockService: func(accountNumber string) error {
				return models.ErrNotFound
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodDelete, "/account/"+tc.accountNumber, nil)
			response := httptest.NewRecorder()

			if tc.mockOn {
				serviceMock.On("DeleteAccount", mock.AnythingOfType("string")).
					Return(tc.mockService).Once()
			}

			handler.DeleteAccount(response, request)

			assert.Equal(t, tc.expectedCode, response.Code)
			assert.Contains(t, response.Body.String(), tc.wantResponse)

			serviceMock.AssertExpectations(t)
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	var accModel models.AccountRepository = &mockModel.AccountModel{}

	repo := &repository.Repository{
		AccountRepo: accModel,
	}

	service := service.NewService(repo)

	handler := NewHandler(service)

	testServer := newTestServer(handler.InitRoutes())

	testCases := []struct {
		name         string
		body         string
		expectedCode int
		wantResponse string
	}{
		{"valid", `/account/1234567890`, 200, `{"message": "account succesfully deleted"}}`},
		{"invalid", `/account/132`, 404, `{"message":"invalid account number"}`},
		{"invalid1", `/account/dfsdf`, 404, `{"message":"not found"}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr, err := testServer.delete(t, tc.body)
			if err != nil {
				t.Fatal(err)
			}

			err = testServer.checkResponse(rr, tc.expectedCode, tc.wantResponse)

			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
