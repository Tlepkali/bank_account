package handler

import (
	"testing"

	"bank_account/internal/models"
	"bank_account/internal/repository"
	"bank_account/internal/service"

	mockModel "bank_account/internal/models/mock"
)

func TestGetAccountByNumber(t *testing.T) {
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
		{"valid", `/account/1234567890`, 200, `{"owner":"test","account_number":"1234567890","balance":100,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`},
		{"invalid", `/account/132`, 404, `{"message":"invalid account number"}`},
		{"invalid1", `/account/12345678901`, 404, `{"message":"not found"}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr, err := testServer.getByNumber(t, tc.body)
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
