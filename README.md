# Calculator Application

A full-stack calculator application built with **React + TypeScript** frontend and **Go** backend, implementing **Hexagonal Architecture (Ports and Adapters)** on both sides for clean separation of concerns.

## Project Structure

```
calculator-react-go/
├── client/              # React + TypeScript + Vite frontend
│   ├── Dockerfile
│   └── nginx.conf
├── server/              # Go backend with Gin framework
│   └── Dockerfile
├── docker-compose.yml   # Docker orchestration
├── README.md            # This file
└── prompt_history/      # Development history
```

## Quick Start

### Option 1: Docker (Recommended)

The fastest way to run the full application:

```bash
# Build and start both services
docker-compose up --build

# Or run in background
docker-compose up -d --build
```

Access the application:
- **Frontend**: http://localhost (puerto 80)
- **API**: http://localhost:8080
- **Swagger Docs**: http://localhost:8080/swagger/index.html

To stop:
```bash
docker-compose down
```

### Option 2: Manual Setup

#### Prerequisites

- [Node.js](https://nodejs.org/) (v18 or later)
- [Go](https://golang.org/doc/install) (v1.26 or later)

### Running the Application

#### 1. Start the Backend Server

```bash
cd server/
go mod tidy
go run main.go
```

The server will start on `http://localhost:8080`.

#### 2. Start the Frontend Client

In a new terminal:

```bash
cd client/
npm install
npm run dev
```

The client will start on `http://localhost:5173` (or another available port).

#### 3. Open the Application

Navigate to the client URL shown in the terminal to use the calculator.

## Architecture Overview

This project implements **Hexagonal Architecture** (also known as Ports and Adapters) on both frontend and backend:

### Backend (Go)
- **Core/App Layer**: Pure business logic with no external dependencies
- **Ports**: Interfaces defining the capabilities the core provides
- **Adapters**: HTTP controllers, request binding, response handling
- **Builder**: Dependency injection wiring everything together

### Frontend (React + TypeScript)
- **Core/Domain**: Business logic, types, and state management
- **Ports**: Interfaces for external communication (API)
- **Adapters**: REST API implementation
- **App Layer**: React hooks orchestrating use cases
- **UI Layer**: React components (pages and reusable components)

## API Documentation

Once the server is running, interactive Swagger documentation is available at:

**[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

### Endpoints

All endpoints accept POST requests at `/api/*`:

| Endpoint | Description | Payload |
|----------|-------------|---------|
| `POST /api/add` | Addition | `{"a": number, "b": number}` |
| `POST /api/subtract` | Subtraction | `{"a": number, "b": number}` |
| `POST /api/multiply` | Multiplication | `{"a": number, "b": number}` |
| `POST /api/divide` | Division | `{"a": number, "b": number}` |
| `POST /api/negate` | Negation (sign change) | `{"value": number}` |
| `POST /api/percentage` | Percentage (value / 100) | `{"value": number}` |

### Response Format

**Success (200)**:
```json
{
  "result": 42.5
}
```

**Error (400)**:
```json
{
  "error": "division by zero is not allowed"
}
```

## Development

### Client Development

```bash
cd client/
npm run dev      # Start dev server
npm run test     # Run tests
npm run test:coverage  # Run tests with coverage
npm run build    # Build for production
npm run lint     # Run ESLint
```

See [client/README.md](client/README.md) for detailed client documentation.

### Server Development

```bash
cd server/
go run main.go              # Start server
go test ./...               # Run all tests
go test -v ./...            # Run tests with verbose output
./run_tests.sh              # Run tests with coverage report
```

See [server/README.md](server/README.md) and [server/TESTING.md](server/TESTING.md) for detailed server documentation.

### Docker Development

```bash
# Build and start all services
docker-compose up --build

# Run in background (detached mode)
docker-compose up -d --build

# View logs
docker-compose logs -f

# View logs for specific service
docker-compose logs -f server
docker-compose logs -f client

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v

# Rebuild specific service
docker-compose up --build server
docker-compose up --build client
```

## Design Rationale

### Why Hexagonal Architecture?

1. **Separation of Concerns**: Business logic is isolated from delivery mechanisms (HTTP, UI)
2. **Testability**: Core logic can be tested without HTTP server or database
3. **Flexibility**: Easy to swap implementations (e.g., REST → GraphQL, React → Vue)
4. **Maintainability**: Changes to one layer don't cascade to others

### Technology Choices

**Backend: Go + Gin**
- Gin provides performant HTTP routing and middleware
- Go's type safety and performance characteristics
- Excellent support for interface-based design (Hexagonal Architecture)

**Frontend: React + TypeScript + Vite**
- React for component-based UI architecture
- TypeScript for type safety across the stack
- Vite for fast development and optimized builds
- Tailwind CSS for utility-first styling
- shadcn/ui components for consistent design

## Features

- Basic arithmetic operations: Add, Subtract, Multiply, Divide
- Unary operations: Negate, Percentage
- Error handling: Division by zero, overflow detection
- Responsive UI with loading states
- Comprehensive test coverage on both frontend and backend

## License

MIT
