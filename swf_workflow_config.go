package remoteconfig

import (
	"github.com/aws/aws-sdk-go/service/swf"
)

type SimpleWorkflowConfig struct {
	Domain       *string           `json:"domain"`
	WorkflowType *swf.WorkflowType `json:"workflow_type"`
}

func (s SimpleWorkflowConfig) GetDomain() string {
	return *s.Domain
}

func (s SimpleWorkflowConfig) GetWorkflowType() swf.WorkflowType {
	return *s.WorkflowType
}
