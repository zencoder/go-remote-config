package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_SWF_CLIENT_REGION      AWSRegion = AWS_REGION_US_EAST_1
	VALID_SWF_CLIENT_ENDPOINT    string    = "http://localhost:8000/swf"
	VALID_SWF_CLIENT_DISABLE_SSL bool      = true
)

type SimpleWorkflowClientConfigSuite struct {
	suite.Suite
}

func TestSimpleWorkflowClientConfigSuite(t *testing.T) {
	suite.Run(t, new(SimpleWorkflowClientConfigSuite))
}

func (s *SimpleWorkflowClientConfigSuite) SetupSuite() {
}

func (s *SimpleWorkflowClientConfigSuite) SetupTest() {
}

func (s *SimpleWorkflowClientConfigSuite) TestValidate() {
	region := VALID_SWF_CLIENT_REGION
	endpoint := VALID_SWF_CLIENT_ENDPOINT

	c := &SimpleWorkflowClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SimpleWorkflowClientConfigSuite) TestValidateWithoutEndpoint() {
	region := VALID_SWF_CLIENT_REGION

	c := &SimpleWorkflowClientConfig{
		Region: &region,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SimpleWorkflowClientConfigSuite) TestValidateErrorEmptyEndpoint() {
	region := VALID_SWF_CLIENT_REGION
	endpoint := ""

	c := &SimpleWorkflowClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Endpoint, contains an empty string"), err)
}

func (s *SimpleWorkflowClientConfigSuite) TestValidateErrorRegion() {
	region := AWSRegion("")

	c := &SimpleWorkflowClientConfig{
		Region: &region,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: Region, failed to validate with error, Region cannot be empty"), err)
}

func (s *SimpleWorkflowClientConfigSuite) TestGetRegion() {
	region := VALID_SWF_CLIENT_REGION

	c := &SimpleWorkflowClientConfig{
		Region: &region,
	}

	assert.Equal(s.T(), VALID_SWF_CLIENT_REGION, c.GetRegion())
}

func (s *SimpleWorkflowClientConfigSuite) TestGetEndpoint() {
	region := VALID_SWF_CLIENT_REGION
	endpoint := VALID_SWF_CLIENT_ENDPOINT

	c := &SimpleWorkflowClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	assert.Equal(s.T(), VALID_SWF_CLIENT_ENDPOINT, c.GetEndpoint())
}

func (s *SimpleWorkflowClientConfigSuite) TestGetEndpointNotSet() {
	region := VALID_SWF_CLIENT_REGION

	c := &SimpleWorkflowClientConfig{
		Region:   &region,
		Endpoint: nil,
	}

	assert.Equal(s.T(), "", c.GetEndpoint())
}

func (s *SimpleWorkflowClientConfigSuite) TestGetDisableSSL() {
	region := VALID_SWF_CLIENT_REGION
	endpoint := VALID_SWF_CLIENT_ENDPOINT
	disableSSL := VALID_SWF_CLIENT_DISABLE_SSL

	c := &SimpleWorkflowClientConfig{
		Region:     &region,
		Endpoint:   &endpoint,
		DisableSSL: &disableSSL,
	}

	assert.Equal(s.T(), VALID_SWF_CLIENT_DISABLE_SSL, c.GetDisableSSL())
}

func (s *SimpleWorkflowClientConfigSuite) TestGetDisableSSLNotSet() {
	region := VALID_SWF_CLIENT_REGION
	endpoint := VALID_SWF_CLIENT_ENDPOINT

	c := &SimpleWorkflowClientConfig{
		Region:     &region,
		Endpoint:   &endpoint,
		DisableSSL: nil,
	}

	assert.Equal(s.T(), false, c.GetDisableSSL())
}
