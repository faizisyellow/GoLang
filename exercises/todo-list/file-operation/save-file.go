package fo

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveToFile[T any](filename string, data []T) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error while add new task: %v", err.Error())
		return nil
	}

	return os.WriteFile(filename, jsonData, 0644)
}
