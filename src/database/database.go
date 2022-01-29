package database

import (
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rrune/goshort/models"
	"github.com/rrune/goshort/util"
)

type Database struct {
	DB *sqlx.DB
}

func New(username string, password string, address string) (d Database, err error) {
	source := fmt.Sprintf("%s:%s@%s", username, password, address)
	db, err := sqlx.Open("mysql", (source))
	if util.Check(err, false) {
		return
	}
	d = Database{
		DB: db,
	}
	return
}

func (d Database) GetShorts(short string) (r []models.Short, err error) {
	if short == "" {
		err = d.DB.Select(&r, "SELECT * FROM shortLinks")
		return
	}
	err = d.DB.Select(&r, "SELECT * FROM shortLinks WHERE short LIKE ?", short)
	return
}

func (d Database) AddShort(url string) (msg string, err error) {
	var random string
	alreadyExists := true
	for alreadyExists {
		random = d.Random(3)
		shorts, err2 := d.GetShorts(random)
		if util.Check(err2, true) {
			msg = "Error while generating short"
			err = err2
			return
		}
		if len(shorts) == 0 {
			alreadyExists = false
		}
	}
	insert, err := d.DB.Query("INSERT INTO shortLinks VALUES(?, ?, ?)", random, url, time.Now())
	if util.Check(err, true) {
		msg = "Error with database"
		return
	}
	defer insert.Close()

	msg = random
	return
}

func (d Database) DelShort(short string) (exists bool, msg string, err error) {
	shorts, err := d.GetShorts(short)
	if util.Check(err, true) {
		msg = "Error with database"
		return
	}
	if len(shorts) == 0 {
		msg = "Short does not exist"
		return
	}
	exists = true
	insert, err := d.DB.Query("DELETE FROM shortLinks WHERE short LIKE ?", short)
	if util.Check(err, true) {
		msg = "Error while deleting from database"
		return
	}
	defer insert.Close()
	msg = "Success"
	return
}

func (d Database) Random(n int) (s string) {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
