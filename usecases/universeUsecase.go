package usecases

import (
	"errors"
	"fmt"
	"github.com/as-ideas/happy-stars-go/domain"
	"log"
)

type UniverseUsecase struct {
	Galaxy *domain.Galaxy
}

func (u *UniverseUsecase) AddUniverse(newUniverse domain.Universe) error {

	if !newUniverse.IsValid() {
		return errors.New("invalid newUniverse given")
	}

	err := u.Galaxy.AddUniverse(newUniverse)
	if err != nil {
		return fmt.Errorf("failed adding new universe to galaxy: %w", err)
	}

	log.Printf("added newUniverse %v to galaxy", newUniverse)
	return nil
}

func (u *UniverseUsecase) RemoveUniverse(ID string) error {
	err := u.Galaxy.RemoveUniverse(ID)

	if err != nil {
		return fmt.Errorf("failed removing universe from galaxy: %w", err)
	}

	log.Printf("added universe with id=%v from galaxy", ID)
	return nil
}

func (u *UniverseUsecase) UpdateUniverse(universe domain.Universe) error {
	return nil
}

func (u *UniverseUsecase) ListUniverses() []domain.Universe {
	return u.Galaxy.GetUniverses()
}

func (u *UniverseUsecase) GetUniverse(id string) (domain.Universe, error) {
	universe, err := u.Galaxy.GetUniverse(id)

	if err != nil {
		return domain.Universe{}, fmt.Errorf("failed retrieving universe from galaxy: %w", err)
	}

	log.Printf("getting universe with id=%v from galaxy", id)
	return universe, nil
}
