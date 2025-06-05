package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONMap map[string]interface{}

func (j *JSONMap) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return fmt.Errorf("Scan source is not []byte")
    }
    return json.Unmarshal(bytes, j)
}

func StructToJSONMap(s interface{}) (JSONMap, error) {
    b, err := json.Marshal(s)
    if err != nil {
        return nil, err
    }
    var m JSONMap
    err = json.Unmarshal(b, &m)
    return m, err
}