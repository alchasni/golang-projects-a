package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSONSQL json.RawMessage

func (j *JSONSQL) Scan(value interface{}) error {
	valueBytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("unexpected JSONSQL type: %T", value))
	}

	result := json.RawMessage{}

	err := json.Unmarshal(valueBytes, &result)
	if err != nil {
		return err
	}

	*j = JSONSQL(result)
	return nil
}

func (j JSONSQL) Value() (driver.Value, error) {
	if len(j) == 0 {
		return make([]byte, 0), nil
	}

	return json.RawMessage(j).MarshalJSON()
}
