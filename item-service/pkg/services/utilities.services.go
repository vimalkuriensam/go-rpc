package services

import (
	"bytes"
	"encoding/gob"
)

func EncodeData(data any, result *bytes.Buffer) error {
	enc := gob.NewEncoder(result)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}
