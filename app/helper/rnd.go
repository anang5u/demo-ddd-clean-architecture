package helper

import (
	"math/rand"
	"time"
)

func RndAphanumeric(length int) string {
	rand.Seed(time.Now().In(time.Local).UnixNano())
	var letterRunes = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func RndAlphabet(length int) string {
	rand.Seed(time.Now().In(time.Local).UnixNano())
	var letterRunes = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func RndNumeric(length int) string {
	rand.Seed(time.Now().In(time.Local).UnixNano())
	var letterRunes = []rune("0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
