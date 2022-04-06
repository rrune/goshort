package database

import (
	"errors"
	"math/rand"
	"time"

	"github.com/rrune/goshort/models"
)

type Database interface {
	GetShorts(string) ([]models.Short, error)
	AddShort(string, string) (string, error)
	DelShort(string) (bool, string, error)
}

func New(dbType string, username string, password string, address string) (d Database, err error) {
	switch dbType {
	case "mysql":
		d, err = newMySQL(username, password, address)
	default:
		err = errors.New("Unknown database type")
	}
	return
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
