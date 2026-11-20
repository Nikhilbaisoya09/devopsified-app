
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



## Docker Deployment

Build the Docker image:

    docker build -t devopsified:v1 .

Run the container:

    docker run -d --name devopsified-go-phase3 -p 8081:8080 devopsified:v1

Check running containers:

    docker ps

Test the API:

    curl.exe http://localhost:8081/health
    curl.exe http://localhost:8081/ready
    curl.exe http://localhost:8081/api/v1/info

View container logs:

    docker logs devopsified-go-phase3

Stop the container:

    docker stop devopsified-go-phase3

## Completed Phases

- [x] Phase 1: Environment Setup
- [x] Phase 2: Golang Application
- [x] Phase 3: Docker Containerization
- [X] Phase 4: Kubernetes
- [X] Phase 5: Helm
- [X] Phase 6: GitHub Actions
