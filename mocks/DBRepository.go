// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "restAPI/pkg/storage/mysql/entity"

	mock "github.com/stretchr/testify/mock"
)

// DBRepository is an autogenerated mock type for the DBRepository type
type DBRepository struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: user
func (_m *DBRepository) AddUser(user entity.User) (entity.User, error) {
	ret := _m.Called(user)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseAccount provides a mock function with given fields: _a0
func (_m *DBRepository) CloseAccount(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAccount provides a mock function with given fields: _a0
func (_m *DBRepository) CreateAccount(_a0 entity.Account) (entity.Account, error) {
	ret := _m.Called(_a0)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(entity.Account) entity.Account); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Account) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccountType provides a mock function with given fields: accountType
func (_m *DBRepository) CreateAccountType(accountType entity.AccountType) error {
	ret := _m.Called(accountType)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.AccountType) error); ok {
		r0 = rf(accountType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: id
func (_m *DBRepository) DeleteUser(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DepositMoney provides a mock function with given fields: _a0, _a1
func (_m *DBRepository) DepositMoney(_a0 string, _a1 float32) (entity.Account, error) {
	ret := _m.Called(_a0, _a1)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(string, float32) entity.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, float32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAccountTypes provides a mock function with given fields:
func (_m *DBRepository) GetAllAccountTypes() ([]entity.AccountType, error) {
	ret := _m.Called()

	var r0 []entity.AccountType
	if rf, ok := ret.Get(0).(func() []entity.AccountType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.AccountType)
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

// GetAllAccounts provides a mock function with given fields:
func (_m *DBRepository) GetAllAccounts() []entity.Account {
	ret := _m.Called()

	var r0 []entity.Account
	if rf, ok := ret.Get(0).(func() []entity.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Account)
		}
	}

	return r0
}

// GetUser provides a mock function with given fields: id
func (_m *DBRepository) GetUser(id string) (entity.User, error) {
	ret := _m.Called(id)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields:
func (_m *DBRepository) GetUsers() ([]entity.User, error) {
	ret := _m.Called()

	var r0 []entity.User
	if rf, ok := ret.Get(0).(func() []entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
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

// UpdateUser provides a mock function with given fields: user
func (_m *DBRepository) UpdateUser(user entity.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithdrawMoney provides a mock function with given fields: _a0, _a1
func (_m *DBRepository) WithdrawMoney(_a0 string, _a1 float32) (entity.Account, error) {
	ret := _m.Called(_a0, _a1)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(string, float32) entity.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, float32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
