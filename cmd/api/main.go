package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/nahidulhasan/sentinel-core/internal/infrastructure/config"
    "github.com/nahidulhasan/sentinel-core/internal/infrastructure/db"
    "github.com/nahidulhasan/sentinel-core/internal/infrastructure/httpserver"
    "github.com/rs/zerolog/log"
)

func main() {
    // load config
    cfg, err := config.Load()
    if err != nil {
        panic(err)
    }

    // logger
    log.Info().Msg("starting sentinel-core")

    // connect to db
    sqlxDB, err := db.NewPostgres(cfg.DB.DSN)
    if err != nil {
        log.Fatal().Err(err).Msg("failed to connect db")
    }
    defer sqlxDB.Close()

    // create http server
    srv := httpserver.New(cfg, sqlxDB)

    // graceful shutdown
    srvAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
    httpServer := &http.Server{
        Addr:    srvAddr,
        Handler: srv,
    }

    go func() {
        log.Info().Msgf("http server listening on %s", srvAddr)
        if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal().Err(err).Msg("httpserver crashed")
        }
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    log.Info().Msg("shutting down")
    if err := httpServer.Shutdown(ctx); err != nil {
        log.Error().Err(err).Msg("shutdown error")
    }
    log.Info().Msg("bye")
}
