# Antigravity - Architecture Definition

**Assistant**: Antigravity

## Frontend/Backend Separation

```
I want to adjust the client and server behavior.

## Frontend
It will only be responsible for receiving input, displaying results, validating input, and handling errors.

## Backend
I want to expose endpoints to handle calculator operations.
Each endpoint validates input and handles edge cases (division by zero, invalid data).
All endpoints respond in JSON format.
```
