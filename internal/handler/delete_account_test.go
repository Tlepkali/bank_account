package handler

import (
	"testing"

	"bank_account/internal/models"
	mockModel "bank_account/internal/models/mock"
	"bank_account/internal/repository"
	"bank_account/internal/service"
)

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
