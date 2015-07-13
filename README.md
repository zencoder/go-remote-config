# go-remote-config

[![godoc](https://godoc.org/github.com/zencoder/go-remote-config?status.svg)](http://godoc.org/github.com/zencoder/go-remote-config)
[![Circle CI](https://circleci.com/gh/zencoder/go-remote-config.svg?style=svg)](https://circleci.com/gh/zencoder/go-remote-config)
[![Coverage Status](https://coveralls.io/repos/zencoder/go-remote-config/badge.svg?branch=master&t=VFcsMv)](https://coveralls.io/r/zencoder/go-remote-config?branch=master)

A Go library for configuration management with JSON files in remote storage.

Install
-------
	go get github.com/zencoder/go-remote-config

Supported Storage Providers
-------
* AWS S3 (Signed URLs)
* HTTP/HTTPS

Features
-------
* Reflection based config validation
  * Required fields
  * Optional fields
  * Custom Validate interface
  * Empty string checks
  * Struct & Slice, nested support
* Built in config structs for services
  * AWS Regions
  * AWS DynamoDB (Client + Table)
  * AWS SQS (Client + Queue)
  * AWS S3
  * Generic HTTP Endpoints

Future Features
-------
* More storage provider support
  * Google Cloud Storage
  * Rackspace CloudFiles
* Default value support
* Config download retry support
* Live config reloading

Example
-------
```go
type SampleConfig struct {
	SQSQueueOptional           *SQSQueueConfig       `json:"sqs_queue_optional,omitempty" remoteconfig:"optional"`
	SQSClientOptional          *SQSClientConfig      `json:"sqs_client_optional,omitempty" remoteconfig:"optional"`
	DynamoDBTableOptional      *DynamoDBTableConfig  `json:"dynamodb_table_optional,omitempty" remoteconfig:"optional"`
	DynamoDBClientOptional     *DynamoDBClientConfig `json:"dynamodb_client_optional,omitempty" remoteconfig:"optional"`
	StrOptional                *string               `json:"str_optional,omitempty" remoteconfig:"optional"`
	StorageConfigOptional      *StorageConfig        `json:"storage_config_optional,omitempty" remoteconfig:"optional"`
	StorageConfigSliceOptional []*StorageConfig      `json:"storage_config_slice_optional,omitempty" remoteconfig:"optional"`
	SQSQueue                   *SQSQueueConfig       `json:"sqs_queue,omitempty"`
	SQSClient                  *SQSClientConfig      `json:"sqs_client,omitempty"`
	DynamoDBTable              *DynamoDBTableConfig  `json:"dynamodb_table,omitempty"`
	DynamoDBClient             *DynamoDBClientConfig `json:"dynamodb_client,omitempty"`
	Str                        *string               `json:"str,omitempty"`
	StorageConfig              *StorageConfig        `json:"storage_config,omitempty"`
	StorageConfigSlice         []*StorageConfig      `json:"storage_config_slice,omitempty"`
}

var s SampleConfig
LoadConfig(s)

import (
	"log"
	"os"

	"github.com/zencoder/go-remote-config"
)

func LoadConfig(config interface{}) {
	// Load the config from S3
	configURL := os.Getenv("S3_CONFIG_URL")
	configRegion := remoteconfig.AWSRegion(os.Getenv("S3_CONFIG_REGION"))

	// Load an endpoint for S3 config (can be used to fake out S3 for testing)
	configEndpoint := os.Getenv("S3_CONFIG_ENDPOINT")

	// We should fail out if config environment variables are not set / valid
	if configURL == "" {
		log.Panic("S3 Configuration URL must be provided.")
	}

	if err := configRegion.Validate(); err != nil {
		log.Panic("Invalid Region for S3 Configuration")
	}

	log.Printf("Loading config file from S3. URL = %s, Region = %s", configURL, configRegion)

	if err := remoteconfig.LoadConfigFromS3(configURL, configRegion, configEndpoint, config); err != nil {
		log.Panicf("Failed to load config file, with error: %s", err.Error())
	}

	log.Printf("Successfully loaded config file from S3. URL = %s, Region = %s", configURL, configRegion)
	log.Printf("%s", config)

}
```
