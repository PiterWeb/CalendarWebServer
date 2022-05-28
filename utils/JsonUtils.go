package utils

import "encoding/json"

func StructToJson(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil

}

func JsonToStruct(data string) (map[string]string, error) {

	var result map[string]string

	err := json.Unmarshal([]byte(data), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}