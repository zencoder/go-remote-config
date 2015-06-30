package remoteconfig

type DynamoDBTableConfig struct {
	TableName *string `json:"table_name,omitempty"`
}

func (d DynamoDBTableConfig) GetTableName() string {
	return *d.TableName
}
