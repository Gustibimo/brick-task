package mocks

import (
	"brick-task/internal/domain/payments/models"
	"context"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (m *Repository) SavePayment(ctx context.Context, payment models.Payment) (*models.Payment, error) {
	ret := m.Called(ctx, payment)

	var r0 *models.Payment
	if rf, ok := ret.Get(0).(func(context.Context, models.Payment) *models.Payment); ok {
		r0 = rf(ctx, payment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Payment) error); ok {
		r1 = rf(ctx, payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *Repository) UpdatePaymentStatus(ctx context.Context, status string, paymentId string) (bool, error) {
	ret := m.Called(ctx, status, paymentId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, status, paymentId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, status, paymentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
