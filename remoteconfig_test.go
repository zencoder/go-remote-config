package remoteconfig

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	SQSOptional      *SQSConfig      `json:"sqs_optional,omitempty" remoteconfig:"optional"`
	DynamoDBOptional *DynamoDBConfig `json:"dynamodb_optional,omitempty" remoteconfig:"optional"`
	StrOptional      *string         `json:"str_optional,omitempty" remoteconfig:"optional"`
	SQS              *SQSConfig      `json:"sqs,omitempty"`
	DynamoDB         *DynamoDBConfig `json:"dynamodb,omitempty"`
	Str              *string         `json:"str,omitempty"`
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
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}

	dynamodbRegion := VALID_REMOTE_CONFIG_DYNAMODB_REGION
	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodb := &DynamoDBConfig{
		Region:    &dynamodbRegion,
		TableName: &dynamodbTableName,
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

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionWithOptional() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}

	dynamodbRegion := VALID_REMOTE_CONFIG_DYNAMODB_REGION
	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodb := &DynamoDBConfig{
		Region:    &dynamodbRegion,
		TableName: &dynamodbTableName,
	}

	str := "testString"

	c := &SampleConfig{
		SQSOptional:      sqs,
		DynamoDBOptional: dynamodb,
		StrOptional:      &str,
		SQS:              sqs,
		DynamoDB:         dynamodb,
		Str:              &str,
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
	assert.Equal(s.T(), errors.New("Field: SQS, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorSQSConfigValidate() {
	sqsRegion := AWSRegion("invalidregion")
	sqs := &SQSConfig{
		Region: &sqsRegion,
	}
	c := &SampleConfig{
		SQS: sqs,
	}
	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("SQSConfig Field: SQS, Failed to validate with error: Invalid SQS Region"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBConfigNotSet() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: DynamoDB, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBConfigValidate() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}

	dynamodbRegion := AWSRegion("invalidregion")
	dynamodb := &DynamoDBConfig{
		Region: &dynamodbRegion,
	}

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: dynamodb,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("DynamoDBConfig Field: DynamoDB, Failed to validate with error: Invalid DynamoDB Region"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorStrNotSet() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}

	dynamodbRegion := VALID_REMOTE_CONFIG_DYNAMODB_REGION
	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodb := &DynamoDBConfig{
		Region:    &dynamodbRegion,
		TableName: &dynamodbTableName,
	}

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: dynamodb,
		Str:      nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: Str, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorStrEmpty() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqs := &SQSConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}

	dynamodbRegion := VALID_REMOTE_CONFIG_DYNAMODB_REGION
	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodb := &DynamoDBConfig{
		Region:    &dynamodbRegion,
		TableName: &dynamodbTableName,
	}

	str := ""

	c := &SampleConfig{
		SQS:      sqs,
		DynamoDB: dynamodb,
		Str:      &str,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Str, contains an empty string"), err)
}

func (s *RemoteConfigSuite) TestLoadConfigFromS3Error() {
	c := &SQSConfig{}
	err := LoadConfigFromS3("invalid", AWSRegion("invalid"), c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("S3 URL does not start with the s3:// scheme"), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidate() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "sqs" : {
        "region" : "us-east-1",
        "aws_account_id" : "345833302425",
        "queue_name" : "testQueue"
      },
      "dynamodb" : {
        "region" : "us-east-1",
        "table_name" : "testTable"
      },
      "str" : "testStr"
    }`)
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := downloadJSONValidate(ts.URL, c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidateErrorDownloadFailed() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "JSON Download Failed", http.StatusNotFound)
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := downloadJSONValidate(ts.URL, c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), fmt.Errorf("Download of JSON failed, URL = %s, Response Code = %d", ts.URL, http.StatusNotFound), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidateErrorValidation() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := downloadJSONValidate(ts.URL, c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: SQS, not set"), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidateErrorInvalidJSON() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is NOT JSON")
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := downloadJSONValidate(ts.URL, c)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Failed to decode JSON, with error, invalid character 'T' looking for beginning of value"), err)
}
