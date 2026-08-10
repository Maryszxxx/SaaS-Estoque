package database

import (
	"database/sql"
	"saas-estoque/entity"
)

// implementando user

type PostgresUserRepository struct {
	db *sql.DB
}

func ConnectPostgresUserRepository() *sql.DB {
	return ConnectPostgresUserRepository()
}
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) SaveUser(user *entity.User) error {
	query := `
		INSERT INTO user (
			name,
			email,
			role,
			passwordhash,
			created_at,
			updated_at,
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Role,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}
