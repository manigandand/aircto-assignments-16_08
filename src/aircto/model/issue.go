package model

import (
	"errors"
	"fmt"
)

type Issue struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AssignedTo  int    `json:"assigned_to"`
	CreatedBy   int    `json:"created_by"`
	Status      string `json:"status"`
}

func GetAllIssues() ([]*Issue, error) {
	var issuesRes []*Issue

	rows, err := db.Query("SELECT * FROM issues")
	if err != nil {
		return issuesRes, err
	}
	defer rows.Close()

	for rows.Next() {
		res1 := &Issue{}
		err = rows.Scan(&res1.ID, &res1.Title, &res1.Description, &res1.AssignedTo, &res1.CreatedBy, &res1.Status)
		if err != nil {
			return issuesRes, err
		}

		issuesRes = append(issuesRes, res1)
	}
	return issuesRes, nil
}

func GetIssue(issueID int) (Issue, error) {
	res1 := Issue{}
	err := db.QueryRow("SELECT * FROM issues WHERE id=?", issueID).Scan(&res1.ID, &res1.Title, &res1.Description, &res1.AssignedTo, &res1.CreatedBy, &res1.Status)
	if err != nil {
		return res1, errors.New("Wrong Issue ID: The given issue id is not a valid one.")
	}
	return res1, nil
}

func CreateIssue(issueReqBody Issue, createdBy interface{}) (string, error) {
	_, err := db.Exec("INSERT INTO issues(title, description, assigned_to, created_by, status) VALUES(?, ?, ?, ?, ?)", issueReqBody.Title, issueReqBody.Description, issueReqBody.AssignedTo, createdBy, issueReqBody.Status)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Server error, unable to create your issue.")
	}
	return "New Issue successfully created", nil
}

// func (u *issue) updateIssue(db *sql.DB) error {
// 	return errors.New("Not Implemented")
// }

// func (u *issue) deleteIssue(db *sql.DB) error {
// 	return errors.New("Not Implemented")
// }

// func getIssues(db *sql.DB, start, count int) ([]issue, error) {
// 	return nil, errors.New("Not Implemented")
// }
