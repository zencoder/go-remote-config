package remoteconfig

import (
	"github.com/aws/aws-sdk-go/service/swf"
)

type SimpleWorkflowConfig struct {
	Domain       *string       `json:"domain"`
	WorkflowType *WorkflowType `json:"workflow_type"`
}

func (s SimpleWorkflowConfig) GetDomain() string {
	return *s.Domain
}

func (s SimpleWorkflowConfig) GetWorkflowType() WorkflowType {
	return *s.WorkflowType
}

type WorkflowType struct {
	Name    *string `json:"name"`
	Version *string `json:"version"`
}

func (w WorkflowType) ToAWS() swf.WorkflowType {
	return swf.WorkflowType{Name: w.Name, Version: w.Version}
}
