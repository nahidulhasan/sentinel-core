package httpserver

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/jmoiron/sqlx"
    "github.com/nahidulhasan/sentinel-core/internal/infrastructure/config"
    "github.com/nahidulhasan/sentinel-core/internal/infrastructure/httpserver/handlers"
)


func New(cfg *config.Config, db *sqlx.DB) http.Handler {
    r := chi.NewRouter()

    r.Get("/health", handlers.HealthHandler)
    api := chi.NewRouter()
    api.Mount("/v1", handlers.NewUserRouter(db))
    r.Mount("/", api)

    return r
}
