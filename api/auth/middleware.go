package auth

import (
	"net/http"
	"strings"

	"github.com/Athlevo/Api_booking_Athlevo/config"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a Gin middleware function that checks for a valid JWT token.
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if the header is present and in the correct format
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		// Extract the token from the header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Initialize the JWT manager
		jwtManager := NewJWTManager(cfg)

		// Verify the token
		claims, err := jwtManager.Verify(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set the user ID and role in the Gin context
		c.Set("userID", claims.GetUserID())
		c.Set("userRole", claims.GetUserRole())

		// Proceed to the next handler
		c.Next()
	}
}

// AuthorizationMiddleware checks if the authenticated user is authorized to access the resource.
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID and role from the context
		userID, ok := c.Get("userID")
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
			return
		}
		userRole, ok := c.Get("userRole")
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User role not found in context"})
			return
		}

		// Get the resource ID from the request parameters (assuming it's named "userId")
		resourceID := c.Param("user_id")

		// Check if the user is authorized: either admin, doctor, or accessing their own resource
		if userRole == "admin" || userRole == "doctor" || userID == resourceID {
			// User is authorized
			c.Next()
			return
		}

		// User is not authorized
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
	}
}
