package model

import (
	"database/sql"
	"errors"
)

type issue struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AssignedTo  int    `json:"assigned_to"`
	CreatedBy   int    `json:"created_by"`
	Status      string `json:"status"`
}

func (u *issue) getIssue(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *issue) updateIssue(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *issue) deleteIssue(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *issue) createIssue(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func getIssues(db *sql.DB, start, count int) ([]issue, error) {
	return nil, errors.New("Not Implemented")
}
