package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGateway struct {
	mock.Mock
}

func (m *MockGateway) Charge() bool {
	args := m.Mock.Called()
	return args.Bool(0)
}

func TestChargeCustomer(t *testing.T) {
	m := &MockGateway{}
	m.On("Charge").Return(true)

	r := ChargeCustomer(m)
	m.Mock.AssertExpectations(t)

	assert.True(t, r, true)
}
