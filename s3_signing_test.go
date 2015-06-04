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
	VALID_S3_BUCKET           string    = "bucketname"
	VALID_S3_KEY              string    = "path/file.json"
	VALID_S3_REGION_US_EAST_1 AWSRegion = AWS_REGION_US_WEST_2
	VALID_S3_REGION_US_WEST_2 AWSRegion = AWS_REGION_US_WEST_2
	VALID_S3_EXPIRY           uint      = 60
	VALID_S3_NO_ENDPOINT      string    = ""
	VALID_S3_LOCAL_ENDPOINT   string    = "http://localhost:8500"
)

var (
	VALID_S3_URL string = fmt.Sprintf("s3://%s/%s", VALID_S3_BUCKET, VALID_S3_KEY)
)

type S3SigningSuite struct {
	suite.Suite
}

func TestS3SigningSuite(t *testing.T) {
	suite.Run(t, new(S3SigningSuite))
}

func (s *S3SigningSuite) SetupSuite() {
}

func (s *S3SigningSuite) SetupTest() {
}

func (s *S3SigningSuite) TestBuildSignedS3URLUSEast1() {
	signedURL, err := BuildSignedS3URL(VALID_S3_URL, VALID_S3_REGION_US_EAST_1, VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Nil(s.T(), err)

	urlRegexp := fmt.Sprintf("^https://s3-%s\\.amazonaws\\.com/bucketname/path/file.json\\?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=.*&X-Amz-Date=.*&X-Amz-Expires=%d&X-Amz-SignedHeaders=host&X-Amz-Signature=.*$", VALID_S3_REGION_US_EAST_1, VALID_S3_EXPIRY)
	assert.Regexp(s.T(), urlRegexp, signedURL)
}

func (s *S3SigningSuite) TestBuildSignedS3URLUSWest2() {
	signedURL, err := BuildSignedS3URL(VALID_S3_URL, VALID_S3_REGION_US_WEST_2, VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Nil(s.T(), err)

	urlRegexp := fmt.Sprintf("^https://s3-%s\\.amazonaws\\.com/bucketname/path/file.json\\?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=.*&X-Amz-Date=.*&X-Amz-Expires=%d&X-Amz-SignedHeaders=host&X-Amz-Signature=.*$", VALID_S3_REGION_US_WEST_2, VALID_S3_EXPIRY)
	assert.Regexp(s.T(), urlRegexp, signedURL)
}

func (s *S3SigningSuite) TestBuildSignedS3URLUSEast1WithEndpoint() {
	signedURL, err := BuildSignedS3URL(VALID_S3_URL, VALID_S3_REGION_US_EAST_1, VALID_S3_EXPIRY, VALID_S3_LOCAL_ENDPOINT)
	assert.Nil(s.T(), err)

	urlRegexp := fmt.Sprintf("^%s/bucketname/path/file.json\\?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=.*&X-Amz-Date=.*&X-Amz-Expires=%d&X-Amz-SignedHeaders=host&X-Amz-Signature=.*$", VALID_S3_LOCAL_ENDPOINT, VALID_S3_EXPIRY)
	assert.Regexp(s.T(), urlRegexp, signedURL)
}

func (s *S3SigningSuite) TestBuildSignedS3URLUSWest2WithEndpoint() {
	signedURL, err := BuildSignedS3URL(VALID_S3_URL, VALID_S3_REGION_US_WEST_2, VALID_S3_EXPIRY, VALID_S3_LOCAL_ENDPOINT)
	assert.Nil(s.T(), err)

	urlRegexp := fmt.Sprintf("^%s/bucketname/path/file.json\\?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=.*&X-Amz-Date=.*&X-Amz-Expires=%d&X-Amz-SignedHeaders=host&X-Amz-Signature=.*$", VALID_S3_LOCAL_ENDPOINT, VALID_S3_EXPIRY)
	assert.Regexp(s.T(), urlRegexp, signedURL)
}

func (s *S3SigningSuite) TestBuildSignedS3URLErrorURLParse() {
	signedURL, err := BuildSignedS3URL("invalid%6", VALID_S3_REGION_US_WEST_2, VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Equal(s.T(), "", signedURL)
	assert.NotNil(s.T(), err)
	assert.IsType(s.T(), &url.Error{}, err)
}

func (s *S3SigningSuite) TestBuildSignedS3URLErrorNotS3URL() {
	signedURL, err := BuildSignedS3URL("s4://invalidurl", VALID_S3_REGION_US_WEST_2, VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Equal(s.T(), "", signedURL)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("S3 URL does not start with the s3:// scheme"), err)
}

func (s *S3SigningSuite) TestBuildSignedS3URLErrorNoRegion() {
	signedURL, err := BuildSignedS3URL(VALID_S3_URL, "", VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Equal(s.T(), "", signedURL)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAWSRegionEmptyString, err)
}

func (s *S3SigningSuite) TestBuildSignedS3URLErrorInvalidRegion() {
	signedURL, err := BuildSignedS3URL(VALID_S3_URL, "invalidregion", VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Equal(s.T(), "", signedURL)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAWSRegionInvalid, err)
}

func (s *S3SigningSuite) TestgenerateSignedS3URLUSEast1() {
	signedURL, err := generateSignedS3URL(VALID_S3_REGION_US_EAST_1, VALID_S3_BUCKET, VALID_S3_KEY, VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Nil(s.T(), err)
	urlRegexp := fmt.Sprintf("^https://s3-%s\\.amazonaws\\.com/bucketname/path/file.json\\?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=.*&X-Amz-Date=.*&X-Amz-Expires=%d&X-Amz-SignedHeaders=host&X-Amz-Signature=.*$", VALID_S3_REGION_US_EAST_1, VALID_S3_EXPIRY)
	assert.Regexp(s.T(), urlRegexp, signedURL)
}

func (s *S3SigningSuite) TestgenerateSignedS3URLUSWest2() {
	signedURL, err := generateSignedS3URL(VALID_S3_REGION_US_WEST_2, VALID_S3_BUCKET, VALID_S3_KEY, VALID_S3_EXPIRY, VALID_S3_NO_ENDPOINT)
	assert.Nil(s.T(), err)
	urlRegexp := fmt.Sprintf("^https://s3-%s\\.amazonaws\\.com/bucketname/path/file.json\\?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=.*&X-Amz-Date=.*&X-Amz-Expires=%d&X-Amz-SignedHeaders=host&X-Amz-Signature=.*$", VALID_S3_REGION_US_WEST_2, VALID_S3_EXPIRY)
	assert.Regexp(s.T(), urlRegexp, signedURL)
}
