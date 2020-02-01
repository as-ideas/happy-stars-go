package domain

import "errors"

type Universe struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MaxSize int    `json:"max_size"`
}

var ErrInvalidUniverse = errors.New("invalid universe")

func (u Universe) Validate() error {
	if len(u.ID) == 0 || len(u.Name) == 0 || u.MaxSize < 0 || u.MaxSize > 1000000 {
		return ErrInvalidUniverse
	}

	return nil
}
