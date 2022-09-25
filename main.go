package main

import (
	"dumbflix/database"
	"dumbflix/pkg/mysql"
	"dumbflix/routes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// intial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello World")
	})
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")
	fmt.Println("server running localhost:" + port)

	// Embed the setup allowed in 2 parameter on this below code
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
