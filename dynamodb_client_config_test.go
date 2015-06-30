package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_DYNAMODB_CLIENT_REGION      AWSRegion = AWS_REGION_US_EAST_1
	VALID_DYNAMODB_CLIENT_ENDPOINT    string    = "http://localhost:8000/dynamodb"
	VALID_DYNAMODB_CLIENT_DISABLE_SSL bool      = true
)

type DynamoDBClientConfigSuite struct {
	suite.Suite
}

func TestDynamoDBClientConfigSuite(t *testing.T) {
	suite.Run(t, new(DynamoDBClientConfigSuite))
}

func (s *DynamoDBClientConfigSuite) SetupSuite() {
}

func (s *DynamoDBClientConfigSuite) SetupTest() {
}

func (s *DynamoDBClientConfigSuite) TestValidate() {
	region := VALID_DYNAMODB_CLIENT_REGION
	endpoint := VALID_DYNAMODB_CLIENT_ENDPOINT

	d := &DynamoDBClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(d)
	assert.Nil(s.T(), err)
}

func (s *DynamoDBClientConfigSuite) TestValidateWithoutEndpoint() {
	region := VALID_DYNAMODB_CLIENT_REGION

	d := &DynamoDBClientConfig{
		Region: &region,
	}

	err := validateConfigWithReflection(d)
	assert.Nil(s.T(), err)
}

func (s *DynamoDBClientConfigSuite) TestValidateErrorEmptyEndpoint() {
	region := VALID_DYNAMODB_CLIENT_REGION
	endpoint := ""

	d := &DynamoDBClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(d)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Endpoint, contains an empty string"), err)
}

func (s *DynamoDBClientConfigSuite) TestValidateErrorRegion() {
	region := AWSRegion("")

	d := &DynamoDBClientConfig{
		Region: &region,
	}

	err := validateConfigWithReflection(d)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: Region, failed to validate with error, Region cannot be empty"), err)
}

func (s *DynamoDBClientConfigSuite) TestGetRegion() {
	region := VALID_DYNAMODB_CLIENT_REGION

	d := &DynamoDBClientConfig{
		Region: &region,
	}

	assert.Equal(s.T(), VALID_DYNAMODB_CLIENT_REGION, d.GetRegion())
}

func (s *DynamoDBClientConfigSuite) TestGetEndpoint() {
	region := VALID_DYNAMODB_CLIENT_REGION
	endpoint := VALID_DYNAMODB_CLIENT_ENDPOINT

	d := &DynamoDBClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	assert.Equal(s.T(), VALID_DYNAMODB_CLIENT_ENDPOINT, d.GetEndpoint())
}

func (s *DynamoDBClientConfigSuite) TestGetEndpointNotSet() {
	region := VALID_DYNAMODB_CLIENT_REGION

	d := &DynamoDBClientConfig{
		Region:   &region,
		Endpoint: nil,
	}

	assert.Equal(s.T(), "", d.GetEndpoint())
}

func (s *DynamoDBClientConfigSuite) TestGetDisableSSL() {
	region := VALID_DYNAMODB_CLIENT_REGION
	endpoint := VALID_DYNAMODB_CLIENT_ENDPOINT
	disableSSL := VALID_DYNAMODB_CLIENT_DISABLE_SSL

	d := &DynamoDBClientConfig{
		Region:     &region,
		Endpoint:   &endpoint,
		DisableSSL: &disableSSL,
	}

	assert.Equal(s.T(), VALID_DYNAMODB_CLIENT_DISABLE_SSL, d.GetDisableSSL())
}

func (s *DynamoDBClientConfigSuite) TestGetDisableSSLNotSet() {
	region := VALID_DYNAMODB_CLIENT_REGION
	endpoint := VALID_DYNAMODB_CLIENT_ENDPOINT

	d := &DynamoDBClientConfig{
		Region:     &region,
		Endpoint:   &endpoint,
		DisableSSL: nil,
	}

	assert.Equal(s.T(), false, d.GetDisableSSL())
}
