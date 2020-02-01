package domain

type Universe struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MaxSize int    `json:"max_size"`
}

func (u Universe) IsValid() bool {
	return len(u.ID) > 0 && len(u.Name) > 0 && u.MaxSize > 0 && u.MaxSize < 1000000
}
