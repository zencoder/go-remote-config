package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_S3_ENDPOINT_EXPIRY_CONFIG_ENDPOINT string = "http://localhost:9500/s3"
	//VALID_S3_CONFIG_BASE_BUCKET string    = "base-bucket"
	//VALID_S3_CONFIG_REGION      AWSRegion = AWS_REGION_US_WEST_2
	//VALID_S3_CONFIG_FILE_EXT    string    = "json"
	VALID_S3_ENDPOINT_EXPIRY_CONFIG_EXPIRY uint = 30
	//VALID_S3_CONFIG_TEST_PATH   string    = "test/path"
)

/*
var (
	VALID_S3_CONFIG_FULL_BUCKET_NAME string = fmt.Sprintf("%s-%s", VALID_S3_CONFIG_BASE_BUCKET, VALID_S3_CONFIG_REGION)
	VALID_S3_CONFIG_FULL_PATH        string = fmt.Sprintf("%s.%s", VALID_S3_CONFIG_TEST_PATH, VALID_S3_CONFIG_FILE_EXT)
	VALID_S3_CONFIG_S3_SCHEME_URL    string = fmt.Sprintf("s3://%s-%s/%s.%s", VALID_S3_CONFIG_BASE_BUCKET, VALID_S3_CONFIG_REGION, VALID_S3_CONFIG_TEST_PATH, VALID_S3_CONFIG_FILE_EXT)
)*/

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
