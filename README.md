# Authentication System with Gin APIs

This project implements a complete authentication system using Go and the Gin web framework. It provides JWT-based authentication with user registration, login, and protected routes.

## Features

- **User Registration**: Create new user accounts with email validation
- **User Login**: Authenticate users with username/password
- **JWT Tokens**: Secure authentication using JSON Web Tokens
- **Password Hashing**: Secure password storage using bcrypt
- **Protected Routes**: Middleware-protected endpoints
- **CORS Support**: Cross-origin resource sharing for web applications
- **Error Handling**: Comprehensive error responses
- **In-Memory Storage**: Simple user storage (can be replaced with database)

## API Endpoints

### Public Endpoints

- `GET /` - API information and available endpoints
- `GET /health` - Health check endpoint
- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login with existing credentials

### Protected Endpoints (Require Authentication)

- `GET /auth/profile` - Get current user profile
- `GET /auth/users` - Get all users (admin-like functionality)

## Quick Start

1. **Install dependencies**:
   ```bash
   go mod tidy
   ```

2. **Build the application**:
   ```bash
   go build -o auth-server .
   ```

3. **Run the server**:
   ```bash
   ./auth-server
   ```
   
   Or with custom port:
   ```bash
   PORT=8081 ./auth-server
   ```

4. **The server will start on port 8080 by default**

## API Usage Examples

### Register a new user
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "password": "securepassword123"
  }'
```

### Access protected endpoint
```bash
# Get the token from login response, then:
curl -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  http://localhost:8080/auth/profile
```

## Configuration

The application can be configured using environment variables:

- `PORT` - Server port (default: 8080)
- `JWT_SECRET` - Secret key for JWT signing (default: "your-secret-key")

For production, create a `.env` file or set environment variables:

```bash
export PORT=8080
export JWT_SECRET="your-super-secret-jwt-key"
```

## Project Structure

```
├── auth/
│   ├── handlers/        # HTTP request handlers
│   │   └── auth.go     # Authentication handlers
│   ├── middleware/      # HTTP middleware
│   │   └── auth.go     # JWT authentication middleware
│   ├── models/          # Data models
│   │   ├── user.go     # User model and DTOs
│   │   └── store.go    # In-memory user storage
│   └── utils/           # Utility functions
│       ├── jwt.go      # JWT token utilities
│       └── password.go # Password hashing utilities
├── config/              # Configuration management
│   └── config.go
├── basics/              # Original Go basics examples
│   └── main.go
├── main.go             # Main application entry point
├── go.mod              # Go module dependencies
└── README.md          # This file
```

## Security Features

- **Password Hashing**: Uses bcrypt for secure password storage
- **JWT Tokens**: Stateless authentication with expiration (24 hours)
- **Input Validation**: Request validation using Gin's built-in validators
- **CORS Protection**: Configurable cross-origin resource sharing
- **Error Handling**: Secure error messages without information leakage

## Development

### Adding New Protected Routes

1. Create handler function in `auth/handlers/`
2. Add route to the protected group in `main.go`:
   ```go
   protected.GET("/new-endpoint", handler.NewEndpoint)
   ```

### Extending User Model

Add new fields to the `User` struct in `auth/models/user.go` and update the storage methods accordingly.

### Database Integration

Replace the `UserStore` in `auth/models/store.go` with a database implementation (PostgreSQL, MySQL, MongoDB, etc.).

## Testing

The system has been tested with the following scenarios:
- ✅ User registration with valid data
- ✅ User login with correct credentials
- ✅ Access to protected endpoints with valid JWT
- ✅ Rejection of requests without authorization
- ✅ Rejection of requests with invalid JWT tokens
- ✅ Prevention of duplicate user registration
- ✅ Password hashing and verification

## Production Considerations

1. **Use a strong JWT secret**: Generate using `openssl rand -hex 32`
2. **Enable HTTPS**: Use TLS certificates in production
3. **Database**: Replace in-memory storage with persistent database
4. **Rate Limiting**: Add rate limiting middleware
5. **Logging**: Implement structured logging
6. **Monitoring**: Add health checks and metrics
7. **Environment Variables**: Use proper environment management
8. **CORS**: Configure CORS for specific origins in production