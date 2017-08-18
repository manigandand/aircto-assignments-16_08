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

/* Validator struct */
type IssueValidator struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	AssignedTo  int    `validate:"required"`
	Status      string `validate:"required"`
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

func checkErr(err error, retErr error) (string, int, error) {
	if err != nil {
		fmt.Println(err)
		return "", 0, retErr
	}
	return "", 0, err
}

func CreateIssue(issueReqBody IssueValidator, createdBy interface{}) (string, int, error) {
	stmt, err := db.Prepare("INSERT INTO issues(title, description, assigned_to, created_by, status) VALUES(?, ?, ?, ?, ?)")
	checkErr(err, errors.New("DB error, unable to preapare statement"))
	res, err := stmt.Exec(issueReqBody.Title, issueReqBody.Description, issueReqBody.AssignedTo, createdBy, issueReqBody.Status)
	checkErr(err, errors.New("DB error, unable to execute statement"))

	id, err := res.LastInsertId()
	checkErr(err, errors.New("DB error, unable to get the last inserted id"))
	fmt.Println("last inserted id:", id)

	return "New Issue successfully created", int(id), nil
	// old
	/*
		_, err = db.Exec("INSERT INTO issues(title, description, assigned_to, created_by, status) VALUES(?, ?, ?, ?, ?)", issueReqBody.Title, issueReqBody.Description, issueReqBody.AssignedTo, createdBy, issueReqBody.Status)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("Server error, unable to create your issue.")
		}
		return "New Issue successfully created", nil
	*/
}

func GetUpdaetIssueId(issueID int, createdBy interface{}) (int, int, error) {
	var id, oldAssignee int
	err := db.QueryRow("SELECT id,assigned_to FROM issues WHERE id=? AND created_by=?", issueID, createdBy).Scan(&id, &oldAssignee)
	if err != nil {
		return oldAssignee, id, errors.New("Unauthorized: You don't have access to update/edit this issue information or the issue id is wrong.")
	}
	return oldAssignee, id, nil
}

func UpdateIssue(issueReqBody IssueValidator, issueID int, createdBy interface{}) (string, error) {
	_, err = db.Exec("UPDATE issues SET title=?, description=?, assigned_to=?, status=? WHERE id=? AND created_by=?", issueReqBody.Title, issueReqBody.Description, issueReqBody.AssignedTo, issueReqBody.Status, issueID, createdBy)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("You don't have access to update/edit this issue information or the issue id is wrong.")
	}
	return "Issue information successfully updated", nil
}

func DeleteIssue(issueID int, createdBy interface{}) (string, error) {
	_, err = db.Exec("DELETE FROM issues WHERE id=? AND created_by=?", issueID, createdBy)
	if err != nil {
		return "", errors.New("You don't have access to delete this issue or the issue id is wrong.")
	}
	return "Issue successfully removed", nil
}

func GetIssuesCreatedByMe(createdBy interface{}) ([]*Issue, error) {
	var issuesRes []*Issue

	rows, err := db.Query("SELECT * FROM issues WHERE created_by=?", createdBy)
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

func GetIssuesAssignedToMe(createdBy interface{}) ([]Issue, error) {
	var issuesRes []Issue

	rows, err := db.Query("SELECT * FROM issues WHERE assigned_to=?", createdBy)
	if err != nil {
		return issuesRes, err
	}
	defer rows.Close()

	for rows.Next() {
		res1 := Issue{}
		err = rows.Scan(&res1.ID, &res1.Title, &res1.Description, &res1.AssignedTo, &res1.CreatedBy, &res1.Status)
		if err != nil {
			return issuesRes, err
		}

		issuesRes = append(issuesRes, res1)
	}
	return issuesRes, nil
}
