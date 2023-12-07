package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/iqra-shams/chat-app/utils"
	"github.com/iqra-shams/chat-app/validation"
)

func HandleRefreshToken(w http.ResponseWriter, r *http.Request) {

	//get token form data
	refrehToken := r.FormValue("refreshtoken")

	// validate token and get claims
	claims, err := validation.ValidateToken(refrehToken)
	if err != nil {
		http.Error(w, "can not get claims", http.StatusBadRequest)
	}

	// create new accesstoken with give claims
	accesstoken, _, err := utils.GenerateToken(claims.Username, claims.Role)

	//print new accesstoken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accesstoken)

}
