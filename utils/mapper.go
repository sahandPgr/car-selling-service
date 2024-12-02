package utils

import "encoding/json"

func TypeConverter[T any](data any) (*T, error) {
	result := new(T)
	jsonRes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonRes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
