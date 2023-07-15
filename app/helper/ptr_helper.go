package helper

import (
	"time"

	"github.com/google/uuid"
)

func StrPtr(s string) *string {
	return &s
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func UuidPtr(s uuid.UUID) *uuid.UUID {
	return &s
}

func TimePtr(p time.Time) *time.Time {
	return &p
}
