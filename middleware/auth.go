package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const userIDKey = "userID"

// Package middleware provides HTTP middleware for the application.
// It includes authentication middleware to validate JWTs and extract user information.
// This is crucial for securing API endpoints and ensuring that only authenticated users can access certain resources.

// GetUserID retrieves the authenticated user ID from the Gin context.
// It returns the user ID as a string and a boolean indicating if it exists.
// This function is used to identify the user making the request, allowing the application to fetch user-specific data.
func GetUserID(c *gin.Context) (string, bool) {
	val, exists := c.Get(userIDKey)
	if !exists {
		return "", false
	}
	userID, ok := val.(string)
	return userID, ok
}

// AuthMiddleware validates a Supabase JWT and attaches the user ID to the context.
// It checks for the JWT secret in environment variables and parses the token.
// If the token is valid, it extracts the user ID and stores it in the context.
// This middleware is essential for protecting routes that require user authentication.
func AuthMiddleware() gin.HandlerFunc {
	jwtSecret := os.Getenv("SUPABASE_JWT_SECRET")
	if jwtSecret == "" {
		panic("SUPABASE_JWT_SECRET is not set")
	}

	return func(c *gin.Context) {
		tokenStr, ok := extractBearerToken(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Ensure token is using HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		sub, err := extractUserIDFromClaims(token.Claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set(userIDKey, sub)
		c.Next()
	}
}

// extractBearerToken parses the Authorization header for a Bearer token.
// It returns the token string and a boolean indicating if it was found.
// This function is used to extract the JWT from incoming requests, which is then validated by the AuthMiddleware.
func extractBearerToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", false
	}
	return strings.TrimPrefix(authHeader, "Bearer "), true
}

// extractUserIDFromClaims pulls the `sub` field from JWT claims.
// It returns the user ID as a string or an error if the claims are invalid.
// This function is used to extract the user ID from the JWT claims, which is then used to fetch user-specific data.
func extractUserIDFromClaims(claims jwt.Claims) (string, error) {
	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}
	sub, ok := mapClaims["sub"].(string)
	if !ok || sub == "" {
		return "", fmt.Errorf("missing sub in token")
	}
	return sub, nil
}
