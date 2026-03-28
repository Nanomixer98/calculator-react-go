# Calculator Backend Server

This is the backend service for the Calculator application, built entirely in Go using the **Gin Web Framework**. It adheres to the **Hexagonal Architecture** (also known as Ports and Adapters) to cleanly separate the core business logic from the HTTP delivery layer.

## Requirements

- [Go](https://golang.org/doc/install) (version 1.26 or later)
- Workspace connected to the `calculator-react-go` project.

## Directory Structure
- `app/`: Contains the pure business logic and the ports (interfaces) of the calculator.
- `controller/`: The HTTP adapters. It includes endpoint handlers (`calculator.go`), route definitions (`router.go`), payload validation (`bind.go`), and standardized outputs (`response.go`).
- `builder/`: Responsible for dependency injection, wiring up the app logic with the Gin router.
- `models/`: Type definitions and DTOs (Data Transfer Objects) for requests and responses.
- `docs/`: Auto-generated Swagger documentation files.

## Running the Server Locally

1. **Navigate to the server directory**
   Ensure your terminal is located in the `server` folder:
   ```bash
   cd server/
   ```

2. **Download Dependencies**
   Make sure all modules are properly synchronized:
   ```bash
   go mod tidy
   ```

3. **Start the Application**
   Run the main entry file:
   ```bash
   go run main.go
   ```
   *The server will start listening on `http://localhost:8080`.*

---

## Interactive API Documentation (Swagger)

The API endpoints are fully documented using Swaggo. Once the server is running, you can view and test the API directly from your browser.

Navigate to:
👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

### Updating Documentation
If you modify the API or update any `// @Summary` annotations inside `controller/calculator.go`, you must regenerate the Swagger docs before restarting the server. 

To regenerate the documentation, run the Swaggo CLI:
```bash
~/go/bin/swag init
```
*(If you do not have Swaggo installed, you can install it using `go install github.com/swaggo/swag/cmd/swag@latest`)*
