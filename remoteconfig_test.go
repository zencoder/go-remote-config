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
	VALID_REMOTE_CONFIG_SQS_REGION             AWSRegion = AWS_REGION_US_EAST_1
	VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID     string    = "345833302425"
	VALID_REMOTE_CONFIG_SQS_QUEUE_NAME         string    = "testQueue"
	VALID_REMOTE_CONFIG_DYNAMODB_CLIENT_REGION AWSRegion = AWS_REGION_US_EAST_1
	VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME    string    = "testTable"
	VALID_REMOTE_CONFIG_NO_ENDPOINT            string    = ""
)

type RemoteConfigSuite struct {
	suite.Suite
}

func TestRemoteConfigSuite(t *testing.T) {
	suite.Run(t, new(RemoteConfigSuite))
}

type SampleConfig struct {
	SQSQueueOptional       *SQSQueueConfig       `json:"sqs_queue_optional,omitempty" remoteconfig:"optional"`
	SQSClientOptional      *SQSClientConfig      `json:"sqs_client_optional,omitempty" remoteconfig:"optional"`
	DynamoDBTableOptional  *DynamoDBTableConfig  `json:"dynamodb_table_optional,omitempty" remoteconfig:"optional"`
	DynamoDBClientOptional *DynamoDBClientConfig `json:"dynamodb_client_optional,omitempty" remoteconfig:"optional"`
	StrOptional            *string               `json:"str_optional,omitempty" remoteconfig:"optional"`
	SQSQueue               *SQSQueueConfig       `json:"sqs_queue,omitempty"`
	SQSClient              *SQSClientConfig      `json:"sqs_client,omitempty"`
	DynamoDBTable          *DynamoDBTableConfig  `json:"dynamodb_table,omitempty"`
	DynamoDBClient         *DynamoDBClientConfig `json:"dynamodb_client,omitempty"`
	Str                    *string               `json:"str,omitempty"`
}

func (s *RemoteConfigSuite) SetupSuite() {
}

func (s *RemoteConfigSuite) SetupTest() {
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflection() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodbTable := &DynamoDBTableConfig{
		TableName: &dynamodbTableName,
	}

	dynamodbClientRegion := VALID_REMOTE_CONFIG_DYNAMODB_CLIENT_REGION
	dynamodbClient := &DynamoDBClientConfig{
		Region: &dynamodbClientRegion,
	}

	str := "testString"

	c := &SampleConfig{
		SQSQueue:       sqsQueue,
		SQSClient:      sqsClient,
		DynamoDBTable:  dynamodbTable,
		DynamoDBClient: dynamodbClient,
		Str:            &str,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionWithOptional() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodbTable := &DynamoDBTableConfig{
		TableName: &dynamodbTableName,
	}

	dynamodbClientRegion := VALID_REMOTE_CONFIG_DYNAMODB_CLIENT_REGION
	dynamodbClient := &DynamoDBClientConfig{
		Region: &dynamodbClientRegion,
	}

	str := "testString"

	c := &SampleConfig{
		SQSQueueOptional:       sqsQueue,
		SQSClientOptional:      sqsClient,
		DynamoDBTableOptional:  dynamodbTable,
		DynamoDBClientOptional: dynamodbClient,
		StrOptional:            &str,
		SQSQueue:               sqsQueue,
		SQSClient:              sqsClient,
		DynamoDBTable:          dynamodbTable,
		DynamoDBClient:         dynamodbClient,
		Str:                    &str,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorSQSQueueConfigNotSet() {
	c := &SampleConfig{
		SQSQueue: nil,
	}
	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: SQSQueue, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorSQSClientConfigNotSet() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	c := &SampleConfig{
		SQSQueue:  sqsQueue,
		SQSClient: nil,
	}
	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: SQSClient, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorSQSQueueConfigValidate() {
	sqsRegion := AWSRegion("invalidregion")
	sqsQueue := &SQSQueueConfig{
		Region: &sqsRegion,
	}
	c := &SampleConfig{
		SQSQueue: sqsQueue,
	}
	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Sub Field of SQSQueue, failed to validate with error, Validater Field: Region, failed to validate with error, Region is invalid"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBTableConfigNotSet() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	c := &SampleConfig{
		SQSQueue:      sqsQueue,
		SQSClient:     sqsClient,
		DynamoDBTable: nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: DynamoDBTable, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBClientConfigNotSet() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodbTable := &DynamoDBTableConfig{
		TableName: &dynamodbTableName,
	}

	c := &SampleConfig{
		SQSQueue:       sqsQueue,
		SQSClient:      sqsClient,
		DynamoDBTable:  dynamodbTable,
		DynamoDBClient: nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: DynamoDBClient, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBClientConfigValidate() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodbTable := &DynamoDBTableConfig{
		TableName: &dynamodbTableName,
	}

	dynamodbRegion := AWSRegion("invalidregion")
	dynamodbClient := &DynamoDBClientConfig{
		Region: &dynamodbRegion,
	}

	c := &SampleConfig{
		SQSQueue:       sqsQueue,
		SQSClient:      sqsClient,
		DynamoDBTable:  dynamodbTable,
		DynamoDBClient: dynamodbClient,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Sub Field of DynamoDBClient, failed to validate with error, Validater Field: Region, failed to validate with error, Region is invalid"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorDynamoDBTableConfigValidate() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	dynamodbTableName := ""
	dynamodbTable := &DynamoDBTableConfig{
		TableName: &dynamodbTableName,
	}

	c := &SampleConfig{
		SQSQueue:      sqsQueue,
		SQSClient:     sqsClient,
		DynamoDBTable: dynamodbTable,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Sub Field of DynamoDBTable, failed to validate with error, String Field: TableName, contains an empty string"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorStrNotSet() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodbTable := &DynamoDBTableConfig{
		TableName: &dynamodbTableName,
	}

	dynamodbClientRegion := VALID_REMOTE_CONFIG_DYNAMODB_CLIENT_REGION
	dynamodbClient := &DynamoDBClientConfig{
		Region: &dynamodbClientRegion,
	}

	c := &SampleConfig{
		SQSQueue:       sqsQueue,
		SQSClient:      sqsClient,
		DynamoDBTable:  dynamodbTable,
		DynamoDBClient: dynamodbClient,
		Str:            nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: Str, not set"), err)
}

func (s *RemoteConfigSuite) TestvalidateConfigWithReflectionErrorStrEmpty() {
	sqsRegion := VALID_REMOTE_CONFIG_SQS_REGION
	sqsAWSAccountID := VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID
	sqsQueueName := VALID_REMOTE_CONFIG_SQS_QUEUE_NAME
	sqsQueue := &SQSQueueConfig{
		Region:       &sqsRegion,
		AWSAccountID: &sqsAWSAccountID,
		QueueName:    &sqsQueueName,
	}
	sqsClient := &SQSClientConfig{
		Region: &sqsRegion,
	}

	dynamodbTableName := VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME
	dynamodbTable := &DynamoDBTableConfig{
		TableName: &dynamodbTableName,
	}

	dynamodbClientRegion := VALID_REMOTE_CONFIG_DYNAMODB_CLIENT_REGION
	dynamodbClient := &DynamoDBClientConfig{
		Region: &dynamodbClientRegion,
	}

	str := ""

	c := &SampleConfig{
		SQSQueue:       sqsQueue,
		SQSClient:      sqsClient,
		DynamoDBTable:  dynamodbTable,
		DynamoDBClient: dynamodbClient,
		Str:            &str,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Str, contains an empty string"), err)
}

func (s *RemoteConfigSuite) TestLoadConfigFromS3Error() {
	c := &SQSQueueConfig{}
	err := LoadConfigFromS3("invalid", AWSRegion("invalid"), VALID_REMOTE_CONFIG_NO_ENDPOINT, c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("S3 URL does not start with the s3:// scheme"), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidate() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "sqs_client" : {
        "region" : "us-east-1",
				"endpoint" : "http://localhost:3000/sqs"
			},
			"sqs_queue" : {
				"region" : "us-east-1",
        "aws_account_id" : "345833302425",
        "queue_name" : "testQueue"
      },
      "dynamodb_client" : {
        "region" : "us-east-1",
				"endpoint" : "http://localhost:8000/dynamodb"
			},
			"dynamodb_table" : {
        "table_name" : "testTable"
      },
      "str" : "testStr"
    }`)
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := DownloadJSONValidate(ts.URL, c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidateErrorDownloadFailed() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "JSON Download Failed", http.StatusNotFound)
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := DownloadJSONValidate(ts.URL, c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), fmt.Errorf("Download of JSON failed, URL = %s, Response Code = %d", ts.URL, http.StatusNotFound), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidateErrorValidation() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := DownloadJSONValidate(ts.URL, c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: SQSQueue, not set"), err)
}

func (s *RemoteConfigSuite) TestdownloadJSONValidateErrorInvalidJSON() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is NOT JSON")
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := DownloadJSONValidate(ts.URL, c)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Failed to decode JSON, with error, invalid character 'T' looking for beginning of value"), err)
}
