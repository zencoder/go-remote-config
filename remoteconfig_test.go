package remoteconfig

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_REMOTE_CONFIG_SQS_REGION              AWSRegion       = AWS_REGION_US_EAST_1
	VALID_REMOTE_CONFIG_SQS_AWS_ACCOUNT_ID      string          = "345833302425"
	VALID_REMOTE_CONFIG_SQS_QUEUE_NAME          string          = "testQueue"
	VALID_REMOTE_CONFIG_DYNAMODB_CLIENT_REGION  AWSRegion       = AWS_REGION_US_EAST_1
	VALID_REMOTE_CONFIG_DYNAMODB_TABLE_NAME     string          = "testTable"
	VALID_REMOTE_CONFIG_STORAGE_CONFIG_PROVIDER StorageProvider = STORAGE_PROVIDER_AWS
	VALID_REMOTE_CONFIG_STORAGE_CONFIG_LOCATION StorageLocation = (StorageLocation)(AWS_REGION_US_WEST_2)
)

type RemoteConfigSuite struct {
	suite.Suite
}

func TestRemoteConfigSuite(t *testing.T) {
	suite.Run(t, new(RemoteConfigSuite))
}

type EmbeddedConfig struct {
	EmbeddedStr *string `json:"embedded_string,omitempty"`
	EmbeddedInt *int64  `json:"embedded_int,omitempty"`
}

type SampleConfig struct {
	EmbeddedConfig
	SQSQueueOptional           *SQSQueueConfig           `json:"sqs_queue_optional,omitempty" remoteconfig:"optional"`
	SQSClientOptional          *SQSClientConfig          `json:"sqs_client_optional,omitempty" remoteconfig:"optional"`
	DynamoDBTableOptional      *DynamoDBTableConfig      `json:"dynamodb_table_optional,omitempty" remoteconfig:"optional"`
	DynamoDBClientOptional     *DynamoDBClientConfig     `json:"dynamodb_client_optional,omitempty" remoteconfig:"optional"`
	StrOptional                *string                   `json:"str_optional,omitempty" remoteconfig:"optional"`
	StorageConfigOptional      *StorageConfig            `json:"storage_config_optional,omitempty" remoteconfig:"optional"`
	StorageConfigSliceOptional []*StorageConfig          `json:"storage_config_slice_optional,omitempty" remoteconfig:"optional"`
	SQSQueue                   *SQSQueueConfig           `json:"sqs_queue,omitempty"`
	SQSClient                  *SQSClientConfig          `json:"sqs_client,omitempty"`
	DynamoDBTable              *DynamoDBTableConfig      `json:"dynamodb_table,omitempty"`
	DynamoDBClient             *DynamoDBClientConfig     `json:"dynamodb_client,omitempty"`
	Str                        *string                   `json:"str,omitempty"`
	StorageConfig              *StorageConfig            `json:"storage_config,omitempty"`
	StorageConfigSlice         []*StorageConfig          `json:"storage_config_slice,omitempty"`
	StorageConfigMap           map[string]*StorageConfig `json:"storage_config_map,omitempty"`
	StrSlice                   []*string                 `json:"str_slice,omitempty" remoteconfig:"optional"`
	MapStrStr                  map[string]*string        `json:"map_str_str,omitempty"`
}

var validConfigJSON = `
	{
		"embedded_string": "abc",
		"embedded_int": 123,
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
		"str" : "testStr",
		"storage_config" : {
			"provider" : "aws",
			"location" : "us-west-2"
		},
		"storage_config_slice" : [{
			"provider" : "aws",
			"location" : "us-west-2"
		},
		{
			"provider" : "aws",
			"location" : "us-east-1"
		}],
		"storage_config_map": {
			"one": {
				"provider": "aws",
				"location": "us-west-2"
			}
		},
		"str_slice": [ "hello" ],
		"map_str_str": { "key": "value" }
	}`

func (s *RemoteConfigSuite) TestValidateConfigWithReflection() {
	c := s.buildValidSampleConfig()
	err := validateConfigWithReflection(c)
	s.Nil(err)
}

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionWithOptional() {
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

	storageProvider := VALID_REMOTE_CONFIG_STORAGE_CONFIG_PROVIDER
	storageLocation := VALID_REMOTE_CONFIG_STORAGE_CONFIG_LOCATION
	storageConfig := &StorageConfig{
		Provider: &storageProvider,
		Location: &storageLocation,
	}

	c := &SampleConfig{
		SQSQueueOptional:           sqsQueue,
		SQSClientOptional:          sqsClient,
		DynamoDBTableOptional:      dynamodbTable,
		DynamoDBClientOptional:     dynamodbClient,
		StrOptional:                &str,
		StorageConfigOptional:      storageConfig,
		StorageConfigSliceOptional: []*StorageConfig{storageConfig},
		SQSQueue:                   sqsQueue,
		SQSClient:                  sqsClient,
		DynamoDBTable:              dynamodbTable,
		DynamoDBClient:             dynamodbClient,
		Str:                        &str,
		StorageConfig:              storageConfig,
		StorageConfigSlice:         []*StorageConfig{storageConfig},
		StorageConfigMap:           map[string]*StorageConfig{"one": storageConfig},
		MapStrStr:                  map[string]*string{"hello": &str},
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorSQSQueueConfigNotSet() {
	c := &SampleConfig{
		SQSQueue: nil,
	}
	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: SQSQueue, not set"), err)
}

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorSQSClientConfigNotSet() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorSQSQueueConfigValidate() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorDynamoDBTableConfigNotSet() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorDynamoDBClientConfigNotSet() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorDynamoDBClientConfigValidate() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorDynamoDBTableConfigValidate() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorStrNotSet() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorStrEmpty() {
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

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorStorageConfigNotSet() {
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
		StorageConfig:  nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: StorageConfig, not set"), err)
}

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorStorageConfigSliceNotSet() {
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

	storageProvider := VALID_REMOTE_CONFIG_STORAGE_CONFIG_PROVIDER
	storageLocation := VALID_REMOTE_CONFIG_STORAGE_CONFIG_LOCATION
	storageConfig := &StorageConfig{
		Provider: &storageProvider,
		Location: &storageLocation,
	}

	c := &SampleConfig{
		SQSQueue:           sqsQueue,
		SQSClient:          sqsClient,
		DynamoDBTable:      dynamodbTable,
		DynamoDBClient:     dynamodbClient,
		Str:                &str,
		StorageConfig:      storageConfig,
		StorageConfigSlice: nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: StorageConfigSlice, not set"), err)
}

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorStorageConfigSliceNested() {
	c := s.buildValidSampleConfig()
	storageProvider := VALID_REMOTE_CONFIG_STORAGE_CONFIG_PROVIDER
	invalidStorageLocation := StorageLocation("")
	c.StorageConfigSlice = []*StorageConfig{
		&StorageConfig{
			Provider: &storageProvider,
			Location: &invalidStorageLocation,
		},
	}

	err := validateConfigWithReflection(c)
	s.NotNil(err)
	s.Equal(errors.New("Validater Field: StorageConfig, failed to validate with error, Region cannot be empty"), err)
}

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorStorageConfigMapNested() {
	c := s.buildValidSampleConfig()
	storageProvider := VALID_REMOTE_CONFIG_STORAGE_CONFIG_PROVIDER
	invalidStorageLocation := StorageLocation("")
	c.StorageConfigMap = map[string]*StorageConfig{
		"one": &StorageConfig{
			Provider: &storageProvider,
			Location: &invalidStorageLocation,
		},
	}

	err := validateConfigWithReflection(c)
	s.NotNil(err)
	s.Equal(errors.New("Sub field of StorageConfigMap with key 'one' failed to validated with error, Validater Field: StorageConfig, failed to validate with error, Region cannot be empty"), err)
}

func (s *RemoteConfigSuite) TestValidateConfigWithReflectionErrorStorageConfigSliceEmpty() {
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

	storageProvider := VALID_REMOTE_CONFIG_STORAGE_CONFIG_PROVIDER
	storageLocation := VALID_REMOTE_CONFIG_STORAGE_CONFIG_LOCATION
	storageConfig := &StorageConfig{
		Provider: &storageProvider,
		Location: &storageLocation,
	}

	c := &SampleConfig{
		SQSQueue:           sqsQueue,
		SQSClient:          sqsClient,
		DynamoDBTable:      dynamodbTable,
		DynamoDBClient:     dynamodbClient,
		Str:                &str,
		StorageConfig:      storageConfig,
		StorageConfigSlice: []*StorageConfig{},
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Slice Field: StorageConfigSlice, is empty"), err)
}

func (s *RemoteConfigSuite) TestLoadConfigFromURL_Gold() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, validConfigJSON)
	}))
	defer ts.Close()

	c := &SampleConfig{}
	err := LoadConfigFromURL(ts.URL, c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestLoadConfigFromURL_NotOK() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not Found")
	}))
	defer ts.Close()

	c := &SQSQueueConfig{}
	err := LoadConfigFromURL(ts.URL, c)
	assert.NotNil(s.T(), err)
	assert.Regexp(s.T(), regexp.MustCompile("Request to 'http://\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d{1,5}' returned non-200 OK status '404: Not Found'"), err.Error())
}

func (s *RemoteConfigSuite) TestLoadConfigFromURLError() {
	c := &SQSQueueConfig{}
	err := LoadConfigFromURL("invalid", c)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "Get invalid: unsupported protocol scheme \"\"")
}

func (s *RemoteConfigSuite) TestReadJSONValidate() {
	cfgBuffer := bytes.NewBufferString(validConfigJSON)

	c := &SampleConfig{}
	err := ReadJSONValidate(cfgBuffer, c)
	assert.Nil(s.T(), err)
}

func (s *RemoteConfigSuite) TestReadJSONParseEmbeddedStruct() {
	cfgBuffer := bytes.NewBufferString(validConfigJSON)

	c := &SampleConfig{}
	err := ReadJSONValidate(cfgBuffer, c)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), c.EmbeddedStr)
	assert.NotNil(s.T(), c.EmbeddedInt)
	assert.EqualValues(s.T(), "abc", *c.EmbeddedStr)
	assert.EqualValues(s.T(), 123, *c.EmbeddedInt)
}

func (s *RemoteConfigSuite) TestReadJSONValidateEmbeddedStruct() {
	invalidConfigJSON := strings.Replace(validConfigJSON, `"embedded_int": 123`, `"embedded_int": "123"`, 1)
	cfgBuffer := bytes.NewBufferString(invalidConfigJSON)

	c := &SampleConfig{}
	err := ReadJSONValidate(cfgBuffer, c)
	assert.EqualError(s.T(), err, "Failed to decode JSON, with error, json: cannot unmarshal string into Go value of type int64")
}

func (s *RemoteConfigSuite) TestReadJSONValidateInvalidJSON() {
	cfgBuffer := bytes.NewBufferString("Not JSON")

	c := &SampleConfig{}
	err := ReadJSONValidate(cfgBuffer, c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Failed to decode JSON, with error, invalid character 'N' looking for beginning of value"), err)
}

func (s *RemoteConfigSuite) TestReadJSONValidateErrorValidation() {
	cfgBuffer := bytes.NewBufferString("{}")

	c := &SampleConfig{}
	err := ReadJSONValidate(cfgBuffer, c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: SQSQueue, not set"), err)
}

func (s *RemoteConfigSuite) TestReadJSONValidateErrorInvalidJSON() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is NOT JSON")
	}))
	defer ts.Close()

	resp, _ := http.Get(ts.URL)
	defer resp.Body.Close()

	c := &SampleConfig{}
	err := ReadJSONValidate(resp.Body, c)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Failed to decode JSON, with error, invalid character 'T' looking for beginning of value"), err)
}

func (s *RemoteConfigSuite) buildValidSampleConfig() *SampleConfig {
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

	storageProvider := VALID_REMOTE_CONFIG_STORAGE_CONFIG_PROVIDER
	storageLocation := VALID_REMOTE_CONFIG_STORAGE_CONFIG_LOCATION
	storageConfig := &StorageConfig{
		Provider: &storageProvider,
		Location: &storageLocation,
	}

	return &SampleConfig{
		SQSQueue:           sqsQueue,
		SQSClient:          sqsClient,
		DynamoDBTable:      dynamodbTable,
		DynamoDBClient:     dynamodbClient,
		Str:                &str,
		StorageConfig:      storageConfig,
		StorageConfigSlice: []*StorageConfig{storageConfig},
		StorageConfigMap:   map[string]*StorageConfig{"one": storageConfig},
		MapStrStr:          map[string]*string{"hello": &str},
	}
}
