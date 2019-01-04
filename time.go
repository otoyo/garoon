package garoon

import (
	"time"
)

type Time struct {
	DateTime time.Time `json:"dateTime"`
	TimeZone string    `json:"timeZone"`
}

type DatePeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type TimePeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type DateTimePeriod struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
