package remoteconfig

const (
	S3_ENDPOINT_EXPIRY_CONFIG_DEFAULT_EXPIRY uint = 60
)

type S3EndpointExpiryConfig struct {
	Endpoint *string `json:"endpoint,omitempty" yaml:"endpoint,omitempty" remoteconfig:"optional"`
	Expiry   *uint   `json:"expiry,omitempty" yaml:"expiry,omitempty" remoteconfig:"optional"`
}

func (c S3EndpointExpiryConfig) GetEndpoint() string {
	if c.Endpoint != nil {
		return *c.Endpoint
	}
	return ""
}

func (c S3EndpointExpiryConfig) GetExpiry() uint {
	if c.Expiry != nil {
		return *c.Expiry
	}
	return S3_ENDPOINT_EXPIRY_CONFIG_DEFAULT_EXPIRY
}
