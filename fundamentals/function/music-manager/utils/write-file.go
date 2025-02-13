package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteDataToJson[T any](filename string, data T) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.WriteFile(filename, jsonData, 0644)

	fmt.Println("successfull created data")
}
