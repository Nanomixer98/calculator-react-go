# Calculator Backend Server

This is the backend service for the Calculator application, built entirely in Go using the **Gin Web Framework**. It adheres to the **Hexagonal Architecture** (also known as Ports and Adapters) to cleanly separate the core business logic from the HTTP delivery layer.

## Requirements

- [Go](https://golang.org/doc/install) (version 1.26 or later)
- Workspace connected to the `calculator-react-go` project.

## Directory Structure

```
server/
├── core/
│   ├── app/           # Business logic (calculator operations)
│   └── port/          # Interface definitions (CalculatorApp interface)
├── controller/        # HTTP adapters and handlers
│   ├── calculator.go  # HTTP endpoint handlers
│   ├── router.go      # Route definitions
│   ├── bind.go        # Request binding and validation
│   └── response.go    # Response formatting
├── builder/           # Dependency injection
│   └── server.go      # Server construction
├── models/            # DTOs for requests/responses
│   └── calculator.go
├── docs/              # Auto-generated Swagger documentation
├── main.go            # Application entry point
├── go.mod             # Go module definition
└── run_tests.sh       # Test runner script
```

## Architecture

### Hexagonal Architecture (Ports and Adapters)

The server follows Hexagonal Architecture to achieve clear separation:

```
┌─────────────────────────────────────────┐
│           HTTP Layer (Gin)              │
│    ┌─────────────────────────────┐      │
│    │    CalculatorController     │      │
│    │     (HTTP Adapter)          │      │
│    └─────────────────────────────┘      │
│                  │                      │
│                  ▼                      │
│    ┌─────────────────────────────┐      │
│    │     CalculatorApp (Port)    │      │
│    │       (Interface)           │      │
│    └─────────────────────────────┘      │
│                  │                      │
│                  ▼                      │
│    ┌─────────────────────────────┐      │
│    │      calculatorApp            │      │
│    │    (Core Business Logic)      │      │
│    └─────────────────────────────┘      │
└─────────────────────────────────────────┘
```

### Layers

1. **Core/App Layer** (`core/app/`)
   - Pure business logic with no external dependencies
   - Contains calculator operations: Add, Subtract, Multiply, Divide, Negate, Percentage
   - Error handling for division by zero and overflow

2. **Ports** (`core/port/`)
   - `CalculatorApp` interface defining the contract for calculator operations
   - Enables dependency injection and testability

3. **Adapters** (`controller/`)
   - HTTP handlers that receive requests and call core logic
   - Request binding and validation (`bind.go`)
   - Response formatting with consistent structure (`response.go`)
   - Route registration (`router.go`)

4. **Builder** (`builder/`)
   - Wires everything together: creates Gin engine, applies CORS middleware, instantiates app and controller
   - Returns a configured `*gin.Engine` ready to run

## Running the Server Locally

1. **Navigate to the server directory**
   ```bash
   cd server/
   ```

2. **Download Dependencies**
   ```bash
   go mod tidy
   ```

3. **Start the Application**
   ```bash
   go run main.go
   ```
   *The server will start listening on `http://localhost:8080`.*

## API Documentation

### Interactive Documentation (Swagger)

Once the server is running, interactive Swagger documentation is available at:

**[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

### Endpoints

All endpoints accept **POST** requests and return JSON responses.

#### Binary Operations

**POST /api/add**
- Adds two numbers
- Body: `{"a": 5.0, "b": 3.0}`
- Response: `{"result": 8.0}`

**POST /api/subtract**
- Subtracts B from A
- Body: `{"a": 10.0, "b": 4.0}`
- Response: `{"result": 6.0}`

**POST /api/multiply**
- Multiplies two numbers
- Body: `{"a": 7.0, "b": 6.0}`
- Response: `{"result": 42.0}`

**POST /api/divide**
- Divides A by B
- Body: `{"a": 15.0, "b": 3.0}`
- Response: `{"result": 5.0}`
- Error: Returns 400 if B is 0

#### Unary Operations

**POST /api/negate**
- Returns the negative of the input
- Body: `{"value": 5.0}`
- Response: `{"result": -5.0}`

**POST /api/percentage**
- Returns value / 100
- Body: `{"value": 50.0}`
- Response: `{"result": 0.5}`

### Request/Response Format

**Binary Operation Request:**
```json
{
  "a": 10.5,
  "b": 5.2
}
```

**Unary Operation Request:**
```json
{
  "value": 42.0
}
```

**Success Response (200 OK):**
```json
{
  "result": 15.7
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "division by zero is not allowed"
}
```

### Error Codes

| Error | Description |
|-------|-------------|
| `Invalid JSON` | Malformed request body |
| `Missing required field` | Required field (a, b, or value) not provided |
| `NaN or Infinity` | Input contains invalid numeric values |
| `division by zero is not allowed` | Attempted division by zero |
| `result is out of representable range` | Numeric overflow occurred |

## Updating Swagger Documentation

If you modify the API or update any `// @Summary` annotations inside `controller/calculator.go`, regenerate the Swagger docs:

```bash
~/go/bin/swag init
```

*(If Swaggo is not installed: `go install github.com/swaggo/swag/cmd/swag@latest`)*

## Testing

See [TESTING.md](TESTING.md) for comprehensive testing documentation.

### Quick Test Commands

```bash
# Run all tests
go test ./...

# Run with coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

# Run specific package tests
go test ./core/app/...
go test ./controller/...

# Run with verbose output
go test -v ./...
```

## Design Rationale

### Why Hexagonal Architecture?

1. **Testability**: Core business logic can be unit tested without an HTTP server
2. **Flexibility**: Easy to swap delivery mechanisms (HTTP → gRPC, CLI, etc.)
3. **Isolation**: Changes to HTTP layer don't affect business logic
4. **Clarity**: Clear boundaries between layers

### Why Gin?

- High performance with minimal memory footprint
- Rich middleware support (CORS, logging, recovery)
- Excellent routing capabilities
- Built-in JSON binding and validation
- Large ecosystem and community

### Error Handling Strategy

- **Core Layer**: Returns specific errors (ErrDivisionByZero, ErrOverflow)
- **Controller Layer**: Translates errors to HTTP status codes and JSON responses
- **Validation**: Request binding validates JSON structure and numeric values before reaching business logic

### Dependency Injection

The `builder` package uses simple constructor injection:

```go
// builder/server.go
calcApp := app.NewCalculatorApp()
calcController := controller.NewCalculatorController(calcApp)
router := controller.NewRouter(calcController)
```

This allows easy mocking for tests and makes dependencies explicit.

## CORS Configuration

The server includes CORS middleware allowing cross-origin requests:

- **Allowed Origins**: `*` (all origins)
- **Allowed Methods**: `POST, OPTIONS`
- **Allowed Headers**: `Content-Type`

## Dependencies

**Core Dependencies:**
- `github.com/gin-gonic/gin` - HTTP web framework
- `github.com/swaggo/gin-swagger` - Swagger UI middleware
- `github.com/swaggo/swag` - Swagger documentation generator

**Testing Dependencies:**
- `github.com/stretchr/testify` - Assertions and mocking

## License

MIT
