package user

import (
	"github.com/mobamoh/go-playground/events-mgmt-api/internal/db"
	"github.com/mobamoh/go-playground/events-mgmt-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding="required"`
	Password string `binding="required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	pwdHash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, pwdHash)
	// res.LastInsertId()
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var hashedPwd string
	if err := row.Scan(&u.ID, &hashedPwd); err != nil {
		return err
	}
	return utils.CheckPassword(hashedPwd, u.Password)
}
