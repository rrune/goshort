package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rrune/goshort/models"
	"github.com/rrune/goshort/util"
)

type Sql struct {
	DB *sqlx.DB
}

func newServerSQL(dbtype string, username string, password string, address string) (d Database, err error) {
	source := fmt.Sprintf("%s:%s@%s", username, password, address)
	db, err := sqlx.Open(dbtype, source)
	if util.Check(err, false) {
		return
	}
	d = Sql{
		DB: db,
	}
	return
}

func newSQLite(dbtype string, filename string) (d Database, err error) {
	db, err := sqlx.Open(dbtype, "../../data/"+filename)
	if util.Check(err, false) {
		return
	}
	d = Sql{
		DB: db,
	}
	return
}

// check if the table exists, and create it if it does not
func (d Sql) InitDB() (err error) {
	_, err = d.DB.Query(`
	CREATE TABLE if not exists shortLinks (
		'short' text NOT NULL,
		'url' text NOT NULL,
		'timestamp' timestamp NOT NULL,
		'ip' text NOT NULL
	)
	`)
	return
}

func (d Sql) GetShorts(short string) (r []models.Short, err error) {
	if short == "" {
		err = d.DB.Select(&r, "SELECT * FROM shortLinks")
		return
	}
	err = d.DB.Select(&r, "SELECT * FROM shortLinks WHERE short LIKE ?", short)
	return
}

func (d Sql) AddShort(url string, ip string) (msg string, err error) {
	var random string
	alreadyExists := true
	for alreadyExists {
		random = Random(3)
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
	insert, err := d.DB.Query("INSERT INTO shortLinks VALUES(?, ?, ?, ?)", random, url, time.Now(), ip)
	if util.Check(err, true) {
		msg = "Error with database"
		return
	}
	defer insert.Close()

	msg = random
	return
}

func (d Sql) DelShort(short string) (exists bool, msg string, err error) {
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
