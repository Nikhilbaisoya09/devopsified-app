
package app

import (
    "encoding/json"
    "fmt"
    "html"
    "log/slog"
    "net/http"
    "time"
)

type Config struct {
    Name        string
    Version     string
    Environment string
}

// Create the HTTP routes.
func NewHandler(cfg Config, logger *slog.Logger) http.Handler {
    mux := http.NewServeMux()

    // Home page
    mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }

        w.Header().Set("Content-Type", "text/html; charset=utf-8")

        fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>DevOpsified</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background: #edf2f7;
            text-align: center;
            padding: 80px 20px;
        }
        .card {
            background: white;
            padding: 40px;
            max-width: 550px;
            margin: auto;
            border-radius: 14px;
            box-shadow: 0 8px 25px #0001;
        }
        h1 { color: #2563eb; }
        a { color: #2563eb; }
    </style>
</head>
<body>
    <div class="card">
        <h1>DevOpsified</h1>
        <h2>Golang DevOps Project</h2>
        <p>Version: %s</p>
        <p>Environment: %s</p>
        <p>Application is running successfully!</p>
        <a href="/health">Health</a> |
        <a href="/ready">Readiness</a> |
        <a href="/api/v1/info">API Info</a>
    </div>
</body>
</html>`,
            html.EscapeString(cfg.Version),
            html.EscapeString(cfg.Environment),
        )
    })

    // Liveness endpoint
    mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, map[string]string{
            "status": "ok",
        })
    })

    // Readiness endpoint
    mux.HandleFunc("GET /ready", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, map[string]string{
            "status": "ready",
        })
    })

    // REST API endpoint
    mux.HandleFunc("GET /api/v1/info", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, map[string]string{
            "name":        cfg.Name,
            "version":     cfg.Version,
            "environment": cfg.Environment,
        })
    })

    // Log every HTTP request.
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        logger.Info("http request",
            "method", r.Method,
            "path", r.URL.Path,
        )
        mux.ServeHTTP(w, r)
        logger.Info("request completed",
            "duration_ms", time.Since(start).Milliseconds(),
        )
    })
}

func writeJSON(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(data)
}
