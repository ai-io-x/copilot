# README.md

# Bookmarks App

## Overview

The Bookmarks App is a full-stack application that allows users to manage bookmarks with a RESTful API built in Go and a React frontend. The application consists of a client for users to submit and view bookmarks, an admin client for managing the SQLite database, and two servers: one for handling RESTful CRUD operations and another for interacting with the SQLite database.

## Project Structure

```
bookmarks-app
├── client                # React client application
│   ├── src
│   ├── package.json
│   └── Dockerfile
├── admin-client          # Admin React client application
│   ├── src
│   ├── package.json
│   └── Dockerfile
├── rest-server           # REST server application in Go
│   ├── cmd
│   ├── internal
│   ├── go.mod
│   └── Dockerfile
├── db-server             # Database server application in Go
│   ├── cmd
│   ├── internal
│   ├── go.mod
│   └── Dockerfile
├── docker-compose.yml    # Docker Compose configuration
└── README.md             # Project documentation
```

## Features

- **User Client**: A React application that allows users to:
  - Submit website information through a form.
  - View bookmarks in a searchable and sortable table.
  - Update and delete bookmarks.

- **Admin Client**: A React application for administrators to:
  - Log in with credentials.
  - Execute SQL commands on the SQLite database.
  - View data from the SQLite database.

- **REST Server**: A Go application that:
  - Handles CRUD operations for bookmarks.
  - Serves data in JSON format.

- **Database Server**: A Go application that:
  - Manages SQLite database interactions.
  - Stores bookmarks, categories, and tags.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Setup

1. Clone the repository:
   ```
   git clone <repository-url>
   cd bookmarks-app
   ```

2. Build and run the application using Docker Compose:
   ```
   docker-compose up --build
   ```

3. Access the client application at `http://localhost:3000` and the admin client at `http://localhost:3001`.

## License

This project is licensed under the MIT License.