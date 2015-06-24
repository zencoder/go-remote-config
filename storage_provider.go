package remoteconfig

import "errors"

type StorageProvider string

const (
	STORAGE_PROVIDER_AWS StorageProvider = "aws"
)

func (s *StorageProvider) UnmarshalText(data []byte) error {
	sString := string(data[:])
	*s = (StorageProvider)(sString)
	return s.Validate()
}

func (s StorageProvider) Validate() error {
	if s != STORAGE_PROVIDER_AWS {
		return errors.New("Invalid storage provider")
	}
	return nil
}
