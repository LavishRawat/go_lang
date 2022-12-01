package controllers

import (
	"encoding/json"
	"net/http"
	models "Server/src/models"
	// "gorm.io/gorm"
	// "fmt"
	// "strconv"
	// "github.com/gorilla/mux"
	utils "Server/src/utils"
)

var NewSignUp models.SignUp

func StoreSignupDetails(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type","application/json")
	// var signup SignUp
	// json.NewDecoder(r.Body).Decode(&signup)
	// Database.Create(&signup)
	// json.NewEncoder(w).Encode(signup)
	StoreSignupDetails := &models.SignUp{}
	utils.ParseBody(r, StoreSignupDetails)
	s := StoreSignupDetails.StoreSignupDetails()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}