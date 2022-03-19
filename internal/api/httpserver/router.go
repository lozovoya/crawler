package httpserver

import (
	v1 "crawler/internal/api/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter (mux *chi.Mux, crawlerController *v1.Crawler) chi.Mux {
	mux.Use(middleware.Logger)
	mux.Route("/api/v1", func(router chi.Router){
		router.Post("/gettitles", crawlerController.GetTitles)
	})
	return *mux
}
