package repository

import (
	"backendgo/app/model"
	"backendgo/database"
	"database/sql"
	"errors"
)

// GetUserByUsernameOrEmail ambil user dari database pakai username atau email
func GetUserByUsernameOrEmail(identifier string) (*model.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, created_at
		FROM users
		WHERE username = $1 OR email = $1
	`

	var user model.User
	err := database.DB.QueryRow(query, identifier).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash, 
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
