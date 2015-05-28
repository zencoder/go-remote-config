package remoteconfig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_SQS_REGION         AWSRegion = AWS_REGION_US_EAST_1
	VALID_SQS_AWS_ACCOUNT_ID string    = "345833302425"
	VALID_SQS_QUEUE_NAME     string    = "testQueue"
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
		region:       &region,
		awsAccountID: &awsAccountID,
		queueName:    &queueName,
	}

	err := c.Validate()
	assert.Nil(s.T(), err)
}

func (s *SQSConfigSuite) TestValidateErrorRegion() {
	region := AWSRegion("invalidregion")
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_NAME

	c := &SQSConfig{
		region:       &region,
		awsAccountID: &awsAccountID,
		queueName:    &queueName,
	}

	err := c.Validate()
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSQSConfigInvalidRegion, err)
}

func (s *SQSConfigSuite) TestValidateErrorAWSAccountID() {
	region := VALID_SQS_REGION
	awsAccountID := ""
	queueName := VALID_SQS_QUEUE_NAME

	c := &SQSConfig{
		region:       &region,
		awsAccountID: &awsAccountID,
		queueName:    &queueName,
	}

	err := c.Validate()
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSQSConfigInvalidAWSAccountID, err)
}

func (s *SQSConfigSuite) TestValidateErrorQueueName() {
	region := VALID_SQS_REGION
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := ""

	c := &SQSConfig{
		region:       &region,
		awsAccountID: &awsAccountID,
		queueName:    &queueName,
	}

	err := c.Validate()
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSQSConfigInvalidQueueName, err)
}

func (s *SQSConfigSuite) TestURL() {
	region := VALID_SQS_REGION
	awsAccountID := VALID_SQS_AWS_ACCOUNT_ID
	queueName := VALID_SQS_QUEUE_NAME

	c := &SQSConfig{
		region:       &region,
		awsAccountID: &awsAccountID,
		queueName:    &queueName,
	}

	url, err := c.URL()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), VALID_SQS_QUEUE_URL, url)

}
