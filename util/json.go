package util

import (
	"encoding/json"
)

func FromJson(jsonString string) (any, error) {
	var obj any

	err := json.Unmarshal([]byte(jsonString), obj)

	if err != nil {
		print(err)
		return nil, err
	}

	return obj, nil
}

func ToJson(obj any) (string, error) {
	bytes, err := json.Marshal(obj)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
