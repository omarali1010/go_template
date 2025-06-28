package controller

import (
	"encoding/json"
	"net/http"

	"github.com/omaraliali1010/go_template/domain"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

func (sc *SignupController) Signup(w http.ResponseWriter, r *http.Request) {
	var req domain.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO: this should happen in the usecase , and here we should just create a DTO and send it to the use case
	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	signupResponse, err := sc.SignupUsecase.Create(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(signupResponse)
}
