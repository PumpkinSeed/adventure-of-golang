package main

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/volatiletech/sqlboiler/randomize"
)

var NullBytes = []byte("null")

// Time is a nullable time.Time. It supports SQL and JSON serialization.
type ExpiratonDate struct {
	Time   time.Time
	String string
	Valid  bool
}

// NewTime creates a new Time.
func NewExpiratonDate(t time.Time, valid bool) ExpiratonDate {
	return ExpiratonDate{
		Time:   t,
		String: ParseTimeToExp(t),
		Valid:  valid,
	}
}

// TimeFrom creates a new Time that will always be valid.
func ExpiratonDateFrom(t time.Time) ExpiratonDate {
	return NewExpiratonDate(t, true)
}

// TimeFromPtr creates a new Time that will be null if t is nil.
func ExpiratonDateFromPtr(t *time.Time) ExpiratonDate {
	if t == nil {
		return NewExpiratonDate(time.Time{}, false)
	}
	return NewExpiratonDate(*t, true)
}

// MarshalJSON implements json.Marshaler.
func (t ExpiratonDate) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return NullBytes, nil
	}

	if t.String != "" {
		return []byte(`"` + t.String + `"`), nil
	}
	return []byte(`"` + ParseTimeToExp(t.Time) + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *ExpiratonDate) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullBytes) {
		t.Valid = false
		t.Time = time.Time{}
		return nil
	}

	str := strings.TrimPrefix(string(data), `"`)
	str = strings.TrimSuffix(str, `"`)
	if len(str) == 5 {
		var err error
		t.Time, err = ParseExpToTime(str)
		t.String = str
		if err != nil {
			return err
		}
	} else {
		t.Valid = false
		t.Time = time.Time{}
		return errors.New("Invalid expiration date")
	}

	t.Valid = true
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (t ExpiratonDate) MarshalText() ([]byte, error) {
	if !t.Valid {
		return NullBytes, nil
	}
	if t.String != "" {
		return []byte(t.String), nil
	}
	return []byte(ParseTimeToExp(t.Time)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *ExpiratonDate) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		t.Valid = false
		return nil
	}
	t.String = string(text)
	var err error
	t.Time, err = ParseExpToTime(string(text))
	if err != nil {
		t.Valid = false
		return nil
	}
	t.Valid = true
	return nil
}

// SetValid changes this Time's value and sets it to be non-null.
func (t *ExpiratonDate) SetValid(v time.Time) {
	t.Time = v
	t.String = ParseTimeToExp(v)
	t.Valid = true
}

func (t *ExpiratonDate) SetValidFromStr(v string) {
	parsed, _ := ParseExpToTime(v)
	t.Time = parsed
	t.String = v
	t.Valid = true
}

// Scan implements the Scanner interface.
func (t *ExpiratonDate) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		t.Time = x
		t.String = ParseTimeToExp(x)
	case string:
		t.String = x
		t.Time, err = ParseExpToTime(x)
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", value, value)
	}
	t.Valid = err == nil
	return err
}

// Value implements the driver Valuer interface.
func (t ExpiratonDate) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// Randomize for sqlboiler
func (t *ExpiratonDate) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	if shouldBeNull {
		t.Time = time.Time{}
		t.Valid = false
	} else {
		t.Time = randomize.Date(nextInt)
		t.String = ParseTimeToExp(t.Time)
		t.Valid = true
	}
}

type Testit struct {
	ExpiratonDate ExpiratonDate `json:"exp_date"`
}

var testit = `{
	"exp_date": "12/20"
}`

func main() {
	var t Testit
	err := json.Unmarshal([]byte(testit), &t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)

	result, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}

func ParseExpToTime(exp string) (time.Time, error) {
	var t time.Time
	length := len(exp)
	if length != 4 && length != 5 {
		return t, errors.New("Invalid expiration date length")
	}
	var month int
	var year int
	if length == 4 {
		month, _ = strconv.Atoi(exp[:2])
		year, _ = strconv.Atoi(exp[2:4])
	}
	if length == 5 {
		month, _ = strconv.Atoi(exp[:2])
		year, _ = strconv.Atoi(exp[3:5])
	}

	if month < 1 || month > 12 {
		return t, errors.New("Invalid month in expiration date")
	}

	year = year + 2000
	if year <= 2000 || year > 2050 {
		return t, errors.New("Invalid year in expiration date")
	}

	t = time.Date(year, time.Month(month), (time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()), 0, 0, 0, 0, time.UTC)
	return t, nil
}

func ParseTimeToExp(t time.Time) string {
	return t.Format("01/06")
}
