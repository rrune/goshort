package util

import (
	"fmt"
	"log"
	"os"
)

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorYellow = "\033[33m"

func Check(err error, log bool) (r bool) {
	if err != nil {
		r = true
		if log {
			fmt.Println(string(colorYellow), err, string(colorReset))
		}
	}
	return
}

func CheckPanic(err error) {
	if err != nil {
		f, err2 := os.OpenFile("./data/err.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err2 != nil {
			log.Fatal(err, err2)
		}
		defer f.Close()
		_, err2 = f.Write([]byte(fmt.Sprint(err)))
		if err2 != nil {
			log.Fatal(err, err2)
		}
		panic(err)
	}
}
