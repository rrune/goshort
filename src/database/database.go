package database

import (
	"math/rand"
	"time"

	"github.com/rrune/goshort/models"
)

type Database interface {
	GetShorts(string) ([]models.Short, error)
	AddShort(string) (string, error)
	DelShort(string) (bool, string, error)
}

func Random(n int) (s string) {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
