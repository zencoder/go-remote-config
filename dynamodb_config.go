package remoteconfig

import "errors"

type DynamoDBConfig struct {
	region    *AWSRegion `json:"region,omitempty"`
	tableName *string    `json:"table_name,omitempty"`
}

var (
	ErrDynamoDBConfigInvalidRegion    error = errors.New("Invalid DynamoDB Region")
	ErrDynamoDBConfigInvalidTableName error = errors.New("Invalid DynamoDB Table Name")
)

// Validates that all the member fields are valid.
func (d DynamoDBConfig) Validate() error {
	if d.region == nil || d.region.Validate() != nil {
		return ErrDynamoDBConfigInvalidRegion
	}
	if d.tableName == nil || *d.tableName == "" {
		return ErrDynamoDBConfigInvalidTableName
	}
	return nil
}

func (d DynamoDBConfig) Region() AWSRegion {
	return *d.region
}

func (d DynamoDBConfig) TableName() string {
	return *d.tableName
}
