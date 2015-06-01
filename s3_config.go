package remoteconfig

import "fmt"

const (
	S3_CONFIG_DEFAULT_EXPIRY uint = 60
)

type S3Config struct {
	Endpoint   *string    `json:"endpoint,omitempty" remoteconfig:"optional"`
	BaseBucket *string    `json:"base_bucket,omitempty"`                    // i.e. base-bucket
	Region     *AWSRegion `json:"region,omitempty"`                         // i.e. us-west-2
	FileExt    *string    `json:"file_ext,omitempty"`                       // i.e. json
	Expiry     *uint      `json:"expiry,omitempty" remoteconfig:"optional"` // i.e. 60
}

func (c S3Config) GetEndpoint() string {
	if c.Endpoint != nil {
		return *c.Endpoint
	}
	return ""
}

func (c S3Config) GetFullBucketName() string {
	// i.e. base-bucket-us-west-2
	return fmt.Sprintf("%s-%s", *c.BaseBucket, *c.Region)
}

func (c S3Config) GetFullPath(basePath string) string {
	// i.e. basepath.json
	return fmt.Sprintf("%s.%s", basePath, *c.FileExt)
}

func (c S3Config) GetExpiry() uint {
	if c.Expiry != nil {
		return *c.Expiry
	}
	return S3_CONFIG_DEFAULT_EXPIRY
}
