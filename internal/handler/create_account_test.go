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

func TestCreateAccountMock(t *testing.T) {
	serviceMock := &mocks.AccountService{}

	service := &service.Service{
		AccountService: serviceMock,
	}

	handler := NewHandler(service)

	cases := []struct {
		name         string
		body         string
		expectedCode int
		wantResponse string
		mockOn       bool
		mockService  interface{}
	}{
		{
			name:         "valid",
			body:         `{"owner": "Tlep", "balance": 2323}`,
			expectedCode: http.StatusCreated,
			wantResponse: `"message": "Account created successfully, `,
			mockOn:       true,
			mockService:  func(account *models.CreateAccountDTO) (string, error) { return "123", nil },
		},
		{
			name:         "invalid",
			body:         `{"owner": "Tlep", "balance": -2323}`,
			expectedCode: http.StatusUnprocessableEntity,
			wantResponse: `"balance": "must be greater than zero"`,
			mockOn:       false,
		},
		{
			name:         "invalid1",
			body:         `{"owner": "", "balance": 2323}`,
			expectedCode: http.StatusUnprocessableEntity,
			wantResponse: `"owner": "must be provided"`,
			mockOn:       false,
		},
		{
			name:         "invalid2",
			body:         `{"owner": "Tlep", "balance": 0}`,
			expectedCode: http.StatusConflict,
			wantResponse: `"error": "account already exists"`,
			mockOn:       true,
			mockService:  func(account *models.CreateAccountDTO) (string, error) { return "", models.ErrDuplicateAccount },
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request, _ := http.NewRequest("POST", "/account", strings.NewReader(tc.body))
			request.Header.Set("Content-Type", "application/json")

			responseRecorder := httptest.NewRecorder()

			if tc.mockOn {
				serviceMock.On("CreateAccount", mock.AnythingOfTypeArgument("*models.CreateAccountDTO")).
					Return(tc.mockService).Once()
			}

			handler.CreateAccount(responseRecorder, request)

			assert.Equal(t, tc.expectedCode, responseRecorder.Code)
			assert.Contains(t, responseRecorder.Body.String(), tc.wantResponse)

			if tc.mockOn {
				serviceMock.AssertCalled(t, "CreateAccount", mock.AnythingOfTypeArgument("*models.CreateAccountDTO"))
			}
		})
	}
}

// Table test without mocks
func TestTableCreateAccount(t *testing.T) {
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
		{"valid", `{"owner": "Tlep", "balance": 2323}`, 201, `{"account_number":"1234567890"}"}`},
		{"invalid", `{"owner": "", "balance": 100}`, 422, `{"error":"owner is empty"}`},
		{"invalid1", `{"owner": "dfds", "balance": -123}`, 422, `{"error":"balance is less than zero"}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr, err := testServer.post(t, "/account", []byte(tc.body))
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
