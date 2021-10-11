package jtime

import "time"

func AddDay(t *time.Time, day int) {
	*t = t.AddDate(0, 0, day)
}
