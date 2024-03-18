package remoteconfig

import (
	"encoding/json"
	"fmt"
	"io"
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
func LoadConfigFromURL(configURL string, configStruct interface{}) error {
	resp, err := http.Get(configURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request to '%s' returned non-200 OK status '%d: %s'", configURL, resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return ReadJSONValidate(resp.Body, configStruct)
}

// Downloads JSON from a URL, decodes it and then validates.
func ReadJSONValidate(cfgReader io.Reader, configStruct interface{}) error {
	// Do a streaming JSON decode
	dec := json.NewDecoder(cfgReader)
	if err := dec.Decode(configStruct); err != nil {
		return fmt.Errorf("Failed to decode JSON, with error, %s", err.Error())
	}

	// Run validation on the config
	if err := validateConfigWithReflection(configStruct); err != nil {
		return err
	}

	return nil
}

func isNilFixed(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		//use of IsNil method
		return v.IsNil()
	}
	return false
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

		if valueField.Kind() == reflect.Struct && typeField.Anonymous {
			continue
		}

		if isNilFixed(valueField) && !optional {
			return fmt.Errorf("Field: %s, not set", typeField.Name)
		} else if isNilFixed(valueField) && optional {
			continue
		}

		// Handle a slice type
		if valueField.Kind() == reflect.Slice {
			if valueField.Len() <= 0 {
				return fmt.Errorf("Slice Field: %s, is empty", typeField.Name)
			}
			for i := 0; i < valueField.Len(); i++ {
				sliceValue := valueField.Index(i)
				if sliceValue.Kind() != reflect.Ptr || sliceValue.IsNil() || sliceValue.Elem().Kind() != reflect.Struct {
					continue
				}
				if err := validateConfigWithReflection(sliceValue.Interface()); err != nil {
					return err
				}
			}
			continue
		}

		// Handle a map type
		if valueField.Kind() == reflect.Map {
			for _, key := range valueField.MapKeys() {
				mapValue := valueField.MapIndex(key)
				if mapValue.Kind() != reflect.Ptr || mapValue.IsNil() || mapValue.Elem().Kind() != reflect.Struct {
					continue
				}
				if err := validateConfigWithReflection(mapValue.Interface()); err != nil {
					return fmt.Errorf("Sub field of %s with key '%s' failed to validated with error, %s", typeField.Name, key, err)
				}
			}
			continue
		}

		// Skip functions
		if valueField.Kind() == reflect.Func {
			continue
		}

		// If this is a string pointer field, check that it isn't empty (unless optional)
		if s, ok := valueField.Interface().(*string); ok {
			if *s == "" {
				return fmt.Errorf("String Field: %s, contains an empty string", typeField.Name)
			}
			continue
		}

		// If this is a string field, check that it isn't empty (unless optional)
		if s, ok := valueField.Interface().(string); ok {
			if s == "" {
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
