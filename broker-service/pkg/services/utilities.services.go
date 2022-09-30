package services

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/vimalkuriensam/broker-service/pkg/config"
)

func ReadRequest(req *http.Request, fields []string) (config.ReadValue, error) {
	cfg := config.GetConfig()
	go cfg.ReadJSON(req)
	data := (<-cfg.DataChan).(config.ReadValue)
	err := AcceptableFields(data.D.(map[string]interface{}), fields)
	if err != nil {
		return config.ReadValue{}, err
	}
	return data, nil
}

func DecodeData(data []byte, result any) error {
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	if err := dec.Decode(&result); err != nil {
		return err
	}
	return nil
}

func AcceptableFields(incoming map[string]interface{}, acceptability []string) error {
	for key := range incoming {
		if isAcceptible := isAcceptable(key, acceptability); !isAcceptible {
			return fmt.Errorf("%s is an not a valid field", key)
		}
	}
	return nil
}

func isAcceptable(key string, acceptability []string) bool {
	for _, value := range acceptability {
		if value == key {
			return true
		}
	}
	return false
}
