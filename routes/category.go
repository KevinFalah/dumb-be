package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/categories", h.FindCategories).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")
	r.HandleFunc("/category", middleware.Auth(h.CreateCategory)).Methods("POST")
	r.HandleFunc("/category/{id}", middleware.Auth(h.UpdateCategory)).Methods("PATCH")
	r.HandleFunc("/category/{id}", middleware.Auth(h.DeleteCategory)).Methods("DELETE")
}
