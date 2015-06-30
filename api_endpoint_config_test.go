package remoteconfig

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_API_ENDPOINT_BASE_PATH string = "http://basepath"
	VALID_API_ENDPOINT_SUB_PATH  string = "sub/path"
	VALID_API_ENDPOINT_ID        string = "testid"
)

var (
	VALID_API_ENDPOINT_FULL_PATH         string = fmt.Sprintf("%s/%s", VALID_API_ENDPOINT_BASE_PATH, VALID_API_ENDPOINT_SUB_PATH)
	VALID_API_ENDPOINT_FULL_PATH_WITH_ID string = fmt.Sprintf("%s/%s/%s", VALID_API_ENDPOINT_BASE_PATH, VALID_API_ENDPOINT_SUB_PATH, VALID_API_ENDPOINT_ID)
)

type APIEndpointConfigSuite struct {
	suite.Suite
}

func TestAPIEndpointConfigSuite(t *testing.T) {
	suite.Run(t, new(APIEndpointConfigSuite))
}

func (s *APIEndpointConfigSuite) SetupSuite() {
}

func (s *APIEndpointConfigSuite) SetupTest() {
}

func (s *APIEndpointConfigSuite) TestValidate() {
	basePath := VALID_API_ENDPOINT_BASE_PATH
	subPath := VALID_API_ENDPOINT_SUB_PATH

	a := &APIEndpointConfig{
		BasePath: &basePath,
		SubPath:  &subPath,
	}

	err := validateConfigWithReflection(a)
	assert.Nil(s.T(), err)
}

func (s *APIEndpointConfigSuite) TestValidateErrorBasePathNotSet() {
	a := &APIEndpointConfig{
		BasePath: nil,
	}

	err := validateConfigWithReflection(a)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: BasePath, not set"), err)
}

func (s *APIEndpointConfigSuite) TestValidateErrorSubPathNotSet() {
	basePath := VALID_API_ENDPOINT_BASE_PATH

	a := &APIEndpointConfig{
		BasePath: &basePath,
		SubPath:  nil,
	}

	err := validateConfigWithReflection(a)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: SubPath, not set"), err)
}

func (s *APIEndpointConfigSuite) TestGetFullPath() {
	basePath := VALID_API_ENDPOINT_BASE_PATH
	subPath := VALID_API_ENDPOINT_SUB_PATH

	a := &APIEndpointConfig{
		BasePath: &basePath,
		SubPath:  &subPath,
	}

	fullPath := a.GetFullPath()
	assert.Equal(s.T(), VALID_API_ENDPOINT_FULL_PATH, fullPath)
}

func (s *APIEndpointConfigSuite) TestGetFullPathWithID() {
	basePath := VALID_API_ENDPOINT_BASE_PATH
	subPath := VALID_API_ENDPOINT_SUB_PATH

	a := &APIEndpointConfig{
		BasePath: &basePath,
		SubPath:  &subPath,
	}

	fullPathWithID := a.GetFullPathWithID(VALID_API_ENDPOINT_ID)
	assert.Equal(s.T(), VALID_API_ENDPOINT_FULL_PATH_WITH_ID, fullPathWithID)
}
