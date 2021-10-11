package jtime

import "time"

func AddMonth(t *time.Time, month int) {
	*t = t.AddDate(0, month, 0)
}
