package remoteconfig

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	DEFAULT_S3_EXPIRY   uint   = 60
	DEFAULT_S3_ENDPOINT string = ""
)

type Validater interface {
	Validate() error
}

// Downloads a configuration JSON file from S3.
// Parses it to a particular struct type and runs a validation.
// URL should be of the format s3://bucket/path/file.json
func LoadConfigFromS3(configURL string, configRegion AWSRegion, configEndpoint string, configStruct interface{}) error {
	// Build a Signed URL to the config file in S3
	signedURL, err := BuildSignedS3URL(configURL, configRegion, DEFAULT_S3_EXPIRY, configEndpoint)
	if err != nil {
		return err
	}

	return DownloadJSONValidate(signedURL, configStruct)
}

// Downloads JSON from a URL, decodes it and then validates.
func DownloadJSONValidate(signedURL string, configStruct interface{}) error {
	// Download the config file from S3
	resp, err := http.Get(signedURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check that we got a valid response code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Download of JSON failed, URL = %s, Response Code = %d", signedURL, resp.StatusCode)
	}

	// Do a streaming JSON decode
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(configStruct); err != nil {
		return fmt.Errorf("Failed to decode JSON, with error, %s", err.Error())
	}

	// Run validation on the config
	err = validateConfigWithReflection(configStruct)
	if err != nil {
		return err
	}

	return nil
}

// Validates a configuration struct.
// Uses reflection to determine and call the correct Validation methods for each type.
func validateConfigWithReflection(c interface{}) error {
	valueElem := reflect.ValueOf(c).Elem()
	typeElem := reflect.TypeOf(c).Elem()

	// Gets a refection Type value for the Validater interface
	validaterType := reflect.TypeOf((*Validater)(nil)).Elem()

	// If the Validater interface is implemented, call the Validate method
	if typeElem.Implements(validaterType) {
		if err := valueElem.Interface().(Validater).Validate(); err != nil {
			return fmt.Errorf("Validater Field: %s, failed to validate with error, %s", typeElem.Name(), err)
		}
	}

	for i := 0; i < valueElem.NumField(); i++ {
		valueField := valueElem.Field(i)
		typeField := typeElem.Field(i)

		tags := typeField.Tag.Get("remoteconfig")
		optional := strings.Contains(tags, "optional")

		if valueField.IsNil() && !optional {
			return fmt.Errorf("Field: %s, not set", typeField.Name)
		} else if valueField.IsNil() && optional {
			continue
		}

		// Handle a slice type
		if valueField.Kind() == reflect.Slice {
			for i := 0; i < valueField.Len(); i++ {
				if err := validateConfigWithReflection(valueField.Index(i).Interface()); err != nil {
					return err
				}
			}
			continue
		}

		// If this is a string field, check that it isn't empty (unless optional)
		if s, ok := valueField.Interface().(*string); ok {
			if *s == "" {
				return fmt.Errorf("String Field: %s, contains an empty string", typeField.Name)
			}
			continue
		}

		// If the Validater interface is implemented, call the Validate method
		if typeField.Type.Implements(validaterType) {
			if err := valueField.Interface().(Validater).Validate(); err != nil {
				return fmt.Errorf("Validater Field: %s, failed to validate with error, %s", typeField.Name, err)
			}
			continue
		}

		// If this field is a struct type, validate it with reflection
		// We can/should only check the sub-fields of a Struct
		if valueField.Elem().Kind() == reflect.Struct && valueField.Elem().NumField() > 0 {
			if err := validateConfigWithReflection(valueField.Interface()); err != nil {
				return fmt.Errorf("Sub Field of %s, failed to validate with error, %s", typeField.Name, err)
			}
		}
	}

	return nil
}
