package schema

import "time"

type TimeRequest struct {
	Date time.Time
	From time.Time
	To   time.Time
}
