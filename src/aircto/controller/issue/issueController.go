package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"aircto/mail"
	DB "aircto/model"
	responseHandler "aircto/response"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

var message string

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

/**
 * [getAllIssuesList - get all the list of issues]
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
var GetAllIssuesList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(context.Get(r, "userID"))
	// query
	dbResult, err := DB.GetAllIssues()
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}
	data := struct {
		AllIssues []*DB.Issue `json:"all_issues"`
	}{dbResult}

	message = "All issue list successfully retrieved"
	response := responseHandler.ResponseWriter(message, true, data, 200)
	result, _ := json.Marshal(response)

	w.Write([]byte(result))
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
		return
	}
	data := struct {
		IssueDetails DB.Issue `json:"issue_details"`
	}{dbResult}

	message = "Issue information successfully retrieved"
	response := responseHandler.ResponseWriter(message, true, data, 200)
	result, _ := json.Marshal(response)

	w.Write([]byte(result))
})

/**
* Create a issue
 */
var CreateIssue = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	issueReqBody := DB.IssueValidator{}

	_ = json.NewDecoder(r.Body).Decode(&issueReqBody)
	createdBy := context.Get(r, "userID")
	// validate the inputs
	issueValidateStruct := &DB.IssueValidator{
		Title:       issueReqBody.Title,
		Description: issueReqBody.Description,
		AssignedTo:  issueReqBody.AssignedTo,
		Status:      issueReqBody.Status,
	}

	validateErr := validate.Struct(issueValidateStruct)
	if validateErr != nil {
		var erors []string
		for _, validateErr := range validateErr.(validator.ValidationErrors) {
			errMsg := validateErr.Field() + " filed is " + validateErr.Tag()
			erors = append(erors, errMsg)
		}
		errData := struct {
			Errors []string `json:"errors"`
		}{erors}

		result, _ := json.Marshal(responseHandler.LoadValidationErrorResponse(errData))
		w.Write([]byte(result))
		return
	}
	/*
	 check if the assignee id vaild one
	*/
	assigneeDetails, err := DB.GetUser(issueReqBody.AssignedTo)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}

	// we can goahead to save this info
	dbResult, lastIssueID, err := DB.CreateIssue(issueReqBody, createdBy)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}
	// return success response
	issteDetailsRes, err := DB.GetIssue(lastIssueID)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}
	data := struct {
		IssueDetails DB.Issue `json:"issue_details"`
	}{issteDetailsRes}

	message = dbResult

	response := responseHandler.ResponseWriter(message, true, data, 201)
	result, _ := json.Marshal(response)
	/*
		Send Mail to the assignee about the new Issue, (goroutine)
	*/
	fmt.Println("initiated goroutine to send mail...")
	go newIssueMailToUser(issteDetailsRes, assigneeDetails, "New Issue", "You are assigned to a new issue")

	w.Write([]byte(result))
})

/**
* Update a issue
 */
var UpdateIssue = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	issueID, _ := strconv.Atoi(vars["issueID"])
	issueReqBody := DB.IssueValidator{}

	_ = json.NewDecoder(r.Body).Decode(&issueReqBody)
	createdBy := context.Get(r, "userID")
	// validate the inputs
	issueValidateStruct := &DB.IssueValidator{
		Title:       issueReqBody.Title,
		Description: issueReqBody.Description,
		AssignedTo:  issueReqBody.AssignedTo,
		Status:      issueReqBody.Status,
	}

	validateErr := validate.Struct(issueValidateStruct)
	if validateErr != nil {
		var erors []string
		for _, validateErr := range validateErr.(validator.ValidationErrors) {
			errMsg := validateErr.Field() + " filed is " + validateErr.Tag()
			erors = append(erors, errMsg)
		}
		errData := struct {
			Errors []string `json:"errors"`
		}{erors}

		result, _ := json.Marshal(responseHandler.LoadValidationErrorResponse(errData))
		w.Write([]byte(result))
		return
	}
	/*
	 check if the assignee id vaild one
	*/
	assigneeDetails, err := DB.GetUser(issueReqBody.AssignedTo)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}

	// we can goahead to update this info, get issue info by id & createdBy
	oldAssignee, resIssueID, err := DB.GetUpdaetIssueId(issueID, createdBy)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}

	dbResult, err := DB.UpdateIssue(issueReqBody, resIssueID, createdBy)
	//update error
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}

	// return success response
	issteDetailsRes, err := DB.GetIssue(resIssueID)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}
	data := struct {
		IssueDetails DB.Issue `json:"issue_details"`
	}{issteDetailsRes}

	message = dbResult
	response := responseHandler.ResponseWriter(message, true, data, 201)
	result, _ := json.Marshal(response)
	/*
		IF new assignee Send Mail to the assignee about the new Issue, (goroutine)
	*/
	if oldAssignee != issueReqBody.AssignedTo {
		fmt.Println("initiated goroutine to send mail...")
		go newIssueMailToUser(issteDetailsRes, assigneeDetails, "New Issue", "You are reassigned to a new issue")
	}
	w.Write([]byte(result))
})

/**
* Delete issue
 */
var DeleteIssue = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	issueID, _ := strconv.Atoi(vars["issueID"])
	createdBy := context.Get(r, "userID")

	// we can goahead to update this info, get issue info by id & createdBy
	_, resIssueID, err := DB.GetUpdaetIssueId(issueID, createdBy)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, errors.New("You don't have access to delete this issue or the issue id is wrong.")))
		w.Write([]byte(result))
		return
	}

	// delete a issue
	dbResult, err := DB.DeleteIssue(resIssueID, createdBy)
	//delete error
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}
	// return success response
	data := dbResult
	message = dbResult
	response := responseHandler.ResponseWriter(message, true, data, 201)
	result, _ := json.Marshal(response)
	w.Write([]byte(result))
})

/**
* get all the issues created by me
 */
var GetAllIssuesByMe = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createdBy := context.Get(r, "userID")
	dbResult, err := DB.GetIssuesCreatedByMe(createdBy)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}

	data := struct {
		AllIssues []*DB.Issue `json:"all_issues"`
	}{dbResult}

	message = "All issue list created by you successfully retrieved"
	response := responseHandler.ResponseWriter(message, true, data, 200)
	result, _ := json.Marshal(response)
	w.Write([]byte(result))
})

/**
* get all the list of issues assigned to me
 */
var GetAllIssuesAssignedToMe = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createdBy := context.Get(r, "userID")
	dbResult, err := DB.GetIssuesAssignedToMe(createdBy)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}

	data := struct {
		AllIssues []*DB.Issue `json:"all_issues"`
	}{dbResult}

	message = "All issue assigned to you successfully retrieved"
	response := responseHandler.ResponseWriter(message, true, data, 200)
	result, _ := json.Marshal(response)
	w.Write([]byte(result))
})

/**
* Send mail: new ticket
 */
func newIssueMailToUser(issteDetailsRes DB.Issue, assigneeDetails DB.User, title string, message string) {
	// send this mail after 12 minutes to the user
	<-time.After(12 * time.Minute)
	mail.PrepareToSendMail(issteDetailsRes, assigneeDetails, title, message)
}

/**
* Send mail: Issue Report
 */
func IssueInfoCronJob() {
	// get all the issues
	dbResult, err := DB.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dbResult)
	for i, v := range dbResult {
		fmt.Println(v)
		fmt.Printf("%d = %s\n", i, v.Email)
	}
	// DB.GetAllIssueGroupBy()
	// get the assignee id
	// get the user details
	// send mail
}
