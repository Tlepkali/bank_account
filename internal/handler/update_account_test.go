package handler

import (
	"testing"

	"bank_account/internal/models"
	mockModel "bank_account/internal/models/mock"
	"bank_account/internal/repository"
	"bank_account/internal/service"
)

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
