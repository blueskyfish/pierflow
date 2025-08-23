package utils

import "encoding/json"

func Stringify[T any](obj T) (string, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
