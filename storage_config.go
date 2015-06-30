package remoteconfig

type StorageConfig struct {
	Provider *StorageProvider `json:"provider,omitempty"`
	Location *StorageLocation `json:"location,omitempty"`
}

func (s StorageConfig) Validate() error {
	if err := s.Provider.Validate(); err != nil {
		return err
	}
	if *s.Provider == STORAGE_PROVIDER_AWS {
		if err := (*AWSRegion)(s.Location).Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (s *StorageConfig) GetProvider() StorageProvider {
	return *s.Provider
}

func (s *StorageConfig) GetLocation() StorageLocation {
	return *s.Location
}
