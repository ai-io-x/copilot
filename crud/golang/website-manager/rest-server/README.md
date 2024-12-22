# REST Server

## Description
REST API server that handles CRUD operations for website bookmarks.

## API Endpoints
- `GET /api/websites` - Get all websites
- `GET /api/websites/category/{name}` - Get websites by category
- `GET /api/websites/tag/{name}` - Get websites by tag
- `POST /api/websites` - Create new website
- `PUT /api/websites/{id}` - Update website
- `DELETE /api/websites/{id}` - Delete website

## Setup
1. Install Go 1.21+
2. Run `go mod download` and ` go mod download github.com/gin-gonic/gin`
3. Run ´go get website-manager/rest-server/internal/controller`
4. Build: `go build -o main ./cmd/main.go`
5. Run: `./main`

## Dependencies
- Go 1.21+
- gorilla/mux - Router
- GORM - ORM