package handler

import (
	"testing"

	"bank_account/internal/models"
	mockModel "bank_account/internal/models/mock"
	"bank_account/internal/repository"
	"bank_account/internal/service"
)

func TestCreateAccount(t *testing.T) {
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
