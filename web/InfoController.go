package web

import (
	"github.com/as-ideas/happy-stars-go/usecases"
	"log"
	"net/http"
)

type InfoController struct {
	InfoUsecase usecases.InfoUsecase
}

func (c *InfoController) ServeColorValues(w http.ResponseWriter, r *http.Request) {
	err := WriteAsJson(w, c.InfoUsecase.GetAvailableColors())

	if err != nil {
		log.Printf("failed serving color values: %s", err)
		http.Error(w, "failed serving color valued", http.StatusInternalServerError)
	}
}
