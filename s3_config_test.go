package remoteconfig

import (
	"errors"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_S3_CONFIG_ENDPOINT  string    = "http://localhost:9500/s3"
	VALID_S3_CONFIG_BUCKET    string    = "bucket"
	VALID_S3_CONFIG_REGION    AWSRegion = AWS_REGION_US_WEST_2
	VALID_S3_CONFIG_EXPIRY    uint      = 30
	VALID_S3_CONFIG_TEST_PATH string    = "test/path.path"
)

var (
	VALID_S3_CONFIG_S3_SCHEME_URL string = fmt.Sprintf("s3://%s/%s", VALID_S3_CONFIG_BUCKET, VALID_S3_CONFIG_TEST_PATH)
)

type S3ConfigSuite struct {
	suite.Suite
	assert *require.Assertions
}

func TestS3ConfigSuite(t *testing.T) {
	suite.Run(t, new(S3ConfigSuite))
}

func (s *S3ConfigSuite) SetupSuite() {
	s.assert = require.New(s.T())
}

func (s *S3ConfigSuite) TestValidate() {
	bucket := VALID_S3_CONFIG_BUCKET
	region := VALID_S3_CONFIG_REGION

	c := &S3Config{
		Bucket: &bucket,
		Region: &region,
	}

	err := validateConfigWithReflection(c)
	s.Nil(err)
}

func (s *S3ConfigSuite) TestValidateWithOptional() {
	endpoint := VALID_S3_CONFIG_ENDPOINT
	bucket := VALID_S3_CONFIG_BUCKET
	region := VALID_S3_CONFIG_REGION
	expiry := VALID_S3_CONFIG_EXPIRY

	c := &S3Config{
		Endpoint: &endpoint,
		Bucket:   &bucket,
		Region:   &region,
		Expiry:   &expiry,
	}

	err := validateConfigWithReflection(c)
	s.Nil(err)
}

func (s *S3ConfigSuite) TestValidateErrorEndpoint() {
	endpoint := ""

	c := &S3Config{
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(c)
	s.NotNil(err)
	s.Equal(errors.New("String Field: Endpoint, contains an empty string"), err)
}

func (s *S3ConfigSuite) TestValidateErrorBucketNotSet() {
	c := &S3Config{
		Bucket: nil,
	}

	err := validateConfigWithReflection(c)
	s.NotNil(err)
	s.Equal(errors.New("Field: Bucket, not set"), err)
}

func (s *S3ConfigSuite) TestValidateErrorBucketEmpty() {
	bucket := ""
	c := &S3Config{
		Bucket: &bucket,
	}

	err := validateConfigWithReflection(c)
	s.NotNil(err)
	s.Equal(errors.New("String Field: Bucket, contains an empty string"), err)
}

func (s *S3ConfigSuite) TestValidateErrorRegionNotSet() {
	bucket := VALID_S3_CONFIG_BUCKET

	c := &S3Config{
		Bucket: &bucket,
		Region: nil,
	}

	err := validateConfigWithReflection(c)
	s.NotNil(err)
	s.Equal(errors.New("Field: Region, not set"), err)
}

func (s *S3ConfigSuite) TestValidateErrorRegionInvalid() {
	bucket := VALID_S3_CONFIG_BUCKET
	region := AWSRegion("invalidregion")

	c := &S3Config{
		Bucket: &bucket,
		Region: &region,
	}

	err := validateConfigWithReflection(c)
	s.NotNil(err)
	s.Equal(errors.New("Validater Field: Region, failed to validate with error, Region is invalid"), err)
}

func (s *S3ConfigSuite) TestGetEndpointNotSet() {
	bucket := VALID_S3_CONFIG_BUCKET
	region := VALID_S3_CONFIG_REGION

	c := &S3Config{
		Endpoint: nil,
		Bucket:   &bucket,
		Region:   &region,
	}

	cEndpoint := c.GetEndpoint()
	s.Equal("", cEndpoint)
}

func (s *S3ConfigSuite) TestS3URLToConfig() {
	s3ConfigURL, path, err := S3URLToConfig(VALID_S3_CONFIG_S3_SCHEME_URL)
	s.Nil(err)
	s.NotNil(s3ConfigURL)
	s.NotEmpty(path)

	bucket := VALID_S3_CONFIG_BUCKET
	s3ConfigExpected := &S3Config{
		Bucket: &bucket,
	}
	s.Equal(s3ConfigExpected, s3ConfigURL)
	s.Equal(VALID_S3_CONFIG_TEST_PATH, path)
}

func (s *S3ConfigSuite) TestS3URLToConfigErrorURLParse() {
	s3ConfigURL, path, err := S3URLToConfig("invalid%6")
	s.Nil(s3ConfigURL)
	s.Empty(path)
	s.NotNil(err)
	s.IsType(&url.Error{}, err)
}

func (s *S3ConfigSuite) TestS3URLToConfigErrorURLScheme() {
	s3ConfigURL, path, err := S3URLToConfig("s4://invalidscheme")
	s.Nil(s3ConfigURL)
	s.Empty(path)
	s.NotNil(err)
	s.Equal(errors.New("URL does not have the s3:// scheme"), err)
}
