package main

import (
	issueController "aircto/controller/issue"
	userController "aircto/controller/user"

	authMiddleware "aircto/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
)

var cronJobInterval uint64

func init() {
	cronJobInterval = 24
}

/**
 * [Goriila mux router]
 * @return {[type]} [request response route handler]
 */
func main() {
	/**
	* create a issue report schedular job to send email to the user with issue details assigned to him
	 */
	go func() {
		/*
		* Give Seconds | Minutes | Hours | Days
		 */
		gocron.Every(cronJobInterval).Hours().Do(issueController.IssueInfoCronJob)
		_, time := gocron.NextRun()
		fmt.Println("Next cron job run on: ", time)
		<-gocron.Start()
	}()

	rtr := mux.NewRouter()

	str := rtr.PathPrefix("/api").Subrouter()
	str.HandleFunc("/", userController.HandleIndex)
	str.HandleFunc("/login", userController.PostLogin).Methods("POST")

	/**
	* user api
	 */
	str.HandleFunc("/users", userController.GetAllUserList).Methods("GET") // reterive all user details

	/**
	* issue api
	 */
	str.Handle("/issues", authMiddleware.JwtMiddleware(issueController.GetAllIssuesList)).Methods("GET")               // get all the list of issues
	str.Handle("/issue/{issueID:[0-9]+}", authMiddleware.JwtMiddleware(issueController.GetIssueInfo)).Methods("GET")   // READ - get the single issue details
	str.Handle("/issue", authMiddleware.JwtMiddleware(issueController.CreateIssue)).Methods("POST")                    // CREATE - create a issue
	str.Handle("/issue/{issueID:[0-9]+}", authMiddleware.JwtMiddleware(issueController.UpdateIssue)).Methods("PUT")    // UPDATE - update a issue
	str.Handle("/issue/{issueID:[0-9]+}", authMiddleware.JwtMiddleware(issueController.DeleteIssue)).Methods("DELETE") // DELETE - delete a issue
	str.Handle("/issues-by-me", authMiddleware.JwtMiddleware(issueController.GetAllIssuesByMe)).Methods("GET")
	str.Handle("/issues-for-me", authMiddleware.JwtMiddleware(issueController.GetAllIssuesAssignedToMe)).Methods("GET")

	http.Handle("/", rtr)

	fmt.Println("********************************************************************")
	fmt.Println("*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*")
	fmt.Println("#                                                                  #")
	fmt.Println("#                                                                  #")
	fmt.Println("#                                                                  #")
	fmt.Println("********** AirCTO SIT API - Hit http://localhost:3011/api **********")
	fmt.Println("#                                                                  #")
	fmt.Println("#                                                                  #")
	fmt.Println("#                                                                  #")
	fmt.Println("*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*")
	fmt.Println("********************************************************************")

	http.ListenAndServe(":3011", handlers.LoggingHandler(os.Stdout, rtr))
}
