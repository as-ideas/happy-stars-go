package usecases

import "github.com/as-ideas/happy-stars-go/domain"

type StarColor string

const (
	BLUE   StarColor = "BLUE"
	GREEN  StarColor = "GREEN"
	YELLOW StarColor = "YELLOW"
	RED    StarColor = "RED"
)

type InfoUsecase struct {
	Galaxy *domain.Galaxy
}

func (*InfoUsecase) GetAvailableColors() []StarColor {
	return []StarColor{BLUE, GREEN, YELLOW, RED}
}
