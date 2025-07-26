package utils

import (
	"encoding/json"
	"fmt"
)

func JsonMarshal(data any) (string, error) {

	byteData, err := json.Marshal(data)

	if err != nil {
		return "", fmt.Errorf("failed to marshal data : %w", err)
	}

	return string(byteData), nil
}

func JsonUnmarshal(data string, destination any) error {

	if err := json.Unmarshal([]byte(data), &destination); err != nil {
		return fmt.Errorf("failed to unmarshal data : %w", err)
	}

	return nil
}
