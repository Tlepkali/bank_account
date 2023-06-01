package mock

import (
	"bank_account/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepoMock is an autogenerated mock type for the AccountRepository type
type AccountRepoMock struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: account
func (_m *AccountRepoMock) CreateAccount(account *models.Account) (string, error) {
	ret := _m.Called(account)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.Account) string); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Account) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountByNumber provides a mock function with given fields: accountNumber
func (_m *AccountRepoMock) GetAccountByNumber(accountNumber string) (*models.Account, error) {
	ret := _m.Called(accountNumber)

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func(string) *models.Account); ok {
		r0 = rf(accountNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAccounts provides a mock function with given fields:
func (_m *AccountRepoMock) GetAllAccounts() ([]*models.Account, error) {
	ret := _m.Called()

	var r0 []*models.Account
	if rf, ok := ret.Get(0).(func() []*models.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccount provides a mock function with given fields: account
func (_m *AccountRepoMock) UpdateAccount(account *models.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAccount provides a mock function with given fields: accountNumber
func (_m *AccountRepoMock) DeleteAccount(accountNumber string) error {
	ret := _m.Called(accountNumber)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(accountNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}