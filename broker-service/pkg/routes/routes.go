package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	// Specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	logger := httplog.NewLogger("gochi.log", httplog.Options{
		JSON: true,
	})
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	mux.Use(httplog.RequestLogger(logger))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(2500 * time.Millisecond))
	mux.Use(middleware.Throttle(1))
	mux.Use(httprate.LimitByIP(100, 1*time.Minute))
	return mux
}
