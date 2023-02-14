package users

import (
	"context"
	"database/sql"
	"time"
)

type UserStore struct {
	*sql.DB
}

func NewStore(db *sql.DB) *UserStore {
	return &UserStore{db}
}

func (h *UserStore) create(r UserRequest) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3);"

	id, err := h.ExecContext(ctx, query, r.Email, r.Name, r.Password)
	if err != nil {
		return 0, err
	}
	insertId, err := id.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertId), nil
}

func (h *UserStore) findById(id int) (*UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "SELECT id, email, name FROM users WHERE id = $1"

	result := h.QueryRowContext(ctx, query, id)
	var u UserResponse

	if err := result.Scan(&u.Id, &u.Email, &u.Name); err != nil {
		return nil, err
	}

	return &u, nil
}

//func parseRows(result *sql.Rows) (*[]UserRequest, error) {
//	var users []UserRequest
//	var err error
//	for result.Next() {
//		var u UserRequest
//		if err = result.Scan(&u.id, &u.email, &u.name); err != nil {
//			break
//		}
//		users = append(users, u)
//	}
//	return &users, err
//}
