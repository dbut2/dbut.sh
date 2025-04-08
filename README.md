# dbut.sh

Multi-service platform providing SSH terminal UI, curl-friendly content, and web services.

## Services

- **SSH**: `ssh dbut.sh` - Terminal UI accessible via SSH (port 8022)
- **cURL**: `curl -s https://dbut.sh` - Curl-friendly content endpoint
- **HTTP**: [`https://dbut.sh`](https://dbut.sh) - Web content service

## Development

Build and run locally:
```bash
go build ./...
go run ./go/cmd/ssh/main.go  # Run SSH server
```

Using Docker:
```bash
docker compose up
```

## Generate Content

```bash
go generate ./go/cmd/gen
```