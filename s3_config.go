package remoteconfig

import (
	"errors"
	"net/url"
)

const (
	S3_CONFIG_DEFAULT_EXPIRY uint = 60
)

type S3Config struct {
	Endpoint *string    `json:"endpoint,omitempty" yaml:"endpoint,omitempty" remoteconfig:"optional"`
	Bucket   *string    `json:"bucket,omitempty" yaml:"bucket,omitempty"`                         // i.e. bucket
	Region   *AWSRegion `json:"region,omitempty" yaml:"region,omitempty"`                         // i.e. us-west-2
	Expiry   *uint      `json:"expiry,omitempty" yaml:"expiry,omitempty" remoteconfig:"optional"` // i.e. 60
}

func (c S3Config) GetEndpoint() string {
	if c.Endpoint != nil {
		return *c.Endpoint
	}
	return ""
}

func (c S3Config) GetExpiry() uint {
	if c.Expiry != nil {
		return *c.Expiry
	}
	return S3_CONFIG_DEFAULT_EXPIRY
}

func S3URLToConfig(s3URL string) (*S3Config, string, error) {
	// i.e. s3://bucket/test/path.json
	c := &S3Config{}

	pURL, err := url.Parse(s3URL)
	if err != nil {
		return nil, "", err
	}

	if pURL.Scheme != "s3" {
		return nil, "", errors.New("URL does not have the s3:// scheme")
	}

	c.Bucket = &pURL.Host

	return c, pURL.Path[1:], nil
}
