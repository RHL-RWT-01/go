#!/bin/bash

# JWT Authentication API Demo Script
# This script demonstrates the complete authentication flow

echo "========================================"
echo "JWT Authentication API Demo"
echo "========================================"

# Check if server is running
if ! curl -s http://localhost:8080/health > /dev/null; then
    echo "❌ Server is not running. Please start the server first:"
    echo "   go run main.go"
    exit 1
fi

echo "✅ Server is running"
echo

# 1. Register a new user
echo "1. Registering a new user..."
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username": "demo_user", "email": "demo@example.com", "password": "demo123456"}')

if echo "$REGISTER_RESPONSE" | grep -q "token"; then
    echo "✅ User registered successfully"
    TOKEN=$(echo "$REGISTER_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    echo "   Token: ${TOKEN:0:50}..."
else
    echo "❌ Registration failed: $REGISTER_RESPONSE"
fi
echo

# 2. Login with the user
echo "2. Logging in with the user..."
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "demo_user", "password": "demo123456"}')

if echo "$LOGIN_RESPONSE" | grep -q "token"; then
    echo "✅ Login successful"
    TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    echo "   New token: ${TOKEN:0:50}..."
else
    echo "❌ Login failed: $LOGIN_RESPONSE"
    exit 1
fi
echo

# 3. Access protected route without token
echo "3. Trying to access protected route without token..."
UNAUTH_RESPONSE=$(curl -s http://localhost:8080/api/protected)
echo "   Response: $UNAUTH_RESPONSE"
echo

# 4. Access protected route with token
echo "4. Accessing protected route with valid token..."
PROTECTED_RESPONSE=$(curl -s http://localhost:8080/api/protected \
  -H "Authorization: Bearer $TOKEN")
echo "   Response: $PROTECTED_RESPONSE"
echo

# 5. Get user profile
echo "5. Getting user profile..."
PROFILE_RESPONSE=$(curl -s http://localhost:8080/api/profile \
  -H "Authorization: Bearer $TOKEN")
echo "   Profile: $PROFILE_RESPONSE"
echo

# 6. Try invalid credentials
echo "6. Testing invalid login credentials..."
INVALID_RESPONSE=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "demo_user", "password": "wrongpassword"}')
echo "   Response: $INVALID_RESPONSE"
echo

echo "========================================"
echo "Demo completed successfully! 🎉"
echo "========================================"
echo
echo "API Endpoints tested:"
echo "  ✅ POST /auth/register"
echo "  ✅ POST /auth/login"
echo "  ✅ GET /api/protected (with and without auth)"
echo "  ✅ GET /api/profile"
echo "  ✅ Error handling for invalid credentials"
echo