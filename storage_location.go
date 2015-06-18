package remoteconfig

type StorageLocation string

func (s *StorageLocation) UnmarshalText(data []byte) error {
	sString := string(data[:])
	*s = (StorageLocation)(sString)
	return nil
}
