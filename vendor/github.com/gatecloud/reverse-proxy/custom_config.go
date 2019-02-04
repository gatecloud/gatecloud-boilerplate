package proxy

import (
	"errors"
	"fmt"
)

// CustomConfig representing the string-interface key-value pairs
type CustomConfig map[string]interface{}

// Get gets the value of the key-value pair by key
func (cc *CustomConfig) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("key can not be empty")
	}

	v, ok := (*cc)[key]
	if !ok {
		return nil, fmt.Errorf("key %s is not valid", key)
	}
	return v, nil
}
