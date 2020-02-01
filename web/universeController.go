package web

import (
	"encoding/json"
	"errors"
	"github.com/as-ideas/happy-stars-go/domain"
	"github.com/as-ideas/happy-stars-go/usecases"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type UniverseController struct {
	UniverseUsecase usecases.UniverseUsecase
}

func (c *UniverseController) AddUniverse(w http.ResponseWriter, r *http.Request) {
	var newUniverse domain.Universe
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "failed adding universe: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &newUniverse)
	if err != nil {
		http.Error(w, "failed adding universe: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = c.UniverseUsecase.AddUniverse(newUniverse)
	if err != nil && errors.Is(err, domain.ErrAlreadyExists) {
		http.Error(w, "failed adding universe: "+err.Error(), http.StatusConflict)
		return
	}

	if err != nil {
		http.Error(w, "failed adding universe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UniverseController) RemoveUniverse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var universeID = vars["id"]
	if len(universeID) == 0 {
		http.Error(w, "missing universe id", http.StatusBadRequest)
		return
	}

	log.Printf("removing universe '%v'", universeID)
	err := c.UniverseUsecase.RemoveUniverse(universeID)

	if err != nil && errors.Is(err, domain.ErrNotFound) {
		http.Error(w, "failed removing universe: "+err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "failed removing universe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *UniverseController) GetUniverse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var universeID = vars["id"]

	if len(universeID) == 0 {
		http.Error(w, "missing universe id", http.StatusBadRequest)
		return
	}

	universe, err := c.UniverseUsecase.GetUniverse(universeID)
	if err != nil && errors.Is(err, domain.ErrNotFound) {
		http.Error(w, "universe not found", http.StatusNotFound)
		return
	}

	log.Printf("returning universe '%v'", universe.ID)
	err = WriteAsJson(w, universe)

	if err != nil {
		http.Error(w, "failed marshalling", http.StatusInternalServerError)
		return
	}
}

func (c *UniverseController) GetUniverses(w http.ResponseWriter, r *http.Request) {

	universes := c.UniverseUsecase.ListUniverses()

	log.Printf("returning %d universes", len(universes))
	err := WriteAsJson(w, universes)

	if err != nil {
		http.Error(w, "failed marshalling", http.StatusInternalServerError)
	}
}
