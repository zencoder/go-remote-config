package remoteconfig

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_SQS_REGION         AWSRegion = AWS_REGION_US_EAST_1
	VALID_SQS_AWS_ACCOUNT_ID string    = "345833302425"
	VALID_SQS_QUEUE_NAME     string    = "testQueue"
	VALID_SQS_ENDPOINT       string    = "http://localhost/testsqs"
)

var (
	VALID_SQS_QUEUE_URL string = fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", VALID_SQS_REGION, VALID_SQS_AWS_ACCOUNT_ID, VALID_SQS_QUEUE_NAME)
)

type SQSConfigSuite struct {
	suite.Suite
}

func TestSQSConfigSuite(t *testing.T) {
	suite.Run(t, new(SQSConfigSuite))
}

func (s *SQSConfigSuite) SetupSuite() {
}

func (s *SQSConfigSuite) SetupTest() {
}

func (s *SQSConfigSuite) TestValidate() {
	region := VALID_SQS_REGION
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_NAME

	c := &SQSConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SQSConfigSuite) TestValidateWithEndpoint() {
	region := VALID_SQS_REGION
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_NAME
	endpoint := VALID_SQS_ENDPOINT

	c := &SQSConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
		Endpoint:     &endpoint,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SQSConfigSuite) TestValidateErrorRegion() {
	region := AWSRegion("invalidregion")
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_NAME

	c := &SQSConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: Region, failed to validate with error, Region is invalid"), err)
}

func (s *SQSConfigSuite) TestValidateErrorAWSAccountID() {
	region := VALID_SQS_REGION
	awsAccountID := ""
	queueName := VALID_SQS_QUEUE_NAME

	c := &SQSConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: AWSAccountID, contains an empty string"), err)
}

func (s *SQSConfigSuite) TestValidateErrorQueueName() {
	region := VALID_SQS_REGION
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := ""

	c := &SQSConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: QueueName, contains an empty string"), err)
}

func (s *SQSConfigSuite) TestGetURL() {
	region := VALID_SQS_REGION
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_NAME

	c := &SQSConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	url := c.GetURL()
	assert.Equal(s.T(), VALID_SQS_QUEUE_URL, url)
}

func (s *SQSConfigSuite) TestGetURLEndpoint() {
	region := VALID_SQS_REGION
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_NAME
	endpoint := VALID_SQS_ENDPOINT

	c := &SQSConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
		Endpoint:     &endpoint,
	}

	url := c.GetURL()
	assert.Equal(s.T(), VALID_SQS_ENDPOINT, url)
}
