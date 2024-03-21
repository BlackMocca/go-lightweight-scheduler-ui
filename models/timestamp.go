package models

import (
	"errors"
	"strings"
	"time"
)

type Timestamp time.Time

/*
------------------------
Timestamp Function
------------------------
*/

func (j *Timestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		return nil
	}
	var layouts = []string{
		string(time.RFC3339),
		"2006-01-02T15:04:05+07:00",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
	}

	var isMatch bool
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			isMatch = true
			*j = Timestamp(t)
			break
		}
	}
	if !isMatch {
		return errors.New("can not parse timestamp please define layout in models.Timestamp")
	}
	return nil
}

func (j Timestamp) ToTime() time.Time {
	return time.Time(j)
}
