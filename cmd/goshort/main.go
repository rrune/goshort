package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rrune/goshort/internal/database"
	"github.com/rrune/goshort/internal/models"
	"github.com/rrune/goshort/internal/router"
	"github.com/rrune/goshort/internal/short"
	"github.com/rrune/goshort/internal/util"
	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.OpenFile("./data/goshort.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetFlags(2 | 3)
	log.Println("")

	var config models.Config
	ymlDatam, err := os.ReadFile("./config/config.yml")
	util.CheckPanic(err)
	err = yaml.Unmarshal(ymlDatam, &config)
	util.CheckPanic(err)

	db, err := database.New(config.Type, config.Username, config.Password, config.Address)
	util.CheckPanic(err)
	shorter := short.New(config.Url, db)

	log.Println("Running on Port " + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router.NewRouter(shorter, config)))
}
