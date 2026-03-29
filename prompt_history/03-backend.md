# Antigravity - Backend Implementation

**Assistant**: Antigravity (Claude Opus 4.6 + Gemini 3 Flash)

## Hexagonal Architecture Setup

Work on the "server/" folder.

- Use Gin for the exposed endpoints.
- Separate the file structure to adhere to a Hexagonal Architecture, with a builder, controller, and app.
- App handles the business logic.

## Router Separation

In accordance with the architecture, "router.go" should handle initializing the routes, not the "server.go" file.

## Controller Structure

Within the "controller/" folder, create the following files:

- bind.go: handles the request binding functions.
- response.go: handles the responses.

## API Documentation

Document the HTTP API endpoints with Swagger.

## Server Documentation

Generate documentation within "server/" on how to run the server locally.
