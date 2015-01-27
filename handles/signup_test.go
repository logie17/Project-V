package handles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

const numAsserts int = 8

type SignupTestSuite struct {
	suite.Suite
	Asserts int
}

// entry point
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(SignupTestSuite))
}

// test the sting validator
func (suite *SignupTestSuite) TestStringValidator() {
	suite.Equal(true, stringValidator("cat", 3), "string should be valid")
	suite.Asserts++
	suite.Equal(false, stringValidator("cat", 5), "string should be invalid")
	suite.Asserts++
	suite.Equal(false, stringValidator("", 5), "string should be invalid")
	suite.Asserts++
	suite.Equal(true, stringValidator("apples", 2), "sting should be valid")
	suite.Asserts++
}

// test the email validator
func (suite *SignupTestSuite) TestEmailValidator() {
	suite.Equal(false, emailValidator("cat"), "cat is not an email addr")
	suite.Asserts++
	suite.Equal(true, emailValidator("cat@cats.com"), "valid email")
	suite.Asserts++
	suite.Equal(false, emailValidator(""), "empty string is not an email")
	suite.Asserts++
	suite.Equal(true, emailValidator("asd.dfs@sdf22.co.uk"), "weird email is still valid")
	suite.Asserts++
}

// test number of asserts
func (suite *SignupTestSuite) TearDownSuite() {
	suite.Equal(numAsserts, suite.Asserts, fmt.Sprintf("we should have %d asserts", numAsserts))
}
