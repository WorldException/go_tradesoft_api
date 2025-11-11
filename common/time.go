package common

import (
	"time"
)

// CustomTime is a custom type for handling time format "04.03.2014"
type CustomTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler interface for CustomTime
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := string(data[1 : len(data)-1])

	// Parse the date string in "04.03.2014" format
	parsedTime, err := time.Parse("02.01.2006", str)
	if err != nil {
		return err
	}

	// Assign the parsed time to the embedded time.Time field
	ct.Time = parsedTime
	return nil
}
