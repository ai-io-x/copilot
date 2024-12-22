# Project: Website Manager

This project is a client-server application with a three-layer architecture. It consists of a React client for managing website information, a REST server in Go for handling CRUD operations, and a separate Go server for storing data in a SQLite database.

## Project Structure

```
website-manager
├── client
│   ├── src
│   │   ├── components
│   │   │   ├── WebsiteForm.tsx
│   │   │   └── WebsiteTable.tsx
│   │   ├── App.tsx
│   │   └── index.tsx
│   ├── Dockerfile
│   ├── package.json
│   └── README.md
├── rest-server
│   ├── cmd
│   │   └── main.go
│   ├── internal
│   │   ├── controller
│   │   │   └── website_controller.go
│   │   └── service
│   │       └── website_service.go
│   ├── Dockerfile
│   └── go.mod
├── db-server
│   ├── cmd
│   │   └── main.go
│   ├── internal
│   │   ├── database
│   │   │   └── sqlite.go
│   │   └── models
│   │       ├── category.go
│   │       ├── tag.go
│   │       └── website.go
│   ├── Dockerfile
│   └── go.mod
├── docker-compose.yml
└── README.md
```

## Client Application

The client is a React application that allows users to input and manage website information. It includes:

- **WebsiteForm**: A form for entering website details such as name, link, description, tags, and categories.
- **WebsiteTable**: A table displaying the list of websites with options to search, sort, update, and delete entries.

## REST Server

The REST server is built in Go and provides endpoints for performing CRUD operations on website data. It includes:

- **Controller**: Handles HTTP requests and responses.
- **Service**: Contains business logic for managing website data.

## Database Server

The database server is also built in Go and manages the SQLite database. It includes:

- **SQLite Functions**: Functions for storing and retrieving website data.
- **Models**: Definitions for the website, category, and tag data structures.

## Docker Setup

The project includes Dockerfiles for each component (client, REST server, and database server) and a `docker-compose.yml` file to orchestrate the services.

## Getting Started

1. Clone the repository.
2. Navigate to the project directory.
3. Run `docker-compose up` to build and start all services. (`docker-compose up --build`)
4. Access the client application in your web browser.

## License

This project is licensed under the MIT License.