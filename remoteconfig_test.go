package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_REMOTE_CONFIG_SQS_REGION          AWSRegion = AWS_REGION_US_EAST_1
	VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID  string    = "345833302425"
	VALID_REMOTE_CONFIG_SQS_QUEUE_NAME      string    = "testQueue"
	VALID_REMOTE_CONFIG_DYNAMODB_REGION     AWSRegion = AWS_REGION_US_EAST_1
	VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME string    = "testTable"
)

type RemoteConfigSuite struct {
	suite.Suite
}

func TestRemoteConfigSuite(t *testing.T) {
	suite.Run(t, new(RemoteConfigSuite))
}

type SampleConfig struct {
	SQS      *SQSConfig      `json:"sqs,omitempty"`
	DynamoDB *DynamoDBConfig `json:"dynamodb,omitempty"`
	Str      *string         `json:"str,omitempty"`
}

func (s *RemoteConfigSuite) SetupSuite() {
}

func (s *RemoteConfigSuite) SetupTest() {
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflection() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		region:       &sqsRegion,
		awsAccountID: &sqsAWSAccountID,
		queueName:    &sqsQueueName,
	}

	dynamodbRegion := VALID_REMOTE_CONFIG_DYNAMODB_REGION
	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodb := &DynamoDBConfig{
		region:    &dynamodbRegion,
		tableName: &dynamodbTableName,
	}

	str := "testString"

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: dynamodb,
		Str:      &str,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorSQSConfigNotSet() {
	c := &SampleConfig{
		SQS: nil,
	}
	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("SQS Config Not Set, SQS"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorSQSConfigValidate() {
	sqsRegion := AWSRegion("invalidregion")
	sqs := &SQSConfig{
		region: &sqsRegion,
	}
	c := &SampleConfig{
		SQS: sqs,
	}
	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Failed to Validate SQSConfig, SQS with error: Invalid SQS Region"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBConfigNotSet() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		region:       &sqsRegion,
		awsAccountID: &sqsAWSAccountID,
		queueName:    &sqsQueueName,
	}

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("DynamoDB Config Not Set, DynamoDB"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBConfigValidate() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		region:       &sqsRegion,
		awsAccountID: &sqsAWSAccountID,
		queueName:    &sqsQueueName,
	}

	dynamodbRegion := AWSRegion("invalidregion")
	dynamodb := &DynamoDBConfig{
		region: &dynamodbRegion,
	}

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: dynamodb,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Failed to Validate DynamoDBConfig, DynamoDB with error: Invalid DynamoDB Region"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorStrEmpty() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		region:       &sqsRegion,
		awsAccountID: &sqsAWSAccountID,
		queueName:    &sqsQueueName,
	}

	dynamodbRegion := VALID_REMOTE_CONFIG_DYNAMODB_REGION
	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodb := &DynamoDBConfig{
		region:    &dynamodbRegion,
		tableName: &dynamodbTableName,
	}

	str := ""

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: dynamodb,
		Str:      &str,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Empty String, Str"), err)
}
