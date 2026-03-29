# Calculator Client

A modern React calculator UI built with **TypeScript**, **Vite**, and **Tailwind CSS**. Implements **Hexagonal Architecture** to cleanly separate business logic from UI components.

## Tech Stack

- **Framework**: React 19 with React Compiler
- **Language**: TypeScript 5.9
- **Build Tool**: Vite 8
- **Styling**: Tailwind CSS 4 + tw-animate-css
- **Icons**: Lucide React
- **Testing**: Vitest + React Testing Library
- **UI Components**: shadcn/ui + Radix UI

## Project Structure

```
client/
├── src/
│   ├── adapters/          # External adapters (REST API)
│   │   └── rest/
│   │       └── RestCalculatorApi.ts
│   ├── app/               # Application layer (hooks, use cases)
│   │   └── hooks/
│   │       └── useCalculator.ts
│   ├── core/              # Core domain (business logic)
│   │   ├── domain/        # Types and state
│   │   │   ├── Types.ts
│   │   │   └── CalculatorState.ts
│   │   └── ports/         # Interface definitions
│   │       └── CalculatorApiPort.ts
│   ├── lib/               # Utility functions
│   ├── test/              # Test utilities
│   └── ui/                # UI layer (React components)
│       ├── components/    # Reusable components
│       │   ├── custom/    # Calculator-specific
│       │   └── ui/        # Generic shadcn components
│       ├── pages/         # Page components
│       │   └── Calculator.tsx
│       ├── index.css      # Global styles
│       ├── index.html     # HTML template
│       └── main.tsx       # Entry point
├── public/                # Static assets
├── package.json
├── vite.config.ts
├── tsconfig.json
└── README.md
```

## Setup Instructions

### Prerequisites

- [Node.js](https://nodejs.org/) v18 or later
- npm (comes with Node.js)

### Installation

```bash
cd client/
npm install
```

### Development

Start the development server:

```bash
npm run dev
```

The application will be available at `http://localhost:5173` (or another available port shown in the terminal).

### Building for Production

```bash
npm run build
```

Output is generated in `../dist/` directory.

## Available Scripts

| Command | Description |
|---------|-------------|
| `npm run dev` | Start development server with HMR |
| `npm run build` | Build for production |
| `npm run preview` | Preview production build locally |
| `npm run lint` | Run ESLint |
| `npm test` | Run tests in watch mode |
| `npm run test:coverage` | Run tests with coverage report |

## Architecture

### Hexagonal Architecture (Ports and Adapters)

The client follows Hexagonal Architecture with clear separation:

```
┌─────────────────────────────────────────┐
│           UI Layer (React)              │
│    ┌─────────────────────────────┐      │
│    │    useCalculator hook       │      │
│    │    (Application Layer)      │      │
│    └─────────────────────────────┘      │
│                  │                      │
│                  ▼                      │
│    ┌─────────────────────────────┐      │
│    │     CalculatorApiPort       │      │
│    │       (Interface)           │      │
│    └─────────────────────────────┘      │
│                  │                      │
│                  ▼                      │
│    ┌─────────────────────────────┐      │
│    │    RestCalculatorApi        │      │
│    │      (Adapter)              │      │
│    └─────────────────────────────┘      │
└─────────────────────────────────────────┘
```

### Layers

1. **Core/Domain** (`src/core/`)
   - `Types.ts`: Domain types and API payload/response interfaces
   - `CalculatorState.ts`: Calculator state interface and initial state
   - `CalculatorApiPort.ts`: Interface defining calculator API contract

2. **Adapters** (`src/adapters/`)
   - `RestCalculatorApi.ts`: HTTP implementation of `CalculatorApiPort`
   - Handles fetch requests, JSON parsing, and error handling

3. **Application** (`src/app/`)
   - `useCalculator.ts`: React hook orchestrating calculator operations
   - Manages state transitions and API calls

4. **UI** (`src/ui/`)
   - `pages/Calculator.tsx`: Main calculator page
   - `components/custom/`: Calculator display and keypad components
   - `components/ui/`: Generic shadcn components

## API Communication

The client communicates with the backend via REST API. During development, Vite's proxy configuration forwards `/api/*` requests to `http://localhost:8080`.

See [Root README.md](../README.md) for complete API documentation.

## Testing

Tests are organized by layer following the architecture:

```
src/
├── adapters/rest/
│   └── RestCalculatorApi.test.ts
├── core/
│   ├── domain/
│   │   ├── Types.test.ts
│   │   └── CalculatorState.test.ts
│   └── ports/
│       └── CalculatorApiPort.test.ts
└── ui/pages/
    └── Calculator.test.tsx
```

### Running Tests

```bash
# Run tests in watch mode
npm test

# Run tests once with coverage
npm run test:coverage
```

### Test Coverage

Coverage reports are generated in `coverage/` directory. Open `coverage/index.html` to view the detailed report.

## Vite Configuration

Key configuration details from `vite.config.ts`:

- **Root**: `src/ui` (UI source directory)
- **Public Dir**: `../../public` (static assets)
- **Build Output**: `../../dist` (production build)
- **Proxy**: `/api` → `http://localhost:8080` (backend server)
- **Path Alias**: `@/` maps to `src/`

## ESLint Configuration

TypeScript-aware ESLint with React-specific rules:

- `@eslint/js` - Core ESLint rules
- `typescript-eslint` - TypeScript support
- `eslint-plugin-react-hooks` - React Hooks rules
- `eslint-plugin-react-refresh` - React Refresh rules

## Design Decisions

### Why Hexagonal Architecture in the Frontend?

1. **Testability**: Business logic in `useCalculator` can be tested independently of React components
2. **Flexibility**: Easy to swap the API implementation (e.g., REST → WebSocket)
3. **Clarity**: Clear boundaries between UI, application logic, and external communication

### Why Vite over Create React App?

- Faster cold start and HMR (Hot Module Replacement)
- Native TypeScript support
- Optimized production builds
- Modern development experience

### Why Tailwind CSS?

- Utility-first approach for rapid UI development
- Consistent design system
- Smaller CSS bundle with PurgeCSS
- Easy theming and dark mode support

## React Compiler

The React Compiler is enabled for automatic optimization. See [React Compiler documentation](https://react.dev/learn/react-compiler) for more information.

## License

MIT
