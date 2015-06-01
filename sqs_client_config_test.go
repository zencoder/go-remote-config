package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_SQS_CLIENT_REGION   AWSRegion = AWS_REGION_US_EAST_1
	VALID_SQS_CLIENT_ENDPOINT string    = "http://localhost/testsqs"
)

type SQSClientConfigSuite struct {
	suite.Suite
}

func TestSQSClientConfigSuite(t *testing.T) {
	suite.Run(t, new(SQSClientConfigSuite))
}

func (s *SQSClientConfigSuite) SetupSuite() {
}

func (s *SQSClientConfigSuite) SetupTest() {
}

func (s *SQSClientConfigSuite) TestValidate() {
	region := VALID_SQS_CLIENT_REGION

	c := &SQSClientConfig{
		Region: &region,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SQSClientConfigSuite) TestValidateWithEndpoint() {
	region := VALID_SQS_CLIENT_REGION
	endpoint := VALID_SQS_CLIENT_ENDPOINT

	c := &SQSClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SQSClientConfigSuite) TestValidateErrorRegion() {
	region := AWSRegion("invalidregion")

	c := &SQSClientConfig{
		Region: &region,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: Region, failed to validate with error, Region is invalid"), err)
}

func (s *SQSClientConfigSuite) TestValidateErrorEndpoint() {
	region := VALID_SQS_CLIENT_REGION
	endpoint := ""

	c := &SQSClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Endpoint, contains an empty string"), err)
}

func (s *SQSClientConfigSuite) TestGetRegion() {
	region := VALID_SQS_CLIENT_REGION

	c := &SQSClientConfig{
		Region: &region,
	}

	sRegion := c.GetRegion()
	assert.Equal(s.T(), VALID_SQS_CLIENT_REGION, sRegion)
}

func (s *SQSClientConfigSuite) TestGetEndpoint() {
	region := VALID_SQS_CLIENT_REGION
	endpoint := VALID_SQS_CLIENT_ENDPOINT

	c := &SQSClientConfig{
		Region:   &region,
		Endpoint: &endpoint,
	}

	sEndpoint := c.GetEndpoint()
	assert.Equal(s.T(), VALID_SQS_CLIENT_ENDPOINT, sEndpoint)
}
