package database

import (
	"database/sql"
	"server/internal/models"
)

type UserTableActions interface {
	GetOne(username string, password string) (models.User, error)
	CheckIfExistByToken(token string) bool
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

func (u UsersTable) GetOneByToken(token string) bool {
	db := u.Db

	count := 0
	err := db.QueryRow("SELECT COUNT(1) FROM users WHERE token = ?", token).Scan(&count)
	if err != nil || count == 0 {
		return false
	}

	return true
}
