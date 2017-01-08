package remoteconfig

type SimpleWorkflowClientConfig struct {
	Region     *AWSRegion `json:"region,omitempty"`
	Endpoint   *string    `json:"endpoint,omitempty" remoteconfig:"optional"`
	DisableSSL *bool      `json:"disable_ssl,omit" remoteconfig:"optional"`
}

func (s SimpleWorkflowClientConfig) GetRegion() AWSRegion {
	return *s.Region
}

func (s SimpleWorkflowClientConfig) GetEndpoint() string {
	if s.Endpoint != nil {
		return *s.Endpoint
	}
	return ""
}

func (s SimpleWorkflowClientConfig) GetDisableSSL() bool {
	if s.DisableSSL != nil {
		return *s.DisableSSL
	}
	return false
}
