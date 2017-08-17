package controller

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"strconv"
	// "strings"
	// "time"

	DB "aircto/model"
	responseHandler "aircto/response"

	// "database/sql"
	// "github.com/dgrijalva/jwt-go"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var message string

/**
 * [getAllIssuesList - get all the list of issues]
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
var GetAllIssuesList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(context.Get(r, "user_id"))
	// query
	dbResult, err := DB.GetAllIssues()
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
	} else {
		data := struct {
			AllIssues []*DB.Issue `json:"all_issues"`
		}{dbResult}

		message = "All issue list successfully retrieved"
		response := responseHandler.ResponseWriter(message, true, data, 200)
		result, _ := json.Marshal(response)

		w.Write([]byte(result))
	}
})

/**
* Get Issue Info by issue id
 */
var GetIssueInfo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	issueID, _ := strconv.Atoi(vars["issueID"])

	dbResult, err := DB.GetIssue(issueID)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
	} else {
		data := struct {
			IssueDetails DB.Issue `json:"issue_details"`
		}{dbResult}

		message = "Issue information successfully retrieved"
		response := responseHandler.ResponseWriter(message, true, data, 200)
		result, _ := json.Marshal(response)

		w.Write([]byte(result))
	}
})

/**
* Create a issue
 */
var CreateIssue = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	issueReqBody := DB.Issue{}
	_ = json.NewDecoder(r.Body).Decode(&issueReqBody)
	createdBy := context.Get(r, "userID")

	dbResult, err := DB.CreateIssue(issueReqBody, createdBy)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
	} else {
		data := dbResult
		message = dbResult
		response := responseHandler.ResponseWriter(message, true, data, 201)
		result, _ := json.Marshal(response)

		w.Write([]byte(result))
	}
})
