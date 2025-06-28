package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/omaraliali1010/go_template/domain"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc *ProfileController) Fetch(w http.ResponseWriter, r *http.Request) {

	log.Println("ProfileController: x-user-id:", r.Context().Value("x-user-id"))
	userID, err := uuid.Parse(r.Context().Value("x-user-id").(string))

	log.Println("ProfileController: userID:", userID)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, jsonError(err.Error()), http.StatusNotAcceptable)
		return
	}
	profile, err := pc.ProfileUsecase.GetProfileByID(r.Context(), userID)
	if err != nil {
		http.Error(w, jsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
