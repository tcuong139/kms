package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"kms_golang/config"
)

// Claims defines the JWT token claims
type Claims struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Auth      int16  `json:"auth"`
	GuardType string `json:"guard_type"` // "web", "customer", "sekosaki"
	jwt.RegisteredClaims
}

// AuthMiddleware verifies JWT tokens for authenticated routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := extractToken(c)
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": "No token provided"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store claims in context
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.UserName)
		c.Set("auth", claims.Auth)
		c.Set("guard_type", claims.GuardType)
		c.Next()
	}
}

// CheckGuardMiddleware verifies the guard type (web=admin/staff, customer, sekosaki)
func CheckGuardMiddleware(guardType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenGuard, exists := c.Get("guard_type")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		if tokenGuard != guardType {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden", "message": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// RoleMiddleware checks if the user has the required role/auth level
// auth: 1=admin, 2=regular user
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if requiredRole == "admin" {
			auth, exists := c.Get("auth")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}
			if auth.(int16) != 1 {
				c.JSON(http.StatusForbidden, gin.H{"error": "forbidden", "message": "Admin access required"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// extractToken retrieves the JWT token from Authorization header or cookie
func extractToken(c *gin.Context) string {
	// Check Authorization header (Bearer token)
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "bearer") {
			return parts[1]
		}
	}

	// Check cookie
	token, err := c.Cookie("token")
	if err == nil && token != "" {
		return token
	}

	// Check query param (for compatibility)
	token = c.Query("token")
	return token
}
