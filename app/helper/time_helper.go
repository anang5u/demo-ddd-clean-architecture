package helper

import "time"

func GetTimeNow() time.Time {
	return time.Now()
}

func GetTimeNowPtr() *time.Time {
	now := GetTimeNow()
	return &now
}
