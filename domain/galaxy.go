package domain

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrAlreadyExists = errors.New("exists already")
	ErrNotFound      = errors.New("not found")
)

// Galaxy consists of universes and in those universes are stars
type Galaxy struct {
	mu        sync.RWMutex
	stars     []Star
	universes []Universe
}

func NewGalaxy() *Galaxy {
	g := Galaxy{}
	g.mu = sync.RWMutex{}
	return &g
}

func (g *Galaxy) addStarToUniverse(newStar Star, universeName string) error {
	return nil
}

func (g *Galaxy) removeStarFromUniverse(starName string, universeName string) error {
	return nil
}

func (g *Galaxy) GetStars() []Star {
	return g.stars
}

func (g *Galaxy) GetUniverses() []Universe {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.universes
}

func (g *Galaxy) AddUniverse(newUniverse Universe) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, current := range g.universes {
		if current.ID == newUniverse.ID {
			return fmt.Errorf("could not add universe '%s':  %w", newUniverse.ID, ErrAlreadyExists)
		}
	}

	g.universes = append(g.universes, newUniverse)
	return nil
}

func (g *Galaxy) RemoveUniverse(ID string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	idx := -1
	for k, current := range g.universes {
		if current.ID == ID {
			idx = k
		}
	}

	if idx == -1 {
		return fmt.Errorf("could not remove universe '%s': %w", ID, ErrNotFound)
	}

	g.universes = append(g.universes[:idx], g.universes[idx+1:]...)
	return nil
}

func (g *Galaxy) GetUniverse(ID string) (Universe, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	idx := -1
	for k, current := range g.universes {
		if current.ID == ID {
			idx = k
		}
	}

	if idx == -1 {
		return Universe{}, fmt.Errorf("could not find universe '%s': %w", ID, ErrNotFound)
	}

	return g.universes[idx], nil
}
