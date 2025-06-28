package controller

import (
	"encoding/json"
	"net/http"

	"github.com/omaraliali1010/go_template/domain"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
}

func (sc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	LoginResponse, err := sc.LoginUsecase.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse)
}
