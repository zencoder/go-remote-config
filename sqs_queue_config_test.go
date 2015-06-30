package remoteconfig

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_SQS_QUEUE_REGION         AWSRegion = AWS_REGION_US_EAST_1
	VALID_SQS_QUEUE_AWS_ACCOUNT_ID string    = "345833302425"
	VALID_SQS_QUEUE_QUEUE_NAME     string    = "testQueue"
	VALID_SQS_QUEUE_NO_ENDPOINT    string    = ""
	VALID_SQS_QUEUE_ENDPOINT       string    = "http://localhost:9500"
)

var (
	VALID_SQS_QUEUE_URL          string = fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", VALID_SQS_QUEUE_REGION, VALID_SQS_QUEUE_AWS_ACCOUNT_ID, VALID_SQS_QUEUE_QUEUE_NAME)
	VALID_SQS_QUEUE_URL_ENDPOINT string = fmt.Sprintf("%s/%s/%s", VALID_SQS_QUEUE_ENDPOINT, VALID_SQS_QUEUE_AWS_ACCOUNT_ID, VALID_SQS_QUEUE_QUEUE_NAME)
)

type SQSQueueConfigSuite struct {
	suite.Suite
}

func TestSQSQueueConfigSuite(t *testing.T) {
	suite.Run(t, new(SQSQueueConfigSuite))
}

func (s *SQSQueueConfigSuite) SetupSuite() {
}

func (s *SQSQueueConfigSuite) SetupTest() {
}

func (s *SQSQueueConfigSuite) TestValidate() {
	region := VALID_SQS_QUEUE_REGION
	awsAccountID := VALID_SQS_QUEUE_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_QUEUE_NAME

	c := &SQSQueueConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SQSQueueConfigSuite) TestValidateErrorRegion() {
	region := AWSRegion("invalidregion")
	awsAccountID := VALID_SQS_QUEUE_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_QUEUE_NAME

	c := &SQSQueueConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: Region, failed to validate with error, Region is invalid"), err)
}

func (s *SQSQueueConfigSuite) TestValidateErrorAWSAccountID() {
	region := VALID_SQS_QUEUE_REGION
	awsAccountID := ""
	queueName := VALID_SQS_QUEUE_QUEUE_NAME

	c := &SQSQueueConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: AWSAccountID, contains an empty string"), err)
}

func (s *SQSQueueConfigSuite) TestValidateErrorQueueName() {
	region := VALID_SQS_QUEUE_REGION
	awsAccountID := VALID_SQS_QUEUE_AWS_ACCOUNT_ID
	queueName := ""

	c := &SQSQueueConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: QueueName, contains an empty string"), err)
}

func (s *SQSQueueConfigSuite) TestGetURLNoEndpoint() {
	region := VALID_SQS_QUEUE_REGION
	awsAccountID := VALID_SQS_QUEUE_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_QUEUE_NAME

	c := &SQSQueueConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	url := c.GetURL(VALID_SQS_QUEUE_NO_ENDPOINT)
	assert.Equal(s.T(), VALID_SQS_QUEUE_URL, url)
}

func (s *SQSQueueConfigSuite) TestGetURLWithEndpoint() {
	region := VALID_SQS_QUEUE_REGION
	awsAccountID := VALID_SQS_QUEUE_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_QUEUE_NAME

	c := &SQSQueueConfig{
		Region:       &region,
		AWSAccountID: &awsAccountID,
		QueueName:    &queueName,
	}

	url := c.GetURL(VALID_SQS_QUEUE_ENDPOINT)
	assert.Equal(s.T(), VALID_SQS_QUEUE_URL_ENDPOINT, url)
}
