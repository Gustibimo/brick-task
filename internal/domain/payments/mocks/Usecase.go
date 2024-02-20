package mocks

import (
	"brick-task/internal/domain/payments/models"
	"context"
	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (m *Usecase) ValidateBankAccount(ctx context.Context, accountNumber string, bankCode string) (bool, error) {
	ret := m.Called(accountNumber, bankCode)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(accountNumber, bankCode)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountNumber, bankCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *Usecase) Disburse(ctx context.Context, request models.DoPaymentRequest) (*models.Payment, error) {
	ret := m.Called(ctx, request)

	var r0 *models.Payment
	if rf, ok := ret.Get(0).(func(context.Context, models.DoPaymentRequest) *models.Payment); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.DoPaymentRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *Usecase) UpdatePaymentStatus(ctx context.Context, request models.PaymentStatusRequest) (bool, error) {

	ret := m.Called(ctx, request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, models.PaymentStatusRequest) bool); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.PaymentStatusRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
