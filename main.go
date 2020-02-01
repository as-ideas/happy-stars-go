package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	"github.com/as-ideas/happy-stars-go/domain"
	"github.com/as-ideas/happy-stars-go/usecases"
	"github.com/as-ideas/happy-stars-go/web"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Printf("using port '%s'\n", port)

	galaxy := domain.NewGalaxy()
	log.Printf("creating galaxy with %d stars in %d universes\n", len(galaxy.GetStars()), len(galaxy.GetUniverses()))

	infoUsecase := usecases.InfoUsecase{Galaxy: galaxy}
	universeUsecase := usecases.UniverseUsecase{Galaxy: galaxy}

	infoController := web.InfoController{InfoUsecase: infoUsecase}
	universeController := web.UniverseController{UniverseUsecase: universeUsecase}

	mux := mux.NewRouter()
	mux.HandleFunc("/api/universes", universeController.GetUniverses).Methods("GET")
	mux.HandleFunc("/api/universes", universeController.AddUniverse).Methods("POST")
	mux.HandleFunc("/api/universes/{id}", universeController.GetUniverse).Methods("GET")
	mux.HandleFunc("/api/universes/{id}", universeController.RemoveUniverse).Methods("DELETE")
	mux.HandleFunc("/api/colors/values", infoController.ServeColorValues)

	log.Printf("starting webserver on port %s ...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}
