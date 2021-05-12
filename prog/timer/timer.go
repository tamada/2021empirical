package main

import "time"

type timer struct {
	number int64
	unit   string
}

func (t *timer) Duration() time.Duration {
	switch t.unit {
	case "min":
		return time.Duration(t.number) * time.Minute
	case "nsec":
		return time.Duration(t.number) * time.Nanosecond
	case "usec": // micro seconds
		return time.Duration(t.number) * time.Nanosecond * time.Duration(1000)
	case "msec":
		return time.Duration(t.number) * time.Nanosecond * time.Duration(1000000)
	case "hour":
		return time.Duration(t.number) * time.Hour
	case "sec":
		return time.Duration(t.number) * time.Second
	default:
		return time.Second
	}
}
