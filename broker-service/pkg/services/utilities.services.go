package services

import (
	"bytes"
	"encoding/gob"
)

func DecodeData(data []byte, result any) error {
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	if err := dec.Decode(&result); err != nil {
		return err
	}
	return nil
}
