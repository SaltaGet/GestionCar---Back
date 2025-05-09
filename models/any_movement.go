package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
)

type AnyMovement struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewAnyMovement(data interface{}) AnyMovement {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	typeName := t.Name()

	return AnyMovement{
		Type: typeName,
		Data: data,
	}
}

func (a *AnyMovement) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AnyMovement) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return json.Unmarshal(bytes, a)
}