
# DevOpsified

End-to-End DevOps Deployment of a Golang Application.

## Tech Stack

- Golang
- Git and GitHub
- Docker
- Kubernetes
- GitHub Actions
- Helm
- Argo CD
- Ingress and DNS

## Features

- Golang HTTP server
- Homepage
- REST API
- Health and readiness endpoints
- JSON logging
- Graceful shutdown
- Unit testing

## Run Application

Run:

    go mod tidy
    go run ./cmd/server

Open: http://localhost:8080

## API Endpoints

| Endpoint | Purpose |
|---|---|
| / | Homepage |
| /health | Health check |
| /ready | Readiness check |
| /api/v1/info | Application information |

## Run Unit Tests

    go test ./... -v

## Project Progress

- [x] Phase 1: Environment Setup
- [x] Phase 2: Golang Application
- [ ] Phase 3: Docker
- [ ] Phase 4: Kubernetes
- [ ] Phase 5: Helm
- [ ] Phase 6: GitHub Actions
- [ ] Phase 7: Argo CD
- [ ] Phase 8: Ingress and DNS
