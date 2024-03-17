package database

import (
	"database/sql"
	"server/internal/models"
)

type UserTableActions interface {
	GetOne(username string, password string) (models.User, error)
}

type UsersTable struct {
	Db *sql.DB
}

func (u UsersTable) GetOne(username string, password string) (models.User, error) {
	var user models.User
	db := u.Db

	err := db.QueryRow("SELECT * FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.ID, &user.Username, &user.Password, &user.Token)
	if err != nil {
		return user, err
	}

	return user, nil
}
