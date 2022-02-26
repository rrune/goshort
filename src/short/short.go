package short

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/rrune/goshort/database"
	"github.com/rrune/goshort/util"
)

type Short struct {
	BaseURL string
	DB      database.Database
}

func New(url string, db database.Database) Short {
	return Short{
		BaseURL: url,
		DB:      db,
	}
}

func (s Short) AddShort(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if util.Check(err, true) {
		w.Write([]byte("Error"))
		return
	}
	defer r.Body.Close()

	url := string(b)
	if !isValidUrl(url) {
		w.Write([]byte("Body needs to be a link"))
		return
	}

	msg, err := s.DB.AddShort(url, r.RemoteAddr)
	if util.Check(err, true) {
		log.Println(msg)
		w.Write([]byte("Error"))
	}

	w.Write([]byte(s.BaseURL + msg))
}

func (s Short) Redirect(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "short")
	shorts, err := s.DB.GetShorts(short)
	if util.Check(err, true) {
		w.Write([]byte("Error"))
		return
	}
	if len(shorts) == 0 {
		w.Write([]byte("Short does not exist"))
		return
	}
	url := shorts[0].Url
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (s Short) DelShort(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if util.Check(err, true) {
		w.Write([]byte("Error"))
		return
	}
	defer r.Body.Close()

	short := string(b)
	exists, msg, err := s.DB.DelShort(short)
	if util.Check(err, true) {
		log.Println(msg)
		w.Write([]byte("Error"))
		return
	}
	if !exists {
		w.Write([]byte("Short does not exist"))
		return
	}
	w.Write([]byte("Deleted " + short))
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
