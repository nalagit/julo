package db

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usr Users
	json.NewDecoder(r.Body).Decode(&usr)
	gr := registerUser(usr.Username, usr.Password)

	json.NewEncoder(w).Encode(gr)
}

func InitializeUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()

	if ok {

		tokenDetails, err := generateToken(username, password)

		if err != nil {
			panic(err.Error())
		} else {
			json.NewEncoder(w).Encode(tokenDetails)
		}
	} else {
		gr := GenericResponse{
			Status:  "Failed",
			Message: "Username/Password Required",
		}
		json.NewEncoder(w).Encode(gr)
	}
}

func Validation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	userDetails, err := validateToken(authToken)

	if err != nil {
		gr := GenericResponse{
			Status:  "Failed",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(gr)
	} else {
		gr := GenericResponse{
			Status:  "Success",
			Message: "Auth " + userDetails.Username + " Success",
		}
		json.NewEncoder(w).Encode(gr)
	}

}
