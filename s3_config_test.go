package remoteconfig

import (
	"errors"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_S3_CONFIG_ENDPOINT    string    = "http://localhost:9500/s3"
	VALID_S3_CONFIG_BASE_BUCKET string    = "base-bucket"
	VALID_S3_CONFIG_REGION      AWSRegion = AWS_REGION_US_WEST_2
	VALID_S3_CONFIG_FILE_EXT    string    = "json"
	VALID_S3_CONFIG_EXPIRY      uint      = 30
	VALID_S3_CONFIG_TEST_PATH   string    = "test/path"
)

var (
	VALID_S3_CONFIG_FULL_BUCKET_NAME string = fmt.Sprintf("%s-%s", VALID_S3_CONFIG_BASE_BUCKET, VALID_S3_CONFIG_REGION)
	VALID_S3_CONFIG_FULL_PATH        string = fmt.Sprintf("%s.%s", VALID_S3_CONFIG_TEST_PATH, VALID_S3_CONFIG_FILE_EXT)
	VALID_S3_CONFIG_S3_SCHEME_URL    string = fmt.Sprintf("s3://%s-%s/%s.%s", VALID_S3_CONFIG_BASE_BUCKET, VALID_S3_CONFIG_REGION, VALID_S3_CONFIG_TEST_PATH, VALID_S3_CONFIG_FILE_EXT)
)

type S3ConfigSuite struct {
	suite.Suite
}

func TestS3ConfigSuite(t *testing.T) {
	suite.Run(t, new(S3ConfigSuite))
}

func (s *S3ConfigSuite) SetupSuite() {
}

func (s *S3ConfigSuite) SetupTest() {
}

func (s *S3ConfigSuite) TestValidate() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *S3ConfigSuite) TestValidateWithOptional() {
	endpoint := VALID_S3_CONFIG_ENDPOINT
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT
	expiry := VALID_S3_CONFIG_EXPIRY

	c := &S3Config{
		Endpoint:   &endpoint,
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
		Expiry:     &expiry,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *S3ConfigSuite) TestValidateErrorEndpoint() {
	endpoint := ""

	c := &S3Config{
		Endpoint: &endpoint,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Endpoint, contains an empty string"), err)
}

func (s *S3ConfigSuite) TestValidateErrorBaseBucketNotSet() {
	c := &S3Config{
		BaseBucket: nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: BaseBucket, not set"), err)
}

func (s *S3ConfigSuite) TestValidateErrorBaseBucketEmpty() {
	baseBucket := ""
	c := &S3Config{
		BaseBucket: &baseBucket,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: BaseBucket, contains an empty string"), err)
}

func (s *S3ConfigSuite) TestValidateErrorRegionNotSet() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     nil,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Field: Region, not set"), err)
}

func (s *S3ConfigSuite) TestValidateErrorRegionInvalid() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := AWSRegion("invalidregion")

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: Region, failed to validate with error, Region is invalid"), err)
}

func (s *S3ConfigSuite) TestValidateNoErrorFileExtNotSet() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    nil,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *S3ConfigSuite) TestValidateErrorFileExtEmpty() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := ""

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: FileExt, contains an empty string"), err)
}

func (s *S3ConfigSuite) TestGetEndpointSet() {
	endpoint := VALID_S3_CONFIG_ENDPOINT
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT

	c := &S3Config{
		Endpoint:   &endpoint,
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}

	cEndpoint := c.GetEndpoint()
	assert.Equal(s.T(), VALID_S3_CONFIG_ENDPOINT, cEndpoint)
}

func (s *S3ConfigSuite) TestGetEndpointNotSet() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT

	c := &S3Config{
		Endpoint:   nil,
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}

	cEndpoint := c.GetEndpoint()
	assert.Equal(s.T(), "", cEndpoint)
}

func (s *S3ConfigSuite) TestGetFullBucketName() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}

	fullBucketName := c.GetFullBucketName()
	assert.Equal(s.T(), VALID_S3_CONFIG_FULL_BUCKET_NAME, fullBucketName)
}

func (s *S3ConfigSuite) TestGetFullPath() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}

	fullPath := c.GetFullPath(VALID_S3_CONFIG_TEST_PATH)
	assert.Equal(s.T(), VALID_S3_CONFIG_FULL_PATH, fullPath)
}

func (s *S3ConfigSuite) TestGetRegion() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}

	cRegion := c.GetRegion()
	assert.Equal(s.T(), VALID_S3_CONFIG_REGION, cRegion)
}

func (s *S3ConfigSuite) TestGetExpirySet() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT
	expiry := VALID_S3_CONFIG_EXPIRY

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
		Expiry:     &expiry,
	}

	cExpiry := c.GetExpiry()
	assert.Equal(s.T(), VALID_S3_CONFIG_EXPIRY, cExpiry)
}

func (s *S3ConfigSuite) TestGetExpiryNotSet() {
	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT

	c := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
		Expiry:     nil,
	}

	cExpiry := c.GetExpiry()
	assert.Equal(s.T(), S3_CONFIG_DEFAULT_EXPIRY, cExpiry)
}

func (s *S3ConfigSuite) TestS3URLToConfig() {
	s3ConfigURL, path, err := S3URLToConfig(VALID_S3_CONFIG_S3_SCHEME_URL)
	assert.NotNil(s.T(), s3ConfigURL)
	assert.NotEmpty(s.T(), path)
	assert.Nil(s.T(), err)

	baseBucket := VALID_S3_CONFIG_BASE_BUCKET
	region := VALID_S3_CONFIG_REGION
	fileExt := VALID_S3_CONFIG_FILE_EXT
	s3ConfigExpected := &S3Config{
		BaseBucket: &baseBucket,
		Region:     &region,
		FileExt:    &fileExt,
	}
	assert.Equal(s.T(), s3ConfigExpected, s3ConfigURL)
	assert.Equal(s.T(), VALID_S3_CONFIG_TEST_PATH, path)
}

func (s *S3ConfigSuite) TestS3URLToConfigErrorURLParse() {
	s3ConfigURL, path, err := S3URLToConfig("invalid%6")
	assert.Nil(s.T(), s3ConfigURL)
	assert.Empty(s.T(), path)
	assert.NotNil(s.T(), err)
	assert.IsType(s.T(), &url.Error{}, err)
}

func (s *S3ConfigSuite) TestS3URLToConfigErrorURLScheme() {
	s3ConfigURL, path, err := S3URLToConfig("s4://invalidscheme")
	assert.Nil(s.T(), s3ConfigURL)
	assert.Empty(s.T(), path)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("URL does not have the s3:// scheme"), err)
}

func (s *S3ConfigSuite) TestS3URLToConfigErrorInvalidRegion() {
	s3ConfigURL, path, err := S3URLToConfig("s3://base-bucket-invalid-region")
	assert.Nil(s.T(), s3ConfigURL)
	assert.Empty(s.T(), path)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Region not found in bucket name, base-bucket-invalid-region"), err)
}

func (s *S3ConfigSuite) TestS3URLToConfigNoExt() {
	assert := assert.New(s.T())
	config, key, err := S3URLToConfig("s3://foo-eu-central-1/bar/bat")
	assert.Nil(err)
	assert.NotNil(config)
	assert.Equal("", *config.FileExt)
	assert.NotEmpty(key)
}
