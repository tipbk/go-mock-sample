package service

import "github.com/stretchr/testify/mock"

type externalServiceMock struct {
	mock.Mock
}

func NewExternalServiceMock() *externalServiceMock {
	return &externalServiceMock{}
}

func (m *externalServiceMock) GetHelloWorld(name string) string {
	args := m.Called(name)
	return args.String(0)
}
