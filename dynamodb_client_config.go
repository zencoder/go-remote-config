package remoteconfig

type DynamoDBClientConfig struct {
	Region   *AWSRegion `json:"region,omitempty"`
	Endpoint *string    `json:"endpoint,omitempty" remoteconfig:"optional"`
}

func (d DynamoDBClientConfig) GetRegion() AWSRegion {
	return *d.Region
}

func (d DynamoDBClientConfig) GetEndpoint() string {
	return *d.Endpoint
}
