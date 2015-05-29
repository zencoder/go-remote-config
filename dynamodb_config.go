package remoteconfig

type DynamoDBConfig struct {
	Region    *AWSRegion `json:"region,omitempty"`
	TableName *string    `json:"table_name,omitempty"`
}

func (d DynamoDBConfig) GetRegion() AWSRegion {
	return *d.Region
}

func (d DynamoDBConfig) GetTableName() string {
	return *d.TableName
}
