package remoteconfig

import "errors"

type DynamoDBConfig struct {
	Region    *AWSRegion `json:"region,omitempty"`
	TableName *string    `json:"table_name,omitempty"`
}

var (
	ErrDynamoDBConfigInvalidRegion    error = errors.New("Invalid DynamoDB Region")
	ErrDynamoDBConfigInvalidTableName error = errors.New("Invalid DynamoDB Table Name")
)

// Validates that all the member fields are valid.
func (d DynamoDBConfig) Validate() error {
	if d.Region == nil || d.Region.Validate() != nil {
		return ErrDynamoDBConfigInvalidRegion
	}
	if d.TableName == nil || *d.TableName == "" {
		return ErrDynamoDBConfigInvalidTableName
	}
	return nil
}

func (d DynamoDBConfig) GetRegion() AWSRegion {
	return *d.Region
}

func (d DynamoDBConfig) GetTableName() string {
	return *d.TableName
}
