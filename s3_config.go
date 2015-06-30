package remoteconfig

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

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

func (c S3Config) GetRegion() AWSRegion {
	return *c.Region
}

func (c S3Config) GetExpiry() uint {
	if c.Expiry != nil {
		return *c.Expiry
	}
	return S3_CONFIG_DEFAULT_EXPIRY
}

func findAWSRegionForBucket(bucket string) (AWSRegion, error) {
	for _, r := range AWSRegions {
		if strings.HasSuffix(bucket, string(r)) {
			return r, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Region not found in bucket name, %s", bucket))
}

func S3URLToConfig(s3URL string) (*S3Config, string, error) {
	// i.e. s3://base-bucket-us-west-2/test/path.json
	c := &S3Config{}

	pURL, err := url.Parse(s3URL)
	if err != nil {
		return nil, "", err
	}
	if pURL.Scheme != "s3" {
		return nil, "", errors.New("URL does not have the s3:// scheme")
	}

	bucket := pURL.Host
	region, err := findAWSRegionForBucket(bucket)
	if err != nil {
		return nil, "", err
	}
	baseBucket := strings.TrimSuffix(bucket, "-"+string(region))

	key := strings.TrimPrefix(pURL.Path, "/")
	fileExt := filepath.Ext(key)[1:]

	key = key[:len(key)-len(fileExt)-1]

	c.BaseBucket = &baseBucket
	c.Region = &region
	c.FileExt = &fileExt

	return c, key, nil
}
