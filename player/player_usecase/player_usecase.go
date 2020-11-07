package playerusecase

import (
	"game_server/domain"
	"time"
)

type playerUsecase struct {
	playerUC       domain.PlayerRepository
	ContextTimeout time.Duration
}

//NewPlayerUsecase make new player usecase objects
func NewPlayerUsecase(pr domain.PlayerRepository, timeout time.Duration) domain.PlayerUsecase {
	return pr
}
