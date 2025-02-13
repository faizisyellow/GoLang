package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile[T any](filename string) (*T, error) {
	var data T

	file, err := os.ReadFile(filename)
	if err != nil {
		// If file doesn't exist, return an initialized empty structure
		if os.IsNotExist(err) {
			return &data, nil
		}
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// If file is empty, return initialized empty structure
	if len(file) == 0 {
		return &data, nil
	}

	// Try to unmarshal if file has content
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &data, nil
}
