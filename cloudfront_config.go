package remoteconfig

type CloudfrontConfig struct {
	Region         *AWSRegion `json:"region,omitempty"`
	Endpoint       *string    `json:"endpoint,omitempty" remoteconfig:"optional"`
	AWSAccountID   *string    `json:"aws_account_id,omitempty"`
	DistributionID *string    `json:"distribution_id,omitempty"`
	BasePath       *string    `json:"base_path,omitempty"`
}

func (c *CloudfrontConfig) GetRegion() AWSRegion {
	return *c.Region
}

func (c *CloudfrontConfig) GetEndpoint() string {
	return *c.Endpoint
}
