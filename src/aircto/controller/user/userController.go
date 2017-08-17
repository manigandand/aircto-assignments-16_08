package controller

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	// "strconv"

	DB "aircto/model"
	responseHandler "aircto/response"

	// "database/sql"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/context"
)

var message string
var status bool
var statusCode int
var mySigningSecretKey = []byte("qwerty123")

/**
 * [HandleIndex description]
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	message = "Welcome to AirCTO!"
	status = true
	data := "Welcome to AirCTO Test - Simple Issue Tracker - SIT V2.0"
	statusCode = 200
	fmt.Println("Welcome")
	response := responseHandler.ResponseWriter(message, status, data, statusCode)
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))
}

/**
* Login handler
 */
func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := strings.TrimSpace(r.PostFormValue("email"))
	password := strings.TrimSpace(r.PostFormValue("password"))

	dbResult, err := DB.CheckLogin(email, password)
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(401, err))
		w.Write([]byte(result))
		return
	}
	acstkn := getTokenHandler(dbResult)
	data := struct {
		AccessToken string  `json:"access_token"`
		UserDetails DB.User `json:"user_details"`
	}{acstkn, dbResult}

	message = "You have successfully logged in."
	response := responseHandler.ResponseWriter(message, true, data, 200)
	result, _ := json.Marshal(response)

	w.Write([]byte(result))
}

/**
* generate JWT Access Token
 */
func getTokenHandler(res DB.User) string {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)
	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)
	// Set token claims
	claims["id"] = res.ID
	claims["email"] = res.Email
	claims["userName"] = res.UserName
	claims["firstName"] = res.FirstName
	claims["lastName"] = res.LastName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningSecretKey)

	/* Finally, write the token to the browser window */
	return tokenString
}

/**
 * [getAllUserList -reterive all user details ]
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
func GetAllUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbResult, err := DB.GetAllUsers()
	if err != nil {
		result, _ := json.Marshal(responseHandler.LoadErrorResponse(500, err))
		w.Write([]byte(result))
		return
	}
	data := struct {
		UserDetails []*DB.User `json:"user_details"`
	}{dbResult}

	message = "All user list successfully retrieved"
	response := responseHandler.ResponseWriter(message, true, data, 200)
	result, _ := json.Marshal(response)

	w.Write([]byte(result))
}
