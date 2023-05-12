// Code generated by jsonenums -type=Month; DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

var (
	_MonthNameToValue = map[string]Month{
		"Shrawan":  Shrawan,
		"Bhadra":   Bhadra,
		"Ashoj":    Ashoj,
		"Karthik":  Karthik,
		"Manghsir": Manghsir,
		"Poush":    Poush,
		"Magh":     Magh,
		"Falgun":   Falgun,
		"Chaitra":  Chaitra,
		"Baisakh":  Baisakh,
		"Jyestha":  Jyestha,
		"Ashad":    Ashad,
	}

	_MonthValueToName = map[Month]string{
		Shrawan:  "Shrawan",
		Bhadra:   "Bhadra",
		Ashoj:    "Ashoj",
		Karthik:  "Karthik",
		Manghsir: "Manghsir",
		Poush:    "Poush",
		Magh:     "Magh",
		Falgun:   "Falgun",
		Chaitra:  "Chaitra",
		Baisakh:  "Baisakh",
		Jyestha:  "Jyestha",
		Ashad:    "Ashad",
	}
)

func init() {
	var v Month
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MonthNameToValue = map[string]Month{
			interface{}(Shrawan).(fmt.Stringer).String():  Shrawan,
			interface{}(Bhadra).(fmt.Stringer).String():   Bhadra,
			interface{}(Ashoj).(fmt.Stringer).String():    Ashoj,
			interface{}(Karthik).(fmt.Stringer).String():  Karthik,
			interface{}(Manghsir).(fmt.Stringer).String(): Manghsir,
			interface{}(Poush).(fmt.Stringer).String():    Poush,
			interface{}(Magh).(fmt.Stringer).String():     Magh,
			interface{}(Falgun).(fmt.Stringer).String():   Falgun,
			interface{}(Chaitra).(fmt.Stringer).String():  Chaitra,
			interface{}(Baisakh).(fmt.Stringer).String():  Baisakh,
			interface{}(Jyestha).(fmt.Stringer).String():  Jyestha,
			interface{}(Ashad).(fmt.Stringer).String():    Ashad,
		}
	}
}

// MarshalJSON is generated so Month satisfies json.Marshaler.
func (r Month) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _MonthValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Month: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Month satisfies json.Unmarshaler.
func (r *Month) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Month should be a string, got %s", data)
	}
	v, ok := _MonthNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Month %q", s)
	}
	*r = v
	return nil
}