package remoteconfig

import (
	"errors"
	"fmt"
)

type SQSConfig struct {
	Region       *AWSRegion `json:"region,omitempty"`
	AWSAccountID *string    `json:"aws_account_id,omitempty"`
	QueueName    *string    `json:"queue_name,omitempty"`
}

var (
	ErrSQSConfigInvalidRegion       error = errors.New("Invalid SQS Region")
	ErrSQSConfigInvalidAWSAccountID error = errors.New("Invalid SQS AWS Account ID")
	ErrSQSConfigInvalidQueueName    error = errors.New("Invalid SQS Queue Name")
)

// Validates that all the member fields are valid.
func (s SQSConfig) Validate() error {
	if s.Region == nil || s.Region.Validate() != nil {
		return ErrSQSConfigInvalidRegion
	}
	if s.AWSAccountID == nil || *s.AWSAccountID == "" {
		return ErrSQSConfigInvalidAWSAccountID
	}
	if s.QueueName == nil || *s.QueueName == "" {
		return ErrSQSConfigInvalidQueueName
	}
	return nil
}

// Returns a full SQS queue URL.
func (s SQSConfig) GetURL() (string, error) {
	url := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", *s.Region, *s.AWSAccountID, *s.QueueName)
	return url, nil
}
