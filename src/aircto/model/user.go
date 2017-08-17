package model

import (
	"errors"
)

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

func CheckLogin(email string, password string) (User, error) {
	res := User{}
	err = db.QueryRow("SELECT * FROM user WHERE email=? AND password=?", email, password).Scan(&res.ID, &res.Email, &res.UserName, &res.FirstName, &res.LastName, &res.Password, &res.AccessToken)
	if err != nil {
		return res, errors.New("Unauthorized: Wrong Credentials. Unfortunately, your login credentials do not yet have access to the app.")
	}

	return res, nil
}

func GetAllUsers() ([]*User, error) {
	var userRes []*User

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return userRes, err
	}
	defer rows.Close()

	for rows.Next() {
		res1 := &User{}
		err = rows.Scan(&res1.ID, &res1.Email, &res1.UserName, &res1.FirstName, &res1.LastName, &res1.Password, &res1.AccessToken)
		if err != nil {
			return userRes, err
		}

		userRes = append(userRes, res1)
	}
	return userRes, nil
}

/*func (u *user) getUser(db *sql.DB) error {
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


*/
