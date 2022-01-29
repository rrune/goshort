package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/rrune/goshort/models"
	"github.com/rrune/goshort/short"
)

var config models.Config

func NewRouter(short short.Short, cfg models.Config) http.Handler {
	config = cfg

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	fs := http.FileServer(http.Dir("./public"))
	r.Handle("/*", http.StripPrefix("/", fs))

	r.Get("/{short}", short.Redirect)

	r.Group(func(r chi.Router) {
		if config.Auth {
			r.Use(auth)
		}
		r.Post("/", short.AddShort)
	})

	r.Group(func(r chi.Router) {
		r.Use(auth)
		r.Delete("/", short.DelShort)
	})
	return r
}
