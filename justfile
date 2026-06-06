# ezPanel build tasks

# Build frontend assets
frontend-build:
    cd web && npm run build

# Build Go binary (development mode, no embedded frontend)
build:
    go build -tags dev -o ezpanel ./cmd/panel/

# Full build: frontend + backend
all: frontend-build build-prod

# Production build (with embedded frontend)
build-prod:
    go build -o ezpanel ./cmd/panel/

# Run in development mode
dev:
    go run -tags dev ./cmd/panel/

# Run tests
test:
    go test ./...

# Clean build artifacts
clean:
    rm -f ezpanel
