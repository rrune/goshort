package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/rrune/goshort/internal/models"
	"github.com/rrune/goshort/internal/short"
)

var config models.Config

func NewRouter(short short.Short, cfg models.Config) http.Handler {
	config = cfg

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{cfg.Url},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"Authorization"},
	}))

	fs := http.FileServer(http.Dir("./web/public"))
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
		r.Get("/shorts", short.GetEveryShort)
		r.Delete("/", short.DelShort)
	})
	return r
}
