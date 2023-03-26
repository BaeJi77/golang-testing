package testify

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ExampleTestSuite struct {
	suite.Suite
	value string
}

func (suite *ExampleTestSuite) SetupSuite() {
	fmt.Println("before test")
}

func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("before each test")
	suite.value = "hello"
}

func (suite *ExampleTestSuite) TearDownSuite() {
	fmt.Println("after test")
}

// before each test
func (suite *ExampleTestSuite) TearDownTest() {
	fmt.Println("after each test")
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	fmt.Println("test1")
	suite.Equal(suite.value, "hello")
}

func (suite *ExampleTestSuite) TestExample2() {
	fmt.Println("test2")
	suite.Equal(suite.value, "hello")
}

func (suite *ExampleTestSuite) TestExample3() {
	fmt.Println("test3")
	suite.Equal(suite.value, "hello")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
