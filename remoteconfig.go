package remoteconfig

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

const DEFAULT_S3_EXPIRY uint = 60

// Downloads a configuration JSON file from S3.
// Parses it to a particular struct type and runs a validation.
// URL should be of the format s3://bucket/path/file.json
func LoadConfigFromS3(configURL string, configRegion AWSRegion, configStruct interface{}) error {
	// Build a Signed URL to the config file in S3
	signedURL, err := BuildSignedS3URL(configURL, configRegion, DEFAULT_S3_EXPIRY)
	if err != nil {
		return err
	}

	return downloadJSONValidate(signedURL, configStruct)
}

func downloadJSONValidate(signedURL string, configStruct interface{}) error {
	// Download the config file from S3
	resp, err := http.Get(signedURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check that we got a valid response code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Download of JSON failed, URL = %s Response Code = %d", signedURL, resp.StatusCode)
	}

	// Do a streaming JSON decode
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(configStruct); err != nil {
		return fmt.Errorf("Failed to decode JSON, with error, %s", err.Error())
	}

	// Run validation on the config
	err = validateConfigWithReflection(configStruct)
	if err != nil {
		return err
	}

	return nil
}

// Validates a configuration struct.
// Uses reflection to determine and call the correct Validation methods for each type.
func validateConfigWithReflection(c interface{}) error {
	valueElem := reflect.ValueOf(c).Elem()
	typeElem := reflect.TypeOf(c).Elem()

	for i := 0; i < valueElem.NumField(); i++ {
		valueField := valueElem.Field(i)
		typeField := typeElem.Field(i)

		switch valueField.Interface().(type) {
		case *SQSConfig:
			sqs := valueField.Interface().(*SQSConfig)
			if sqs == nil {
				return fmt.Errorf("SQS Config Not Set, %s", typeField.Name)
			}
			if err := sqs.Validate(); err != nil {
				return fmt.Errorf("Failed to Validate SQSConfig, %s with error: %s", typeField.Name, err.Error())
			}
		case *DynamoDBConfig:
			dynamodb := valueField.Interface().(*DynamoDBConfig)
			if dynamodb == nil {
				return fmt.Errorf("DynamoDB Config Not Set, %s", typeField.Name)
			}
			if err := dynamodb.Validate(); err != nil {
				return fmt.Errorf("Failed to Validate DynamoDBConfig, %s with error: %s", typeField.Name, err.Error())
			}
		case *string:
			s := valueField.Interface().(*string)
			if s == nil || *s == "" {
				return fmt.Errorf("Empty String, %s", typeField.Name)
			}
		}
	}

	return nil
}
