package services

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"strings"
)

func EncodeData(data any, result *bytes.Buffer) error {
	enc := gob.NewEncoder(result)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func GetStructFieldByTag(tag string, s interface{}) (string, error) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		return "", fmt.Errorf("bad type")
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get("json"), ",")[0] // use split to ignore tag "options"
		if v == tag {
			return f.Name, nil
		}
	}
	return "", nil
}
