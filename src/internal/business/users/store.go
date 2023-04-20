package users

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type UserStore struct {
	*sql.DB
}

func NewStore(db *sql.DB) *UserStore {
	return &UserStore{db}
}

func (h *UserStore) create(name, mail, password string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id;"

	var id int
	err := h.QueryRowContext(ctx, query, mail, name, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (h *UserStore) findById(id int) (*BasicUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "SELECT id, email, name FROM users WHERE id = $1"

	result := h.QueryRowContext(ctx, query, id)
	var u BasicUser

	if err := result.Scan(&u.Id, &u.Email, &u.Name); err != nil {
		return nil, err
	}

	return &u, nil
}

func (h *UserStore) findPasswordByEmail(email string) (*UserIdAndPwd, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "SELECT id, password FROM users WHERE email = $1"

	result := h.QueryRowContext(ctx, query, email)
	var user UserIdAndPwd

	if err := result.Scan(&user.Id, &user.Pwd); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (h *UserStore) userExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "SELECT EXISTS(SELECT 1 from users WHERE email=$1)"

	var exists = false
	err := h.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return exists, err
		}
		return exists, nil
	}

	return exists, nil
}
