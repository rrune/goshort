package database

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/rrune/goshort/models"
	"github.com/rrune/goshort/util"
	"gopkg.in/yaml.v2"
)

var DB Database
var short string

func TestMain(m *testing.M) {
	var config models.Config
	ymlData, err := os.ReadFile("../config.yml")
	util.CheckPanic(err)
	err = yaml.Unmarshal((ymlData), &config)
	util.CheckPanic(err)

	DB, err = New("mysql", config.Username, config.Password, config.Address)
	util.CheckPanic(err)

	m.Run()
}

func TestAddShort(t *testing.T) {
	is := is.New(t)
	url := "https://example.com"
	msg, err := DB.AddShort(url)
	short = msg

	is.NoErr(err) // TestAddDelShort: Error (1)
}

func TestDelShort(t *testing.T) {
	is := is.New(t)

	s, _, err := DB.DelShort(short)

	is.True(s)    // TestAddDelShort: Not true (1)
	is.NoErr(err) // TestAddDelShort: Error (1)
}
