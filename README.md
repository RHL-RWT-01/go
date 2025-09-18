# JWT Authentication API with Gin

A complete JWT-based authentication system implemented in Go using the Gin framework.

## Features

- User registration and login
- JWT token generation and validation
- Password hashing with bcrypt
- Protected routes with middleware
- CORS support
- RESTful API design

## Project Structure

```
├── auth/           # JWT utilities
│   └── jwt.go
├── handlers/       # HTTP route handlers
│   └── auth.go
├── middleware/     # Authentication middleware
│   └── auth.go
├── models/         # Data models
│   └── user.go
├── basics/         # Original Go learning examples
│   └── main.go
├── main.go         # Main server file
├── go.mod          # Go module file
└── README.md       # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/RHL-RWT-01/go.git
cd go
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o jwt-auth-server main.go
```

4. Run the server:
```bash
./jwt-auth-server
```

The server will start on `http://localhost:8080`

## API Endpoints

### Public Endpoints

#### GET /
Returns API information and available endpoints.

#### GET /health
Health check endpoint.

**Response:**
```json
{
  "status": "healthy",
  "message": "JWT Auth API is running"
}
```

#### POST /auth/register
Register a new user.

**Request Body:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com"
  }
}
```

#### POST /auth/login
Login with existing credentials.

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com"
  }
}
```

### Protected Endpoints

All protected endpoints require an `Authorization` header with a valid JWT token:

```
Authorization: Bearer <jwt_token>
```

#### GET /api/profile
Get the authenticated user's profile.

**Response:**
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john@example.com"
}
```

#### GET /api/protected
Example protected route that returns user information.

**Response:**
```json
{
  "message": "This is a protected route",
  "user_id": 1,
  "username": "john_doe"
}
```

## Usage Examples

### Register a new user
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "email": "test@example.com", "password": "password123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'
```

### Access protected route
```bash
curl http://localhost:8080/api/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Security Features

- **Password Hashing**: Uses bcrypt for secure password storage
- **JWT Tokens**: Stateless authentication with 24-hour expiration
- **Middleware Protection**: Routes are protected using JWT validation middleware
- **CORS Support**: Configured for cross-origin requests

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- `400 Bad Request`: Invalid request body or validation errors
- `401 Unauthorized`: Missing or invalid authentication token
- `404 Not Found`: Resource not found
- `409 Conflict`: Username already exists during registration
- `500 Internal Server Error`: Server-side errors

## Development

### Adding New Protected Routes

1. Create a new handler function in `handlers/` directory
2. Add the route to the protected group in `main.go`:

```go
api := r.Group("/api")
api.Use(middleware.JWTAuthMiddleware())
{
    api.GET("/your-route", handlers.YourHandler)
}
```

### Customizing JWT Settings

Edit the `auth/jwt.go` file to modify:
- Token expiration time
- JWT secret key (use environment variables in production)
- Claims structure

## Production Considerations

- Use environment variables for sensitive configuration (JWT secret, database connection)
- Implement proper database storage instead of in-memory maps
- Add rate limiting and request validation
- Use HTTPS in production
- Add logging and monitoring
- Implement refresh token mechanism

## Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt) - JWT implementation
- [golang.org/x/crypto/bcrypt](https://golang.org/x/crypto) - Password hashing

## License

This project is part of a Go learning repository and is available for educational purposes.