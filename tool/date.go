package tool

import "time"

func DateStr2Timestamp(date string) int64 {
	if date == "" {
		return 0
	}
	t, _ := time.Parse("2006-01-02", date)
	return t.Unix()
}
