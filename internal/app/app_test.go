
package app

import (
    "io"
    "log/slog"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestEndpoints(t *testing.T) {
    logger := slog.New(
        slog.NewTextHandler(io.Discard, nil),
    )

    handler := NewHandler(Config{
        Name:        "DevOpsified",
        Version:     "1.0.0",
        Environment: "test",
    }, logger)

    tests := []struct {
        path     string
        expected string
    }{
        {"/health", `"status":"ok"`},
        {"/ready", `"status":"ready"`},
        {"/api/v1/info", `"version":"1.0.0"`},
    }

    for _, tc := range tests {
        t.Run(tc.path, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodGet, tc.path, nil)
            recorder := httptest.NewRecorder()

            handler.ServeHTTP(recorder, req)

            if recorder.Code != http.StatusOK {
                t.Fatalf("expected 200, got %d", recorder.Code)
            }

            if !strings.Contains(recorder.Body.String(), tc.expected) {
                t.Errorf("unexpected response: %s", recorder.Body.String())
            }
        })
    }
}
