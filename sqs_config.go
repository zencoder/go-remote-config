package remoteconfig

import "fmt"

type SQSConfig struct {
	Region       *AWSRegion `json:"region,omitempty"`
	AWSAccountID *string    `json:"aws_account_id,omitempty"`
	QueueName    *string    `json:"queue_name,omitempty"`
	Endpoint     *string    `json:"endpoint,omitempty" remoteconfig:"optional"`
}

// Returns a full SQS queue URL.
// If Endpoint is set it will be returned instead of building the URL.
func (s SQSConfig) GetURL() string {
	if s.Endpoint != nil && *s.Endpoint != "" {
		return *s.Endpoint
	}
	return fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", *s.Region, *s.AWSAccountID, *s.QueueName)
}
