package model

import (
	"database/sql"
	"errors"
)

type user struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

func (u *user) getUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *user) updateUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *user) deleteUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *user) createUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not Implemented")
}
