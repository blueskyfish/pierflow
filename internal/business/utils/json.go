package utils

import "encoding/json"

func Stringify[T any](obj T) (string, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func CovertTo[T any](data interface{}, value *T) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonBytes, value); err != nil {
		return err
	}
	return nil
}
