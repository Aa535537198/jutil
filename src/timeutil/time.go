package jtime

import (
	"time"
)

var h float64 = 3600000000000
var m float64 = 60000000000
var s float64 = 1000000000

//year, month, day, hour, minute, second, miss
func AddTime(t *time.Time, hour, minute, second int) {
	duration := time.Duration(h*float64(hour) + m*float64(minute) + s*float64(second))
	*t = (*t).Add(duration)
}
