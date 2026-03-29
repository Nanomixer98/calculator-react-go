# Antigravity - Frontend Refactoring

**Assistant**: Antigravity

## Hexagonal Architecture Migration

Consider "client/" folder.
The current calculator behaviour must be adapted to consume server REST API endpoints.
Frontend must be refactored to only be responsible of receiving input, displaying results, validating input, and handling errors.
Folder structure must follow an hexagonal archicture as well.
Make an implemententation plan, then guide me to implement it

## Component Refactoring

Simplify the "Calculator.tsx" component by splitting it into smaller components, arrange the components in a folder according to the architecture.

## Display Enhancement

Currently "CalculatorDisplay.tsx" only shows the last digit pressed; I want to show the entire ongoing operation and once a response is received, only show the response.
