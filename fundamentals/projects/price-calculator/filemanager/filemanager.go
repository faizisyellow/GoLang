package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, err
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputPath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}
