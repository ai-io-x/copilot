# Spring Boot REST API

This project is a simple Spring Boot REST API that demonstrates how to create a RESTful service using Java 21 and Spring Boot 3.4.1. The API includes a single controller that responds with a JSON object containing a username.

## Project Structure

```
spring-boot-rest-api
├── src
│   ├── main
│   │   ├── java
│   │   │   └── com
│   │   │       └── example
│   │   │           └── restapi
│   │   │               ├── RestApiApplication.java
│   │   │               └── controller
│   │   │                   └── UserController.java
│   │   └── resources
│   │       ├── application.properties
│   │       └── static
│   └── test
│       └── java
│           └── com
│               └── example
│                   └── restapi
│                       └── RestApiApplicationTests.java
├── .gitignore
├── mvnw
├── mvnw.cmd
├── pom.xml
└── README.md
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd spring-boot-rest-api
   ```

2. **Build the project:**
   ```
   ./mvnw clean install
   ```

3. **Run the application:**
   ```
   ./mvnw spring-boot:run
   ```

4. **Access the API:**
   You can access the API endpoint using the following URL:
   ```
   http://localhost:8080/api/user/{username}
   ```
   Replace `{username}` with the desired username.

## Usage

The API will respond with a JSON object in the following format:
```json
{"name":"username"}
```

## Dependencies

This project uses the following dependencies:
- Spring Boot Starter Web

## License

This project is licensed under the MIT License.