package auth

import (
	"fmt"
	"time"

	"github.com/Athlevo/Api_booking_Athlevo/config"
	"github.com/dgrijalva/jwt-go"
)

// JWTManager manages JWT tokens.
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

// NewJWTManager creates a new JWTManager.
func NewJWTManager(cfg *config.Config) *JWTManager {
	return &JWTManager{
		secretKey:     cfg.JWTSecretKey,
		tokenDuration: time.Duration(cfg.JWTExpiry) * time.Minute,
	}
}

// Verify verifies the signature of the given JWT token and returns the user claims if valid.
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

// UserClaims represents the claims embedded in a JWT token.
type UserClaims struct {
	jwt.StandardClaims
	ID   string `json:"id"`
	Role string `json:"role"`
	Iat  int64  `json:"iat"`
}

// GetUserID returns the user ID from the token claims.
func (c *UserClaims) GetUserID() string {
	return c.ID
}

// GetUserRole returns the user role from the token claims.
func (c *UserClaims) GetUserRole() string {
	return c.Role
}

// GetUserRole returns the user role from the token claims.
func (c *UserClaims) GetIat() int64 {
	return c.Iat
}
