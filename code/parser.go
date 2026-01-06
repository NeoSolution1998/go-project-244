package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func parseFile(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		return parseJSON(data)
	default:
		return nil, fmt.Errorf("Неподдерживаемый формат: %s", ext)
	}
}

func parseJSON(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
