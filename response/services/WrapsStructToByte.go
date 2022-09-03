package services

import (
	"encoding/json"
	"fmt"
)

func WrapsStructToByte(_value any) ([]byte, error) {
	if payload, err := json.Marshal(_value); err != nil {
		return nil, fmt.Errorf("%s{%v}", "Error step wraps json to byte.Error msg : ", err)
	} else {
		return payload, nil
	}
}

