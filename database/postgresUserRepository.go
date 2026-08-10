package database

import (
	"database/sql"
	"errors"
	"saas-estoque/entity"
	"time"
)

// implementando user

type PostgresUserRepository struct {
	db *sql.DB
}

func (r *PostgresUserRepository) FindByID(id int64) (*entity.User, error) {
	query := "SELECT id, name, email, role, password_hash, created_at, updated_at  FROM users WHERE id = $1"
	user := &entity.User{}

	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")

		}
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) FindByEmail(email string) (*entity.User, error) {
	query := "SELECT id, name, email, role, password_hash, created_at, updated_at  FROM users WHERE email = $1"
	user := &entity.User{}

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")

		}
		return nil, err
	}
	return user, nil

}

func (r *PostgresUserRepository) Update(user *entity.User) error {
	query := `UPDATE users SET name = $1, email = $2, role = $3, password_hash = $4, updated_at = $5 WHERE id = $6`
	result, err := r.db.Exec(query, user.Name, user.Email, user.Role, user.PasswordHash, time.Now(), user.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *PostgresUserRepository) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Save(user *entity.User) error {
	query := `
		INSERT INTO users (
			name,
			email,
			role,
			password_hash,
			created_at,
			updated_at
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
