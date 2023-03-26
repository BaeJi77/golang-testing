package testify

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type BComponent interface {
	Do() string
}

type RealComponent struct{}

func (r *RealComponent) Do() string {
	return "I'm real"
}

type MockBComponent struct {
	mock.Mock
}

func (md *MockBComponent) Do() string {
	args := md.Called()
	return args.String(0)
}

type AComponent struct {
	b BComponent
}

func (out *AComponent) CallBComponent() string {
	return out.b.Do()
}

func TestAComponent(t *testing.T) {
	// Create a mock dependency
	mockBComponent := new(MockBComponent)
	mockBComponent.On("Do").Return("mock result")

	// Inject the mock dependency into the object under test
	aComponent := AComponent{b: mockBComponent}

	// Test the object under test
	result := aComponent.CallBComponent()
	expectedResult := "mock result"
	if result != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, result)
	}

	// Verify that the mock dependency was called
	mockBComponent.AssertCalled(t, "Do")
}
