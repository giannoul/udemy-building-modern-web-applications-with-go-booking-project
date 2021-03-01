package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/giannoul/udemy-building-modern-web-applications-with-go-booking-project/pkg/config"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go-booking-project/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// http.Handler is also known as multiplexer
	/*mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))*/

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//mux.Use(middleware.Logger) <-- this avoids the `panic: interface conversion: *scs.bufferedResponseWriter is not io.ReaderFrom: missing method ReadFrom`

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//Tell the server where to search for static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
