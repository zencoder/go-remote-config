package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_DYNAMODB_TABLE_TABLENAME string = "testTable"
)

type DynamoDBTableConfigSuite struct {
	suite.Suite
}

func TestDynamoDBTableConfigSuite(t *testing.T) {
	suite.Run(t, new(DynamoDBTableConfigSuite))
}

func (s *DynamoDBTableConfigSuite) SetupSuite() {
}

func (s *DynamoDBTableConfigSuite) SetupTest() {
}

func (s *DynamoDBTableConfigSuite) TestValidate() {
	tableName := VALID_DYNAMODB_TABLE_TABLENAME

	d := &DynamoDBTableConfig{
		TableName: &tableName,
	}

	err := validateConfigWithReflection(d)
	assert.Nil(s.T(), err)
}

func (s *DynamoDBTableConfigSuite) TestValidateErrorTableName() {
	tableName := ""

	d := &DynamoDBTableConfig{
		TableName: &tableName,
	}

	err := validateConfigWithReflection(d)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: TableName, contains an empty string"), err)
}

func (s *DynamoDBTableConfigSuite) TestGetTableName() {
	tableName := VALID_DYNAMODB_TABLE_TABLENAME

	d := &DynamoDBTableConfig{
		TableName: &tableName,
	}

	assert.Equal(s.T(), VALID_DYNAMODB_TABLE_TABLENAME, d.GetTableName())
}
