package remoteconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AWSRegionSuite struct {
	suite.Suite
}

func TestAWSRegionSuite(t *testing.T) {
	suite.Run(t, new(AWSRegionSuite))
}

func (s *AWSRegionSuite) SetupSuite() {
}

func (s *AWSRegionSuite) SetupTest() {
}

func (s *AWSRegionSuite) TestValidateUSEast1() {
	r := AWS_REGION_US_EAST_1
	err := r.Validate()
	assert.Nil(s.T(), err)
}

func (s *AWSRegionSuite) TestValidateErrorNoRegion() {
	r := AWSRegion("")
	err := r.Validate()
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAWSRegionEmptyString, err)
}

func (s *AWSRegionSuite) TestValidateErrorInvalidRegion() {
	r := AWSRegion("invalidregion")
	err := r.Validate()
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAWSRegionInvalid, err)
}

func (s *AWSRegionSuite) TestUnmarshalText() {
	r := AWSRegion("")
	d := []byte(AWS_REGION_US_EAST_1)
	err := r.UnmarshalText(d)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), AWS_REGION_US_EAST_1, r)
}

func (s *AWSRegionSuite) TestUnmarshalErrorNoRegion() {
	r := AWSRegion("")
	d := []byte("")
	err := r.UnmarshalText(d)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAWSRegionEmptyString, err)
}

func (s *AWSRegionSuite) TestUnmarshalErrorInvalidRegion() {
	r := AWSRegion("")
	d := []byte("invalidregion")
	err := r.UnmarshalText(d)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAWSRegionInvalid, err)
}
