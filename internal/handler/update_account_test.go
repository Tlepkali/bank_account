package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bank_account/internal/models"
	mockModel "bank_account/internal/models/mock"
	"bank_account/internal/repository"
	"bank_account/internal/service"
	"bank_account/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateAccountMock(t *testing.T) {
	serviceMock := &mocks.AccountService{}

	service := &service.Service{
		AccountService: serviceMock,
	}

	handler := NewHandler(service)

	cases := []struct {
		name           string
		body           string
		accountNumber  string
		expectedCode   int
		wantResponse   string
		mockOn         bool
		getAccountMock interface{}
		mockService    interface{}
	}{
		{
			name:          "valid",
			body:          `{"owner": "Tlep", "balance": 2323}`,
			accountNumber: "1234567890",
			expectedCode:  200,
			wantResponse:  `"account_number": "1234567890"`,
			mockOn:        true,
			getAccountMock: func(accountNumber string) (*models.Account, error) {
				return &models.Account{
					Owner:         "test",
					AccountNumber: "1234567890",
					Balance:       100,
				}, nil
			},
			mockService: func(account *models.Account) error {
				return nil
			},
		},
		{
			name:          "invalid",
			body:          `{"owner": "", "balance": 100}`,
			accountNumber: "1234567890",
			expectedCode:  422,
			wantResponse:  `"owner": "must be provided"`,
			mockOn:        false,
		},
		{
			name:          "invalid1",
			body:          `{"owner": "dfds", "balance": 123}`,
			accountNumber: "34234324",
			expectedCode:  404,
			wantResponse:  `"error": "account not found"`,
			mockOn:        false,
			getAccountMock: func(accountNumber string) (*models.Account, error) {
				return nil, models.ErrNotFound
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPut, "/account/"+tc.accountNumber, strings.NewReader(tc.body))
			request.Header.Set("Content-Type", "application/json")

			response := httptest.NewRecorder()

			if tc.getAccountMock != nil {
				serviceMock.On("GetAccountByNumber", mock.AnythingOfType("string")).
					Return(tc.getAccountMock).Once()
			}

			if tc.mockOn {
				serviceMock.On("UpdateAccount", mock.AnythingOfType("*models.Account")).
					Return(tc.mockService).Once()
			}

			handler.UpdateAccount(response, request)

			assert.Equal(t, tc.expectedCode, response.Code)
			assert.Contains(t, response.Body.String(), tc.wantResponse)

			serviceMock.AssertExpectations(t)

			if tc.mockOn {
				serviceMock.AssertCalled(t, "UpdateAccount", mock.AnythingOfType("*models.Account"))
			}
		})
	}
}

func TestUpdateAccount(t *testing.T) {
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
		urlQuery     string
		expectedCode int
		wantResponse string
	}{
		{"valid", `{"owner": "Tlep", "balance": 2323}`, "/account/1234567890", 200, `{"account_number":"1234567890"}"}`},
		{"invalid", `{"owner": "", "balance": 100}`, "/account/1234567890", 422, `{"error":"owner is empty"}`},
		{"invalid1", `{"owner": "dfds", "balance": 123}`, "/account/34234324", 404, `{"error":"not found"}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr, err := testServer.put(t, tc.urlQuery, []byte(tc.body))
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
