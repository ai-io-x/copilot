# FILE: /client-server-app/client-server-app/README.md

# Client-Server Application

This project is a client-server application built with a three-layer architecture. It consists of a React client, a REST server, and a SQLite server, all designed to handle CRUD operations for website information.

## Project Structure

```
client-server-app
├── client
│   ├── public
│   ├── src
│   │   ├── components
│   │   │   ├── FormComponent.jsx
│   │   │   └── TableComponent.jsx
│   │   ├── App.jsx
│   │   └── index.js
│   ├── Dockerfile
│   └── package.json
├── rest-server
│   ├── src
│   │   ├── main
│   │   │   ├── java
│   │   │   │   └── com
│   │   │   │       └── example
│   │   │   │           └── restserver
│   │   │   │               ├── controller
│   │   │   │               │   └── DataController.java
│   │   │   │               ├── service
│   │   │   │               │   └── DataService.java
│   │   │   │               └── model
│   │   │   │                   └── DataModel.java
│   │   └── resources
│   │       └── application.properties
│   ├── Dockerfile
│   └── pom.xml
├── sqlite-server
│   ├── src
│   │   ├── main
│   │   │   ├── java
│   │   │   │   └── com
│   │   │   │       └── example
│   │   │   │           └── sqliteserver
│   │   │   │               ├── controller
│   │   │   │               │   └── SQLiteController.java
│   │   │   │               ├── service
│   │   │   │               │   └── SQLiteService.java
│   │   │   │               └── model
│   │   │   │                   └── SQLiteModel.java
│   │   └── resources
│   │       └── application.properties
│   ├── Dockerfile
│   └── pom.xml
├── docker-compose.yml
└── README.md
```

## Client

The client is a React application that provides a user interface for inputting and managing website information. It includes:

- **FormComponent**: A form for entering website details such as name, link, description, tags, and categories.
- **TableComponent**: A table that displays the entered website data, allowing users to search, sort, update, and delete entries.

## REST Server

The REST server is built using the Spring Framework and handles CRUD operations for website data. It includes:

- **DataController**: Manages HTTP requests for creating, reading, updating, and deleting website information.
- **DataService**: Contains the business logic for interacting with the SQLite server.

## SQLite Server

The SQLite server is responsible for storing and fetching data from a SQLite database. It includes:

- **SQLiteController**: Handles HTTP requests for data storage and retrieval.
- **SQLiteService**: Contains the logic for interacting with the SQLite database.

## Docker Setup

The project includes Dockerfiles for the client, REST server, and SQLite server, as well as a `docker-compose.yml` file to build and run all services together. To start the application, run:

```
docker-compose up --build
```

## Initial Data

The SQLite database is initialized with example data for demonstration purposes. 

## Requirements

- Docker
- Docker Compose

## License

This project is licensed under the MIT License.