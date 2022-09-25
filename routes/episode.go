package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func EpisodeRoutes(r *mux.Router) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(EpisodeRepository)

	r.HandleFunc("/film/{id}/episodes", h.FindEpisodes).Methods("GET")
	r.HandleFunc("/film/{id}/episode/{id}", h.GetEpisode).Methods("GET")
	r.HandleFunc("/episode", middleware.Auth(middleware.UploadFile(h.CreateEpisode))).Methods("POST")
	r.HandleFunc("/episode/{id}", middleware.Auth(h.UpdateEpisode)).Methods("PATCH")
	r.HandleFunc("/episode/{id}", middleware.Auth(h.DeleteEpisode)).Methods("DELETE")
}
