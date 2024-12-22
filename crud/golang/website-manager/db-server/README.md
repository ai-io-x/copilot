# Database Server

## Description
SQLite database server for storing website bookmarks data.

## Database Schema

### Websites Table
```sql
CREATE TABLE websites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    description TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE website_tags (
    website_id INTEGER,
    tag_id INTEGER,
    FOREIGN KEY (website_id) REFERENCES websites(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

```

## Setup
1. Install Go 1.21+
2. Install SQLite development files:
   ```bash
   # For Ubuntu/Debian
   sudo apt-get install sqlite3 libsqlite3-dev
   # For CentOS/RHEL
   sudo yum install sqlite sqlite-devel
   ```
3. go mod download
4. Run go mod download github.com/mattn/go-sqlite3
5. Create data directory: `mkdir -p /data`
6. # Plan
1. Add Setup section to DB server README
2. Add Dependencies section with required packages
3. Include SQLite specific setup instructions
4. Add database initialization steps

# Implementation

```markdown


## Setup
1. Install Go 1.21+
2. Install SQLite development files:
   ```bash
   # For Ubuntu/Debian
   sudo apt-get install sqlite3 libsqlite3-dev
   # For CentOS/RHEL
   sudo yum install sqlite sqlite-devel
   ```
3. Run `go mod download`
4. Run `go mod download github.com/mattn/go-sqlite3`
5. Create data directory:
   ```bash
   mkdir -p /data
   ```
6. Build: `go build -o main ./cmd/main.go`
7. Run: `./main`

## Dependencies
- Go 1.21+
- sqlite3 - SQLite driver
- GORM - ORM library
- github.com/mattn/go-sqlite3 - SQLite driver for Go
