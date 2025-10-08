package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	PasswordHash string `json:"-"`
}

// Untuk request login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Untuk response login
type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// JWTClaims dipakai di JWT
type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
