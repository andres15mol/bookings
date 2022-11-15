package main

import (
	"net/http"

	"github.com/andres15mol/booking/pkg/config"
	"github.com/andres15mol/booking/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler{ 
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/rooms/generals-quaters", handlers.Repo.General)
	mux.Get("/rooms/majors-suite", handlers.Repo.Major)
	mux.Get("/contact",handlers.Repo.Contact)
	mux.Get("/make-reservation",handlers.Repo.MakeReservation)
	mux.Get("/reservation",handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}