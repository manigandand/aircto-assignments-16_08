package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	responseHandler "aircto/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

var mySigningSecretKey = []byte("qwerty123")

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.1Gvg0ahLLUKTdyBBR-KMOEOu8fnl24UF2_71MiVZdKU"
		// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1hbmlnYW5kYW4uamVmZkBnbWFpbC5jb20iLCJleHAiOjE1MDI2NDE4ODgsImZpcnN0X25hbWUiOiJNYW5pZ2FuZGFuIiwiaWQiOjEsImxhc3RfbmFtZSI6IkRoYXJtYWxpbmdhbSIsInVzZXJfbmFtZSI6Im1hbmlnYW5kYW5qZWZmIn0.0Lu5Vzil8y34fb1AzQEYrKENu4ylkXsY2OVWRjXQXOs"
		tokenString := r.Header.Get("Authorization")
		fmt.Println(tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningSecretKey, nil
		})

		fmt.Println(token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["id"], claims["email"], claims["user_name"], claims["first_name"], claims["last_name"], claims["exp"])

			context.Set(r, "userID", claims["id"])
			context.Set(r, "userEmail", claims["email"])
			context.Set(r, "userName", claims["user_name"])
			context.Set(r, "userFirstName", claims["first_name"])
			context.Set(r, "userLastName", claims["last_name"])

			next.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)

			result, _ := json.Marshal(&responseHandler.Response{
				false,
				"Unauthorized: Token is not valid",
				false,
				"Oops!",
				responseHandler.ErrorType{true, err.Error()},
				401,
				responseHandler.TimeType{int32(time.Now().Unix())}})

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(result))
		}
	})
}
