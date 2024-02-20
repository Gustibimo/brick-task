package usecases

import (
	"brick-task/internal/domain/payments/mocks"
	"brick-task/internal/domain/payments/models"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBankGateway is a mock implementation of the BankGateway interface
type MockBankGateway struct {
	mock.Mock
}

// BankAccountInquiry mocks the BankAccountInquiry method
func (m *MockBankGateway) BankAccountInquiry(accountNumber, bankCode string) (bool, error) {
	args := m.Called(accountNumber, bankCode)
	return args.Bool(0), args.Error(1)
}

func (m *MockBankGateway) DoPayment(request models.DoPaymentRequest) (bool, error) {
	args := m.Called(request)
	return args.Bool(0), args.Error(1)
}

func TestValidateBankAccount(t *testing.T) {
	// Mock data
	accountNumber := "1234567890"
	bankCode := "123"

	// Mocking the bank gateway
	mockGateway := new(MockBankGateway)
	mockRepo := new(mocks.Repository)
	mockGateway.On("BankAccountInquiry", accountNumber, bankCode).Return(true, nil)

	// Create payment use case with the mocked bank gateway
	useCase := NewPaymentUsecase(mockRepo, mockGateway)

	// Test case 1: Valid inquiry response
	valid, err := useCase.ValidateBankAccount(context.Background(), accountNumber, bankCode)
	assert.NoError(t, err)
	assert.True(t, valid)

	// Test case 2: Error from bank gateway
	expectedErr := errors.New("error from bank gateway")
	mockGateway.On("BankAccountInquiry", accountNumber, bankCode).Return(false, expectedErr)
	valid, err = useCase.ValidateBankAccount(context.Background(), accountNumber, bankCode)
	assert.Error(t, err)
	assert.False(t, valid)
	assert.Equal(t, expectedErr, err)
}
