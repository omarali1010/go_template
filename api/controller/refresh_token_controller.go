package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/omaraliali1010/go_template/domain"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
}

func (rt *RefreshTokenController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	log.Println("RefreshTokenController called with req ", r)
	var req domain.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	refreshTokenResponse, err := rt.RefreshTokenUsecase.GetRefreshAndAccessToken(req.RefreshToken, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(refreshTokenResponse)
}
