package remoteconfig

import "fmt"

type SQSQueueConfig struct {
	Region       *AWSRegion `json:"region,omitempty"`
	AWSAccountID *string    `json:"aws_account_id,omitempty"`
	QueueName    *string    `json:"queue_name,omitempty"`
}

// Returns a full SQS queue URL.
func (s SQSQueueConfig) GetURL() string {
	return fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", *s.Region, *s.AWSAccountID, *s.QueueName)
}
