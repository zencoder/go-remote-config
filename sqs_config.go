package remoteconfig

import (
	"errors"
	"fmt"
)

type SQSConfig struct {
	region       *AWSRegion `json:"region,omitempty"`
	awsAccountID *string    `json:"aws_account_id,omitempty"`
	queueName    *string    `json:"queue_name,omitempty"`
}

var (
	ErrSQSConfigInvalidRegion       error = errors.New("Invalid SQS Region")
	ErrSQSConfigInvalidAWSAccountID error = errors.New("Invalid SQS AWS Account ID")
	ErrSQSConfigInvalidQueueName    error = errors.New("Invalid SQS Queue Name")
)

// Validates that all the member fields are valid.
func (s SQSConfig) Validate() error {
	if s.region == nil || s.region.Validate() != nil {
		return ErrSQSConfigInvalidRegion
	}
	if s.awsAccountID == nil || *s.awsAccountID == "" {
		return ErrSQSConfigInvalidAWSAccountID
	}
	if s.queueName == nil || *s.queueName == "" {
		return ErrSQSConfigInvalidQueueName
	}
	return nil
}

// Returns a full SQS queue URL.
func (s SQSConfig) URL() (string, error) {
	url := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", *s.region, *s.awsAccountID, *s.queueName)
	return url, nil
}
