package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rrune/goshort/database"
	"github.com/rrune/goshort/models"
	"github.com/rrune/goshort/router"
	"github.com/rrune/goshort/short"
	"github.com/rrune/goshort/util"
	"gopkg.in/yaml.v2"
)

func main() {
	var config models.Config
	ymlDatam, err := os.ReadFile("../config.yml")
	util.CheckPanic(err)
	err = yaml.Unmarshal(ymlDatam, &config)
	util.CheckPanic(err)

	db, err := database.New(config.Username, config.Password, config.Address)
	util.CheckPanic(err)
	shorter := short.New(config.Url, db)

	go util.WaitForExit()

	fmt.Println("Running on Port " + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router.NewRouter(shorter, config)))
}
