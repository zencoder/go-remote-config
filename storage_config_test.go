package remoteconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_STORAGE_CONFIG_PROVIDER StorageProvider = STORAGE_PROVIDER_AWS
	VALID_STORAGE_CONFIG_LOCATION StorageLocation = (StorageLocation)(AWS_REGION_US_WEST_2)
)

type StorageConfigSuite struct {
	suite.Suite
}

func TestStorageConfigSuite(t *testing.T) {
	suite.Run(t, new(StorageConfigSuite))
}

func (s *StorageConfigSuite) SetupSuite() {
}

func (s *StorageConfigSuite) SetupTest() {
}

func (s *StorageConfigSuite) TestValidateConfigWithReflection() {
	p := VALID_STORAGE_CONFIG_PROVIDER
	l := VALID_STORAGE_CONFIG_LOCATION
	c := &StorageConfig{
		Provider: &p,
		Location: &l,
	}

	err := validateConfigWithReflection(c)
	assert.Nil(s.T(), err)
}

func (s *StorageConfigSuite) TestValidateConfigWithReflectionErrorProvider() {
	p := (StorageProvider)("invalid_provider")
	l := VALID_STORAGE_CONFIG_LOCATION
	c := &StorageConfig{
		Provider: &p,
		Location: &l,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: StorageConfig, failed to validate with error, Invalid storage provider"), err)
}

func (s *StorageConfigSuite) TestValidateConfigWithReflectionErrorLocation() {
	p := VALID_STORAGE_CONFIG_PROVIDER
	l := (StorageLocation)("invalid_location")
	c := &StorageConfig{
		Provider: &p,
		Location: &l,
	}

	err := validateConfigWithReflection(c)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.New("Validater Field: StorageConfig, failed to validate with error, Region is invalid"), err)
}

func (s *StorageConfigSuite) TestGetProvider() {
	p := VALID_STORAGE_CONFIG_PROVIDER
	c := &StorageConfig{
		Provider: &p,
	}

	assert.Equal(s.T(), VALID_STORAGE_CONFIG_PROVIDER, c.GetProvider())
}

func (s *StorageConfigSuite) TestGetLocation() {
	l := VALID_STORAGE_CONFIG_LOCATION
	c := &StorageConfig{
		Location: &l,
	}

	assert.Equal(s.T(), VALID_STORAGE_CONFIG_LOCATION, c.GetLocation())
}
