package database

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v3"
)

type jsonTime struct {
	time.Time
}

func (t jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// Scan implements the Scanner interface.
func (t *jsonTime) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		t.Time = x
	case nil:
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", value, value)
	}
	return err
}

// Value implements the driver Valuer interface.
func (t jsonTime) Value() (driver.Value, error) {
	return t.Time, nil
}

type jsonNullTime struct {
	null.Time
}

func (t jsonNullTime) MarshalJSON() ([]byte, error) {
	if !t.Time.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Time.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
