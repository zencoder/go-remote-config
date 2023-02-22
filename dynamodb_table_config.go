package remoteconfig

type DynamoDBTableConfig struct {
	TableName *string `mapstructure:"table_name" json:"table_name,omitempty"`
}

func (d DynamoDBTableConfig) GetTableName() string {
	return *d.TableName
}
