package remoteconfig

type AthenaClientConfig struct {
	Region *AWSRegion `json:"region,omitempty"`
}

func (s AthenaClientConfig) GetRegion() AWSRegion {
	return *s.Region
}
