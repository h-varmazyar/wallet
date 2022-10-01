package walletApi

import (
	"database/sql/driver"
	"encoding/json"
)

func (TransactionStatus) InRange(v interface{}) bool {
	_, ok := TransactionStatus_value[v.(TransactionStatus).String()]
	return ok
}
func (x *TransactionStatus) Scan(value interface{}) error {
	*x = TransactionStatus(TransactionStatus_value[value.(string)])
	return nil
}
func (x *TransactionStatus) Value() (driver.Value, error) {
	return x.String(), nil
}
func (x *TransactionStatus) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*x = TransactionStatus(TransactionStatus_value[str])
	return nil
}
func (x TransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

func (TransactionType) InRange(v interface{}) bool {
	_, ok := TransactionType_value[v.(TransactionType).String()]
	return ok
}
func (x *TransactionType) Scan(value interface{}) error {
	*x = TransactionType(TransactionType_value[value.(string)])
	return nil
}
func (x *TransactionType) Value() (driver.Value, error) {
	return x.String(), nil
}
func (x *TransactionType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*x = TransactionType(TransactionType_value[str])
	return nil
}
func (x TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
