package helper

import "github.com/google/uuid"

func UuidNew() uuid.UUID {
	return uuid.New()
}

func UuidMustParse(s string) uuid.UUID {
	return uuid.MustParse(s)
}
