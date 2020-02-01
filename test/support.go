package test

import (
	"github.com/as-ideas/happy-stars-go/domain"
	"github.com/as-ideas/happy-stars-go/usecases"
	"github.com/as-ideas/happy-stars-go/web"
)

func GivenEmptyGalaxy() *domain.Galaxy {
	return domain.NewGalaxy()
}

func GivenUniverseController(galaxy *domain.Galaxy) *web.UniverseController {
	u := usecases.UniverseUsecase{Galaxy: galaxy}
	c := web.UniverseController{UniverseUsecase: u}
	return &c
}

func GivenInfoController(galaxy *domain.Galaxy) *web.InfoController {
	return &web.InfoController{InfoUsecase: usecases.InfoUsecase{Galaxy: galaxy}}
}
