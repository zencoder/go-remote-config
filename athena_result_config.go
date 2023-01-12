package remoteconfig

type AthenaQueryConfig struct {
	OutputBucket *string `json:"output_bucket,omitempty"`
}

func (s AthenaQueryConfig) GetOutputBucket() string {
	return *s.OutputBucket
}
