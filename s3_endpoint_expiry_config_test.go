package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_S3_ENDPOINT_EXPIRY_CONFIG_ENDPOINT string = "http://localhost:9500/s3"
	VALID_S3_ENDPOINT_EXPIRY_CONFIG_EXPIRY   uint   = 30
)

type S3EndpointExpiryConfigSuite struct {
	suite.Suite
}

func TestS3EndpointExpiryConfigSuite(t *testing.T) {
	suite.Run(t, new(S3EndpointExpiryConfigSuite))
}

func (s *S3EndpointExpiryConfigSuite) SetupSuite() {
}

func (s *S3EndpointExpiryConfigSuite) SetupTest() {
}

func (s *S3EndpointExpiryConfigSuite) TestValidate() {
	endpoint := VALID_S3_CONFIG_ENDPOINT
	expiry := VALID_S3_CONFIG_EXPIRY

	c := &S3EndpointExpiryConfig{
		Endpoint: &endpoint,
		Expiry:   &expiry,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *S3EndpointExpiryConfigSuite) TestValidateErrorEndpoint() {
	endpoint := ""

	c := &S3EndpointExpiryConfig{
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Endpoint, contains an empty string"), err)
}

func (s *S3EndpointExpiryConfigSuite) TestGetExpirySet() {
	expiry := VALID_S3_CONFIG_EXPIRY

	c := &S3EndpointExpiryConfig{
		Expiry: &expiry,
	}

	cExpiry := c.GetExpiry()
	assert.Equal(s.T(), VALID_S3_ENDPOINT_EXPIRY_CONFIG_EXPIRY, cExpiry)
}

func (s *S3EndpointExpiryConfigSuite) TestGetExpiryNotSet() {
	c := &S3EndpointExpiryConfig{
		Expiry: nil,
	}

	cExpiry := c.GetExpiry()
	assert.Equal(s.T(), S3_ENDPOINT_EXPIRY_CONFIG_DEFAULT_EXPIRY, cExpiry)
}

func (s *S3EndpointExpiryConfigSuite) TestGetEndpointSet() {
	endpoint := VALID_S3_CONFIG_ENDPOINT

	c := &S3EndpointExpiryConfig{
		Endpoint: &endpoint,
	}

	cEndpoint := c.GetEndpoint()
	assert.Equal(s.T(), VALID_S3_ENDPOINT_EXPIRY_CONFIG_ENDPOINT, cEndpoint)
}

func (s *S3EndpointExpiryConfigSuite) TestGetEndpointNotSet() {
	c := &S3EndpointExpiryConfig{
		Endpoint: nil,
	}

	cEndpoint := c.GetEndpoint()
	assert.Equal(s.T(), "", cEndpoint)
}
