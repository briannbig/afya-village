package util

import (
	"encoding/json"
)

func GetBytes(payload any) ([]byte, error) {

	str, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	buf := []byte(str)

	return buf, nil
}
