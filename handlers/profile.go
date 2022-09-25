package handlers

import (
	profiledto "dumbflix/dto/profile"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

// Function di dalam route (/profile)
func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get Isi dari token
	userInfoToken := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfoToken["id"].(float64)) // membuat string menjadi int

	profile, err := h.ProfileRepository.GetProfile(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(profile)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProfile) GetAllProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		ID:        u.ID,
		Phone:     u.Phone,
		Gender:    u.Gender,
		Address:   u.Address,
		Subscribe: u.Subscribe,
	}
}
