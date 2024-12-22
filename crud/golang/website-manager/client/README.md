# README for Client Application

This is the client application for the Website Manager project, built using React. The client provides a user interface for managing website information, including the ability to add, update, delete, and view websites.

## Features

- **Website Form**: A form to input website details such as name, link, description, tags, and categories.
- **Website Table**: A table that displays all entered websites with options to search, sort, update, and delete entries.

## Getting Started

### Prerequisites

- Node.js (version 14 or higher)
- npm (Node Package Manager)

### Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd website-manager/client
   ```

2. Install dependencies:
   ```
   npm install
   ```

### Running the Application

To start the application in development mode, run:
```
npm start
```

This will start the React application and open it in your default web browser.

### Building for Production

To build the application for production, run:
```
npm run build
```

This will create an optimized build of the application in the `build` directory.

## Docker

To build and run the client application using Docker, use the following command in the root of the project:
```
docker-compose up --build
```

This will build the Docker image and start the client application along with the REST server and database server.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.