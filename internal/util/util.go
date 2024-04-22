package util

import "fmt"

func GetBytes(payload any) ([]byte, error) {
	buf, ok := payload.([]byte)
	if !ok {
		return nil, fmt.Errorf("error transforming payload to bytes")
	}

	return buf, nil
}
