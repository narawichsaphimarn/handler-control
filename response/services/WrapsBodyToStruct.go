package services

import (
	"encoding/json"
	"fmt"
)

func WrapsSuccess(_value []byte, body any) error {
	if err := json.Unmarshal([]byte(_value), body); err != nil {
		return fmt.Errorf("%s{%v}", "Error step wraps success body.Error mgs : ", err)
	}
	return nil
}

func WrapsUnSuccess(_value []byte, body any) error {
	if err := json.Unmarshal([]byte(_value), body); err != nil {
		return fmt.Errorf("%s{%v}", "Error step wraps unsuccess body.Error mgs : ", err)
	}
	return nil
}
