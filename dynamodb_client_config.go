package remoteconfig

type DynamoDBClientConfig struct {
	Region     *AWSRegion `json:"region,omitempty"`
	Endpoint   *string    `json:"endpoint,omitempty" remoteconfig:"optional"`
	DisableSSL *bool      `json:"disable_ssl,omit" remoteconfig:"optional"`
}

func (d DynamoDBClientConfig) GetRegion() AWSRegion {
	return *d.Region
}

func (d DynamoDBClientConfig) GetEndpoint() string {
	if d.Endpoint != nil {
		return *d.Endpoint
	}
	return ""
}

func (d DynamoDBClientConfig) GetDisableSSL() bool {
	if d.DisableSSL != nil {
		return *d.DisableSSL
	}
	return false
}
