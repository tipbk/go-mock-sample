package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CallExternalFunction(t *testing.T) {
	// setup
	externalServiceMock := NewExternalServiceMock()
	externalServiceMock.On("GetHelloWorld", "Somchai").Return("Mai Chai Somchai")
	myService := NewMyService(externalServiceMock)
	// test
	actual := myService.CallExternalService("Somchai")
	expected := "Mai Chai Somchai"
	// assertion
	assert.Equal(t, expected, actual)
}
