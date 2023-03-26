package testify

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type Database interface {
	Get(key string) (string, error)
}

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func TestGetData(t *testing.T) {
	mockDB := new(MockDatabase)
	mockDB.On("Get", "foo").Return("bar", nil)

	data, _ := mockDB.Get("foo")

	if data != "bar" {
		t.Errorf("Unexpected data: %s", data)
	}

	mockDB.AssertExpectations(t)
}
