# Students API

A RESTful API service built with Go for managing student records. This project demonstrates clean architecture principles with a simple yet effective student management system using SQLite as the database.

## Features

- **Create Student** - Add new student records with validation
- **Get Student by ID** - Retrieve individual student information
- **Get All Students** - List all students in the database
- **Request Validation** - Automatic validation of incoming requests
- **Graceful Shutdown** - Proper signal handling for clean server shutdown
- **Structured Logging** - Built-in logging with `log/slog`
- **Clean Architecture** - Separation of concerns with interface-based storage layer

## Tech Stack

- **Language:** Go 1.24.3
- **Database:** SQLite
- **HTTP Server:** Go's built-in `net/http` package
- **Validation:** `go-playground/validator/v10`
- **Configuration:** `cleanenv` (YAML support)

### Dependencies

- `github.com/go-playground/validator/v10` - Request validation
- `github.com/mattn/go-sqlite3` - SQLite database driver
- `github.com/ilyakaznacheev/cleanenv` - Configuration management
- `github.com/joho/godotenv` - Environment variable loading

## Project Structure

```
students-api/
├── cmd/
│   └── students-api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go     # HTTP request handlers
│   ├── storage/
│   │   ├── storage.go             # Storage interface definition
│   │   └── sqlite/
│   │       └── sqlite.go          # SQLite implementation
│   ├── types/
│   │   └── types.go               # Data models
│   └── utils/
│       └── response/
│           └── response.go         # HTTP response utilities
├── config/
│   └── local.yaml                 # Configuration file
├── go.mod                         # Go module definition
└── go.sum                         # Dependency checksums
```

## Prerequisites

- Go 1.24.3 or higher
- SQLite (included via cgo with go-sqlite3)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/bishal05das/students-api.git
cd students-api
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o students-api ./cmd/students-api
```

## Configuration

The application uses a YAML configuration file located at `config/local.yaml`:

```yaml
env: "dev"                              # Environment (dev/prod)
storage_path: "storage/storage.db"      # SQLite database file path
http_server:
  address: "localhost:8082"             # Server address and port
```

You can specify a custom config file using:
- Environment variable: `CONFIG_PATH=./config/custom.yaml`
- Command-line flag: `./students-api -config ./config/custom.yaml`

## Usage

### Starting the Server

Run the application with the default configuration:

```bash
./students-api -config ./config/local.yaml
```

Or using environment variable:

```bash
CONFIG_PATH=./config/local.yaml ./students-api
```

The server will start on `http://localhost:8082` (default).

### API Endpoints

#### 1. Create Student

Create a new student record.

**Endpoint:** `POST /api/students`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20
}
```

**Response (201 Created):**
```json
{
  "id": 1
}
```

**Validation:**
- `name` - Required
- `email` - Required
- `age` - Required

**Error Response (400 Bad Request):**
```json
{
  "status": "error",
  "error": "validation error message"
}
```

**Example using curl:**
```bash
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","age":20}'
```

#### 2. Get Student by ID

Retrieve a specific student by their ID.

**Endpoint:** `GET /api/students/{id}`

**Path Parameter:**
- `id` - Student ID (integer)

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "status": "error",
  "error": "student not found"
}
```

**Example using curl:**
```bash
curl http://localhost:8082/api/students/1
```

#### 3. Get All Students

Retrieve all students from the database.

**Endpoint:** `GET /api/students`

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane.smith@example.com",
    "age": 22
  }
]
```

**Example using curl:**
```bash
curl http://localhost:8082/api/students
```

## Database

The application uses SQLite as its database. The database file is created automatically at the path specified in the configuration file (`storage/storage.db` by default).

### Student Table Schema

```sql
CREATE TABLE IF NOT EXISTS students (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  email TEXT,
  age INTEGER
)
```

The table is automatically created when the application starts if it doesn't exist.

## Development

### Running in Development

```bash
go run ./cmd/students-api -config ./config/local.yaml
```

### Code Organization

The project follows Clean Architecture principles:

- **Handlers Layer** (`internal/http/handlers/`) - Handles HTTP requests, validates input, and returns responses
- **Storage Layer** (`internal/storage/`) - Interface-based design for data persistence with SQLite implementation
- **Types Layer** (`internal/types/`) - Domain models and data structures
- **Config Layer** (`internal/config/`) - Configuration loading and management
- **Utils Layer** (`internal/utils/`) - Shared utilities and helpers

### Adding New Features

1. Define data models in `internal/types/`
2. Add storage interface methods in `internal/storage/storage.go`
3. Implement storage methods in `internal/storage/sqlite/sqlite.go`
4. Create HTTP handlers in `internal/http/handlers/`
5. Register routes in `cmd/students-api/main.go`

## Graceful Shutdown

The application implements graceful shutdown with a 5-second timeout. It listens for:
- `SIGINT` (Ctrl+C)
- `SIGTERM`

This ensures all ongoing requests are completed before the server shuts down.

## Error Handling

The API uses consistent error responses:

```json
{
  "status": "error",
  "error": "descriptive error message"
}
```

Common HTTP status codes:
- `200 OK` - Successful GET requests
- `201 Created` - Successful POST requests
- `400 Bad Request` - Validation errors or malformed requests
- `500 Internal Server Error` - Server-side errors

## License

This project is available under the MIT License.

## Author

**Bishal Das**
- GitHub: [@bishal05das](https://github.com/bishal05das)
