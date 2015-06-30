package remoteconfig

import "fmt"

type SQSQueueConfig struct {
	Region       *AWSRegion `json:"region,omitempty"`
	AWSAccountID *string    `json:"aws_account_id,omitempty"`
	QueueName    *string    `json:"queue_name,omitempty"`
}

// Returns a full SQS queue URL.
func (s SQSQueueConfig) GetURL(endpoint string) string {
	if endpoint == "" {
		return fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", *s.Region, *s.AWSAccountID, *s.QueueName)
	}
	return fmt.Sprintf("%s/%s/%s", endpoint, *s.AWSAccountID, *s.QueueName)
}
