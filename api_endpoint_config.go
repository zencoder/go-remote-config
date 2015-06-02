package remoteconfig

import "fmt"

type APIEndpointConfig struct {
	BasePath *string `json:"base_path,omitempty"`
	SubPath  *string `json:"sub_path,omitempty"`
}

func (c APIEndpointConfig) GetFullPath() string {
	return fmt.Sprintf("%s/%s", *c.BasePath, *c.SubPath)
}

func (c APIEndpointConfig) GetFullPathWithID(id string) string {
	return fmt.Sprintf("%s/%s/%s", *c.BasePath, *c.SubPath, id)
}
