package remoteconfig

type SQSClientConfig struct {
	Region   *AWSRegion `json:"region,omitempty"`
	Endpoint *string    `json:"endpoint,omitempty" remoteconfig:"optional"`
}

func (s SQSClientConfig) GetRegion() AWSRegion {
	return *s.Region
}

func (s SQSClientConfig) GetEndpoint() string {
	return *s.Endpoint
}
