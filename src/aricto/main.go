package main

import (
	// authMiddleware "aricto/middleware"
	userController "aricto/controller/user"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

/**
 * [Goriila mux router]
 * @return {[type]} [request response route handler]
 */
func main() {
	rtr := mux.NewRouter()

	str := rtr.PathPrefix("/api").Subrouter()
	str.HandleFunc("/", userController.HandleIndex)
	str.HandleFunc("/login", userController.PostLogin).Methods("POST")

	//user subroute api
	// usr_str := str.PathPrefix("/user").Subrouter()
	// usr_str.HandleFunc("/all-user-list", aricto.GetAllUserList).Methods("GET")

	//issue subroute api
	// isu_str := str.PathPrefix("/issues").Subrouter()
	// isu_str.Handle("/all-issues-list", authMiddleware.JwtMiddleware(aricto.GetAllIssuesList)).Methods("GET")
	// isu_str.Handle("/issue-info", authMiddleware.JwtMiddleware(aricto.GetIssueInfo)).Methods("GET")
	// isu_str.Handle("/create-issue", authMiddleware.JwtMiddleware(aricto.CreateIssue)).Methods("POST")
	// isu_str.Handle("/update-issue", authMiddleware.JwtMiddleware(aricto.UpdateIssue)).Methods("PUT")
	// isu_str.Handle("/delete-issue", authMiddleware.JwtMiddleware(aricto.DeleteIssue)).Methods("DELETE")
	// isu_str.Handle("/issues-by-me", authMiddleware.JwtMiddleware(aricto.GetAllIssuesByMe)).Methods("GET")
	// isu_str.Handle("/issues-for-me", authMiddleware.JwtMiddleware(aricto.GetAllIssuesAssignedToMe)).Methods("GET")

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
