package domain

import "fmt"

// Galaxy consists of universes and in those universes are stars
type Galaxy struct {
	stars     []Star
	universes []Universe
}

func (g *Galaxy) addStarToUniverse(newStar Star, universeName string) error {
	return nil
}

func (g *Galaxy) removeStarFromUniverse(starName string, universeName string) error {
	return nil
}

func (g *Galaxy) String() string {
	return fmt.Sprintf("stars: %v, universes: %v", g.stars, g.universes)
}

func (g *Galaxy) GetStars() []Star {
	return g.stars
}

func (g *Galaxy) GetUniverses() []Universe {
	return g.universes
}
