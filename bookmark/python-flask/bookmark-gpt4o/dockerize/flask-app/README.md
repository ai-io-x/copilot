# Flask Application

This project is a simple Flask application that manages bookmarks. It allows users to add, view, and manage their bookmarks through a web interface.

## Project Structure

```
flask-app
├── src
│   ├── app.py              # Flask application code
│   └── templates
│       └── index.html      # HTML template for the main page
├── Dockerfile               # Dockerfile to build the application image
├── docker-compose.yml       # Docker Compose configuration
└── README.md                # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd flask-app
   ```

2. **Build the Docker image:**
   ```
   docker build -t flask-app .
   ```

3. **Run the application using Docker Compose:**
   ```
   docker-compose up
   ```

4. **Access the application:**
   Open your web browser and go to `http://localhost:5000`.

## Usage

- You can add bookmarks by filling out the form on the main page.
- View and manage your bookmarks directly from the interface.

## Dependencies

- Flask
- SQLite3

## License

This project is licensed under the MIT License.