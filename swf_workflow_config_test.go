package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_SWF_WORKFLOW_DOMAIN       string = "testDomain"
	VALID_SWF_WORKFLOW_TYPE_NAME    string = "testWorkflow"
	VALID_SWF_WORKFLOW_TYPE_VERSION string = "1.0"
)

type SimpleWorkflowConfigSuite struct {
	suite.Suite
}

func TestSimpleWorkflowConfigSuite(t *testing.T) {
	suite.Run(t, new(SimpleWorkflowConfigSuite))
}

func (s *SimpleWorkflowConfigSuite) SetupSuite() {
}

func (s *SimpleWorkflowConfigSuite) SetupTest() {
}

func (s *SimpleWorkflowConfigSuite) TestValidate() {
	domain := VALID_SWF_WORKFLOW_DOMAIN
	name := VALID_SWF_WORKFLOW_TYPE_NAME
	version := VALID_SWF_WORKFLOW_TYPE_VERSION
	workflowType := WorkflowType{Name: &name, Version: &version}

	c := &SimpleWorkflowConfig{
		Domain:       &domain,
		WorkflowType: &workflowType,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *SimpleWorkflowConfigSuite) TestValidateErrorDomain() {
	domain := ""
	name := VALID_SWF_WORKFLOW_TYPE_NAME
	version := VALID_SWF_WORKFLOW_TYPE_VERSION
	workflowType := WorkflowType{Name: &name, Version: &version}

	c := &SimpleWorkflowConfig{
		Domain:       &domain,
		WorkflowType: &workflowType,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("String Field: Domain, contains an empty string"), err)
}

func (s *SimpleWorkflowConfigSuite) TestGetDomain() {
	domain := VALID_SWF_WORKFLOW_DOMAIN
	name := VALID_SWF_WORKFLOW_TYPE_NAME
	version := VALID_SWF_WORKFLOW_TYPE_VERSION
	workflowType := WorkflowType{Name: &name, Version: &version}

	c := &SimpleWorkflowConfig{
		Domain:       &domain,
		WorkflowType: &workflowType,
	}

	assert.Equal(s.T(), VALID_SWF_WORKFLOW_DOMAIN, c.GetDomain())
}

func (s *SimpleWorkflowConfigSuite) TestValidateErrorWorkflowType() {
	domain := VALID_SWF_WORKFLOW_DOMAIN
	name := ""
	version := ""
	workflowType := WorkflowType{Name: &name, Version: &version}

	c := &SimpleWorkflowConfig{
		Domain:       &domain,
		WorkflowType: &workflowType,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Sub Field of WorkflowType, failed to validate with error, String Field: Name, contains an empty string"), err)
}

func (s *SimpleWorkflowConfigSuite) TestGetWorkflowType() {
	domain := VALID_SWF_WORKFLOW_DOMAIN
	name := VALID_SWF_WORKFLOW_TYPE_NAME
	version := VALID_SWF_WORKFLOW_TYPE_VERSION
	workflowType := WorkflowType{Name: &name, Version: &version}

	c := &SimpleWorkflowConfig{
		Domain:       &domain,
		WorkflowType: &workflowType,
	}

	assert.Equal(s.T(), workflowType, c.GetWorkflowType())
}

func (s *SimpleWorkflowConfigSuite) TestWorkflowTypeToAWS() {
	name := VALID_SWF_WORKFLOW_TYPE_NAME
	version := VALID_SWF_WORKFLOW_TYPE_VERSION
	workflowType := WorkflowType{Name: &name, Version: &version}

	c := workflowType.ToAWS()

	assert.Equal(s.T(), *workflowType.Name, *c.Name)
	assert.Equal(s.T(), *workflowType.Version, *c.Version)
}
