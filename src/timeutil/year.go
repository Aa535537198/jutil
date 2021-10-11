package jtime

import "time"

func AddYear(t *time.Time, year int) {
	*t = t.AddDate(year, 0, 0)
}
