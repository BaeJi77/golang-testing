package testify

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func errorFunc() error {
	return fmt.Errorf("hello world")
}

func noErrorFunc() error {
	return nil
}

func TestAssert(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, 1, "1 should be equal to 1")
	a.NotEqual(1, 2, "1 should not be equal to 2")
	a.Nil(nil, "nil should be nil")
	a.NotNil("not nil", "not nil should not be nil")
	a.True(1 == 1, "1 should be equal to 1")
	a.False(1 == 2, "1 should not be equal to 2")

	err := errorFunc()
	a.Error(err, "someFuncThatReturnsAnError should return an error")

	err = noErrorFunc()
	a.NoError(err, "someFuncThatDoesNotReturnAnError should not return an error")
}
