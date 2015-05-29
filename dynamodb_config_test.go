package remoteconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_DYNAMODB_REGION     AWSRegion = AWS_REGION_US_EAST_1
	VALID_DYNAMODB_TABLE_NAME string    = "testTable"
)

type DynamoDBConfigSuite struct {
	suite.Suite
}

func TestDynamoDBConfigSuite(t *testing.T) {
	suite.Run(t, new(DynamoDBConfigSuite))
}

func (s *DynamoDBConfigSuite) SetupSuite() {
}

func (s *DynamoDBConfigSuite) SetupTest() {
}

func (s *DynamoDBConfigSuite) TestValidate() {
	region := VALID_DYNAMODB_REGION
	tableName := VALID_DYNAMODB_TABLE_NAME

	d := &DynamoDBConfig{
		Region:    &region,
		TableName: &tableName,
	}

	err := d.Validate()
	assert.Nil(s.T(), err)
}

func (s *DynamoDBConfigSuite) TestValidateErrorRegion() {
	region := AWSRegion("")
	tableName := VALID_DYNAMODB_TABLE_NAME

	d := &DynamoDBConfig{
		Region:    &region,
		TableName: &tableName,
	}

	err := d.Validate()
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrDynamoDBConfigInvalidRegion, err)
}

func (s *DynamoDBConfigSuite) TestValidateErrorTableName() {
	region := VALID_DYNAMODB_REGION
	tableName := ""

	d := &DynamoDBConfig{
		Region:    &region,
		TableName: &tableName,
	}

	err := d.Validate()
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrDynamoDBConfigInvalidTableName, err)
}

func (s *DynamoDBConfigSuite) TestGetRegion() {
	region := VALID_DYNAMODB_REGION
	tableName := VALID_DYNAMODB_TABLE_NAME

	d := &DynamoDBConfig{
		Region:    &region,
		TableName: &tableName,
	}

	assert.Equal(s.T(), VALID_DYNAMODB_REGION, d.GetRegion())
}

func (s *DynamoDBConfigSuite) TestGetTableName() {
	region := VALID_DYNAMODB_REGION
	tableName := VALID_DYNAMODB_TABLE_NAME

	d := &DynamoDBConfig{
		Region:    &region,
		TableName: &tableName,
	}

	assert.Equal(s.T(), VALID_DYNAMODB_TABLE_NAME, d.GetTableName())
}
