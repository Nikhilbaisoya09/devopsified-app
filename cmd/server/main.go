
package main

import (
    "context"
    "errors"
    "log/slog"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "example.com/devopsified-app/internal/app"
)

// Reads an environment variable or uses a default value.
func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}

func main() {
    // Create a JSON logger.
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    port := getEnv("PORT", "8080")

    config := app.Config{
        Name:        "DevOpsified",
        Version:     getEnv("APP_VERSION", "1.0.0"),
        Environment: getEnv("APP_ENV", "development"),
    }

    // Create the HTTP server.
    server := &http.Server{
        Addr:              ":" + port,
        Handler:           app.NewHandler(config, logger),
        ReadHeaderTimeout: 5 * time.Second,
    }

    listener, err := net.Listen("tcp", server.Addr)
    if err != nil {
        logger.Error("cannot start server", "error", err)
        os.Exit(1)
    }

    errorsChan := make(chan error, 1)

    go func() {
        if err := server.Serve(listener); err != nil &&
            !errors.Is(err, http.ErrServerClosed) {
            errorsChan <- err
        }
    }()

    logger.Info("server started", "port", port,
        "version", config.Version,
        "environment", config.Environment)

    // Listen for shutdown signals.
    ctx, stop := signal.NotifyContext(
        context.Background(),
        os.Interrupt,
        syscall.SIGTERM,
    )
    defer stop()

    select {
    case <-ctx.Done():
        logger.Info("shutdown requested")
    case err := <-errorsChan:
        logger.Error("server failed", "error", err)
        os.Exit(1)
    }

    // Gracefully close active requests.
    shutdownCtx, cancel := context.WithTimeout(
        context.Background(),
        10*time.Second,
    )
    defer cancel()

    if err := server.Shutdown(shutdownCtx); err != nil {
        logger.Error("shutdown failed", "error", err)
        os.Exit(1)
    }

    logger.Info("server stopped")
}
